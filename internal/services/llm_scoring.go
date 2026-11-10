// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/llm"
	"github.com/heathcetide/infectious-disease/internal/models"
)

type llmDiseaseScore struct {
	DiseaseTypeName string   `json:"diseaseTypeName"`
	Score           float64  `json:"score"`
	Reason          string   `json:"reason,omitempty"`
	Evidence        []string `json:"evidence,omitempty"`
}

type llmScoreResponse struct {
	Diseases []llmDiseaseScore `json:"diseases"`
}

// ComputeLLMScores asks the configured LLM to rank disease risks from active symptoms,
// persists AIDiseaseScoring (with reason/evidence), and returns DiseaseScoring-shaped rows.
func ComputeLLMScores(ctx context.Context, db *gorm.DB, client *llm.Client, statusID uint, symptoms AllSymptoms) ([]models.DiseaseScoring, map[string]llmDiseaseScore, error) {
	if client == nil || !client.Enabled() {
		return nil, nil, fmt.Errorf("llm not enabled")
	}

	active := activeSymptomFields(symptoms)
	fields := make([]string, 0, len(active))
	for f := range active {
		fields = append(fields, f)
	}
	sort.Strings(fields)

	var userID uint
	var st models.DailyHealthStatus
	if err := db.First(&st, statusID).Error; err == nil {
		userID = st.UserID
	}
	scoreCtx := loadScoringContext(db, userID, statusID)

	catalog := models.JavaScoringDiseaseNames
	system := `你是高原沿线传染病风险筛查助手。综合：阳性症状、暴露史、海拔/驻地、本月气压气温、既往病史、季节先验、近期血常规/影像 OCR、以及皮肤皮疹/溃疡视觉辅助标签（YOLO/VLM，仅作辅助特征，不可单独下诊断），评估病种风险（0-100）。
只输出 JSON，不要 Markdown。格式：
{"diseases":[{"diseaseTypeName":"...","score":12.5,"reason":"因A、B症状+季节/海拔/检验/视觉先验指向该病","evidence":["hasFever","hasCough"]}]}
规则：
1. diseaseTypeName 必须使用给定病种名称原文字符串，不要改写。
2. 每个给定病种都要给出一条；无关病种给低分（通常 0-15）。
3. score 为 0-100；症状越吻合分越高，并按上下文、OCR 与视觉辅助特征微调。
4. reason 必须用中文说明「阳性症状/暴露/检验影像/视觉辅助 → 病种」以及用到的先验，不超过 60 字；不得留空。
5. evidence 必须是当日阳性症状字段名的子集（可为空数组）；高分病种至少 1 个证据字段。
6. 视觉标签置信度低或为 normal_skin/other 时忽略；不得仅凭图片给出高分。`

	user := fmt.Sprintf("候选病种：\n- %s\n\n当日阳性症状字段：\n- %s\n\n多源上下文：\n%s",
		strings.Join(catalog, "\n- "),
		strings.Join(fields, "\n- "),
		formatContextForLLM(scoreCtx),
	)
	if len(fields) == 0 {
		user += "\n（无阳性症状，请整体给极低分，reason 写「无明显阳性症状」）"
	}

	raw, err := client.Chat(ctx, system, user)
	if err != nil {
		return nil, nil, err
	}
	parsed, err := parseLLMScoreJSON(raw)
	if err != nil {
		return nil, nil, err
	}

	byName := map[string]llmDiseaseScore{}
	for _, d := range parsed.Diseases {
		name := strings.TrimSpace(d.DiseaseTypeName)
		if name == "" {
			continue
		}
		if d.Score < 0 {
			d.Score = 0
		}
		if d.Score > 100 {
			d.Score = 100
		}
		d.Reason = strings.TrimSpace(d.Reason)
		if d.Reason == "" {
			d.Reason = "模型未给出依据"
		}
		// Keep only evidence that appears in today's positive fields.
		if len(d.Evidence) > 0 && len(active) > 0 {
			filtered := make([]string, 0, len(d.Evidence))
			for _, e := range d.Evidence {
				e = strings.TrimSpace(e)
				if _, ok := active[e]; ok {
					filtered = append(filtered, e)
				}
			}
			d.Evidence = filtered
		}
		byName[name] = d
	}

	today := time.Now()
	dayStart := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	if err := db.Where("status_id = ? AND symptom_time = ?", statusID, dayStart).
		Delete(&models.AIDiseaseScoring{}).Error; err != nil {
		return nil, nil, err
	}

	typeIDByName := diseaseTypeIDs(db)
	out := make([]models.DiseaseScoring, 0, len(catalog))
	for _, name := range catalog {
		item, ok := byName[name]
		if !ok {
			for k, v := range byName {
				if strings.Contains(name, k) || strings.Contains(k, name) {
					item = v
					ok = true
					break
				}
			}
		}
		score := 0.0
		reason := ""
		evJSON := "[]"
		if ok {
			score = item.Score
			reason = item.Reason
			if b, err := json.Marshal(item.Evidence); err == nil {
				evJSON = string(b)
			}
			byName[name] = item
		} else {
			reason = "未返回该病种评分"
		}
		aiRow := models.AIDiseaseScoring{
			StatusID:        statusID,
			SymptomTime:     &dayStart,
			DiseaseTypeID:   typeIDByName[name],
			DiseaseTypeName: name,
			Score:           score,
			Reason:          reason,
			EvidenceJSON:    evJSON,
		}
		if err := db.Create(&aiRow).Error; err != nil {
			return nil, nil, err
		}
		out = append(out, models.DiseaseScoring{
			StatusID:        statusID,
			SymptomTime:     &dayStart,
			DiseaseTypeID:   typeIDByName[name],
			DiseaseTypeName: name,
			Score:           score,
		})
	}
	// Soft local priors as a safety net if the model underuses context.
	notes := applyContextFusion(out, scoreCtx, active)
	for _, s := range out {
		updates := map[string]any{"score": s.Score}
		if item, ok := byName[s.DiseaseTypeName]; ok {
			if why, ok2 := notes[s.DiseaseTypeName]; ok2 && why != "" && !strings.Contains(item.Reason, "先验") {
				item.Reason = strings.TrimSpace(item.Reason + "；" + why)
				if len([]rune(item.Reason)) > 80 {
					item.Reason = string([]rune(item.Reason)[:80]) + "…"
				}
				byName[s.DiseaseTypeName] = item
			}
			updates["reason"] = byName[s.DiseaseTypeName].Reason
		}
		_ = db.Model(&models.AIDiseaseScoring{}).
			Where("status_id = ? AND disease_type_name = ?", statusID, s.DiseaseTypeName).
			Updates(updates).Error
	}
	return out, byName, nil
}

func parseLLMScoreJSON(raw string) (*llmScoreResponse, error) {
	raw = strings.TrimSpace(raw)
	if i := strings.Index(raw, "{"); i >= 0 {
		if j := strings.LastIndex(raw, "}"); j > i {
			raw = raw[i : j+1]
		}
	}
	var out llmScoreResponse
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		return nil, fmt.Errorf("parse llm json: %w; body=%s", err, truncateRunes(raw, 240))
	}
	if len(out.Diseases) == 0 {
		return nil, fmt.Errorf("llm json: empty diseases")
	}
	return &out, nil
}

func diseaseTypeIDs(db *gorm.DB) map[string]int {
	out := map[string]int{}
	var rows []models.DiseaseType
	if err := db.Find(&rows).Error; err != nil {
		return out
	}
	for _, r := range rows {
		out[r.DiseaseTypeName] = int(r.DiseaseTypeID)
	}
	return out
}

func truncateRunes(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n]) + "…"
}

// ComputeAIWeightScores is the offline-weight fallback when LLM fails / is unset.
func ComputeAIWeightScores(db *gorm.DB, statusID uint, symptoms AllSymptoms) ([]models.DiseaseScoring, error) {
	var weights []models.AISymptomWeighting
	if err := db.Find(&weights).Error; err != nil {
		return nil, err
	}
	projected := make([]models.SymptomWeighting, 0, len(weights))
	for _, w := range weights {
		projected = append(projected, models.SymptomWeighting(w))
	}

	active := activeSymptomFields(symptoms)
	today := time.Now()
	dayStart := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	if err := db.Where("status_id = ? AND symptom_time = ?", statusID, dayStart).
		Delete(&models.AIDiseaseScoring{}).Error; err != nil {
		return nil, err
	}

	out := make([]models.DiseaseScoring, 0, len(models.JavaScoringDiseaseNames))
	for _, diseaseName := range models.JavaScoringDiseaseNames {
		score := scoreForDisease(projected, diseaseName, active)
		typeID := 0
		for _, w := range projected {
			if strings.EqualFold(w.DiseaseTypeName, diseaseName) {
				typeID = w.DiseaseTypeID
				break
			}
		}
		ev := evidenceForDisease(projected, diseaseName, active)
		evJSON := "[]"
		if b, err := json.Marshal(ev); err == nil {
			evJSON = string(b)
		}
		reason := "离线 AI 权重匹配"
		if len(ev) > 0 {
			reason = "依据权重命中：" + strings.Join(ev, "、")
			if len([]rune(reason)) > 60 {
				reason = string([]rune(reason)[:60])
			}
		}
		aiRow := models.AIDiseaseScoring{
			StatusID: statusID, SymptomTime: &dayStart,
			DiseaseTypeID: typeID, DiseaseTypeName: diseaseName, Score: score,
			Reason: reason, EvidenceJSON: evJSON,
		}
		if err := db.Create(&aiRow).Error; err != nil {
			return nil, err
		}
		out = append(out, models.DiseaseScoring{
			StatusID: statusID, SymptomTime: &dayStart,
			DiseaseTypeID: typeID, DiseaseTypeName: diseaseName, Score: score,
		})
	}
	var userID uint
	var st models.DailyHealthStatus
	if err := db.First(&st, statusID).Error; err == nil {
		userID = st.UserID
	}
	scoreCtx := loadScoringContext(db, userID, statusID)
	notes := applyContextFusion(out, scoreCtx, active)
	for i := range out {
		name := out[i].DiseaseTypeName
		updates := map[string]any{"score": out[i].Score}
		if why, ok := notes[name]; ok && why != "" {
			updates["reason"] = why
		}
		_ = db.Model(&models.AIDiseaseScoring{}).
			Where("status_id = ? AND disease_type_name = ?", statusID, name).
			Updates(updates).Error
	}
	log.Printf("ai scoring: used offline ai_symptom_weighting fallback for status=%d", statusID)
	return out, nil
}

func evidenceForDisease(weights []models.SymptomWeighting, diseaseName string, active map[string]struct{}) []string {
	out := make([]string, 0, 8)
	for _, w := range weights {
		if !strings.EqualFold(w.DiseaseTypeName, diseaseName) {
			continue
		}
		if _, ok := active[w.SymptomFieldName]; ok && w.WeightScore > 0 {
			out = append(out, w.SymptomFieldName)
		}
	}
	sort.Strings(out)
	if len(out) > 6 {
		out = out[:6]
	}
	return out
}

func topDiseaseNames(scores []models.DiseaseScoring, k int) []string {
	cp := make([]models.DiseaseScoring, len(scores))
	copy(cp, scores)
	sort.Slice(cp, func(i, j int) bool { return cp[i].Score > cp[j].Score })
	if k > len(cp) {
		k = len(cp)
	}
	out := make([]string, 0, k)
	for i := 0; i < k; i++ {
		out = append(out, cp[i].DiseaseTypeName)
	}
	return out
}

// expertAIDisagree reports Top1 mismatch or Top3 intersection < 2.
func expertAIDisagree(expert, ai []models.DiseaseScoring) bool {
	eTop := topDiseaseNames(expert, 3)
	aTop := topDiseaseNames(ai, 3)
	if len(eTop) == 0 || len(aTop) == 0 {
		return false
	}
	if eTop[0] != aTop[0] {
		return true
	}
	set := map[string]struct{}{}
	for _, n := range eTop {
		set[n] = struct{}{}
	}
	overlap := 0
	for _, n := range aTop {
		if _, ok := set[n]; ok {
			overlap++
		}
	}
	return overlap < 2
}
