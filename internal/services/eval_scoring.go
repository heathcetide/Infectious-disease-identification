// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// ScoringSnapshotPayload is the restorable body stored inside ScoringVersion.SnapshotJSON.
type ScoringSnapshotPayload struct {
	Mode          string                      `json:"mode"`
	PromptText    string                      `json:"promptText"`
	Notes         string                      `json:"notes,omitempty"`
	CreatedAt     time.Time                   `json:"createdAt,omitempty"`
	ExpertWeights []models.SymptomWeighting   `json:"expertWeights"`
	AIWeights     []models.AISymptomWeighting `json:"aiWeights"`
}

// CreateScoringVersion snapshots current mode + prompt text + full weight tables.
func CreateScoringVersion(db *gorm.DB, name, notes, promptText string) (*models.ScoringVersion, error) {
	mode := "expert"
	if m, err := GetAIModeDTO(db); err == nil && m.IsAI {
		mode = "ai"
	}
	if strings.TrimSpace(name) == "" {
		name = time.Now().Format("2006-01-02_15:04") + "_" + mode
	}
	ph := sha256.Sum256([]byte(promptText))
	var weightBlob strings.Builder
	var weights []models.SymptomWeighting
	_ = db.Find(&weights).Error
	for _, w := range weights {
		weightBlob.WriteString(fmt.Sprintf("%s|%s|%.4f;", w.DiseaseTypeName, w.SymptomFieldName, w.WeightScore))
	}
	var aiWeights []models.AISymptomWeighting
	_ = db.Find(&aiWeights).Error
	for _, w := range aiWeights {
		weightBlob.WriteString(fmt.Sprintf("ai|%s|%s|%.4f;", w.DiseaseTypeName, w.SymptomFieldName, w.WeightScore))
	}
	wh := sha256.Sum256([]byte(weightBlob.String()))
	payload := ScoringSnapshotPayload{
		Mode: mode, PromptText: promptText, Notes: notes,
		CreatedAt: time.Now(), ExpertWeights: weights, AIWeights: aiWeights,
	}
	blob, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	_ = db.Model(&models.ScoringVersion{}).Where("is_active = ?", true).Update("is_active", false).Error
	v := &models.ScoringVersion{
		Name: name, Mode: mode, Notes: notes,
		PromptText: promptText, PromptHash: hex.EncodeToString(ph[:8]),
		WeightsHash:  hex.EncodeToString(wh[:8]),
		SnapshotJSON: string(blob),
		IsActive:     true, CreatedAt: time.Now(),
	}
	if err := db.Create(v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func ListScoringVersions(db *gorm.DB, limit int) ([]models.ScoringVersion, error) {
	if limit <= 0 || limit > 100 {
		limit = 30
	}
	var rows []models.ScoringVersion
	// Omit heavy snapshot blob in list; clients fetch detail for view/download/restore.
	err := db.Select("version_id", "name", "mode", "prompt_hash", "prompt_text", "weights_hash", "notes", "is_active", "created_at").
		Order("version_id DESC").Limit(limit).Find(&rows).Error
	return rows, err
}

// GetScoringVersion returns one version including SnapshotJSON.
func GetScoringVersion(db *gorm.DB, id uint) (*models.ScoringVersion, error) {
	var v models.ScoringVersion
	if err := db.First(&v, id).Error; err != nil {
		return nil, err
	}
	return &v, nil
}

// ParseScoringSnapshot unmarshals the restore payload from a version row.
func ParseScoringSnapshot(v *models.ScoringVersion) (*ScoringSnapshotPayload, error) {
	if v == nil {
		return nil, fmt.Errorf("version is nil")
	}
	if strings.TrimSpace(v.SnapshotJSON) != "" {
		var p ScoringSnapshotPayload
		if err := json.Unmarshal([]byte(v.SnapshotJSON), &p); err != nil {
			return nil, fmt.Errorf("invalid snapshot json: %w", err)
		}
		if p.Mode == "" {
			p.Mode = v.Mode
		}
		if p.PromptText == "" {
			p.PromptText = v.PromptText
		}
		return &p, nil
	}
	// Legacy rows without SnapshotJSON: still expose prompt/mode for view/download.
	return &ScoringSnapshotPayload{
		Mode: v.Mode, PromptText: v.PromptText, Notes: v.Notes, CreatedAt: v.CreatedAt,
	}, nil
}

// RestoreScoringVersion reapplies mode + weight tables from a snapshot and marks it active.
func RestoreScoringVersion(db *gorm.DB, id uint) (*models.ScoringVersion, error) {
	v, err := GetScoringVersion(db, id)
	if err != nil {
		return nil, err
	}
	payload, err := ParseScoringSnapshot(v)
	if err != nil {
		return nil, err
	}
	if len(payload.ExpertWeights) == 0 && len(payload.AIWeights) == 0 {
		return nil, fmt.Errorf("该快照没有可恢复的权重数据（请重新保存一份完整快照）")
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		isAI := strings.EqualFold(payload.Mode, "ai") || strings.EqualFold(payload.Mode, "hybrid")
		if _, err := SetSystemAIMode(tx, isAI); err != nil {
			return err
		}
		for _, w := range payload.ExpertWeights {
			q := tx.Model(&models.SymptomWeighting{})
			if w.SymptomWeightingID > 0 {
				q = q.Where("symptom_weighting_id = ?", w.SymptomWeightingID)
			} else {
				q = q.Where("disease_type_name = ? AND symptom_field_name = ?", w.DiseaseTypeName, w.SymptomFieldName)
			}
			if err := q.Update("weight_score", w.WeightScore).Error; err != nil {
				return err
			}
		}
		for _, w := range payload.AIWeights {
			q := tx.Model(&models.AISymptomWeighting{})
			if w.SymptomWeightingID > 0 {
				q = q.Where("symptom_weighting_id = ?", w.SymptomWeightingID)
			} else {
				q = q.Where("disease_type_name = ? AND symptom_field_name = ?", w.DiseaseTypeName, w.SymptomFieldName)
			}
			if err := q.Update("weight_score", w.WeightScore).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(&models.ScoringVersion{}).Where("is_active = ?", true).Update("is_active", false).Error; err != nil {
			return err
		}
		return tx.Model(&models.ScoringVersion{}).Where("version_id = ?", id).Update("is_active", true).Error
	})
	if err != nil {
		return nil, err
	}
	v.IsActive = true
	v.Mode = payload.Mode
	return v, nil
}

// RunTopKEval compares expert/AI ranked diseases against confirmed diagnoses (P5).
func RunTopKEval(db *gorm.DB, name string, k int, from, to *time.Time, versionID uint) (*models.EvalRun, error) {
	if k < 1 {
		k = 3
	}
	if versionID == 0 {
		var v models.ScoringVersion
		if err := db.Where("is_active = ?", true).Order("version_id DESC").First(&v).Error; err == nil {
			versionID = v.VersionID
		}
	}
	q := db.Model(&models.DiagnosisResults{}).Where("disease_type <> '' AND disease_type IS NOT NULL")
	if from != nil {
		q = q.Where("submission_time >= ?", *from)
	}
	if to != nil {
		q = q.Where("submission_time < ?", to.Add(24*time.Hour))
	}
	var diagnoses []models.DiagnosisResults
	if err := q.Order("diagnosis_results_id DESC").Limit(2000).Find(&diagnoses).Error; err != nil {
		return nil, err
	}
	run := &models.EvalRun{
		VersionID: versionID, Name: name, K: k,
		DateFrom: from, DateTo: to, CreatedAt: time.Now(),
	}
	if run.Name == "" {
		run.Name = "eval_" + time.Now().Format("20060102_150405")
	}
	if err := db.Create(run).Error; err != nil {
		return nil, err
	}

	var e1, ek, a1, ak, n int
	for _, d := range diagnoses {
		gold := normalizeDiseaseName(d.DiseaseType)
		if gold == "" || gold == "其他" || gold == "健康" {
			continue
		}
		day := time.Now()
		if d.SubmissionTime != nil {
			day = *d.SubmissionTime
		}
		dayStart := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
		dayEnd := dayStart.Add(24 * time.Hour)

		var status models.DailyHealthStatus
		_ = db.Where("user_id = ? AND check_in_date >= ? AND check_in_date < ?", d.UserID, dayStart, dayEnd).
			Order("status_id DESC").First(&status).Error
		if status.StatusID == 0 {
			continue
		}

		expertTop := topNamesFromTable(db, &models.DiseaseScoring{}, status.StatusID, dayStart, k)
		aiTop := topNamesFromTable(db, &models.AIDiseaseScoring{}, status.StatusID, dayStart, k)
		if len(aiTop) == 0 {
			aiTop = expertTop
		}
		// Skip diagnoses that have a check-in row but no scored diseases that day.
		if len(expertTop) == 0 && len(aiTop) == 0 {
			continue
		}
		c := models.EvalCase{
			RunID: run.RunID, DiagnosisID: d.DiagnosisResultsID, UserID: d.UserID,
			GoldDisease: gold,
			ExpertHitAt1: hitAt(expertTop, gold, 1),
			ExpertHitAtK: hitAt(expertTop, gold, k),
			AIHitAt1:     hitAt(aiTop, gold, 1),
			AIHitAtK:     hitAt(aiTop, gold, k),
		}
		eb, _ := json.Marshal(expertTop)
		ab, _ := json.Marshal(aiTop)
		c.ExpertTopJSON, c.AITopJSON = string(eb), string(ab)
		_ = db.Create(&c).Error
		n++
		if c.ExpertHitAt1 {
			e1++
		}
		if c.ExpertHitAtK {
			ek++
		}
		if c.AIHitAt1 {
			a1++
		}
		if c.AIHitAtK {
			ak++
		}
	}
	run.CaseCount = n
	if n > 0 {
		run.ExpertHitAt1 = float64(e1) / float64(n)
		run.ExpertHitAtK = float64(ek) / float64(n)
		run.AIHitAt1 = float64(a1) / float64(n)
		run.AIHitAtK = float64(ak) / float64(n)
	}
	_ = db.Save(run).Error
	return run, nil
}

func topNamesFromTable(db *gorm.DB, model any, statusID uint, day time.Time, k int) []string {
	type row struct {
		DiseaseTypeName string
		Score           float64
	}
	var rows []row
	switch model.(type) {
	case *models.DiseaseScoring:
		var src []models.DiseaseScoring
		_ = db.Where("status_id = ? AND symptom_time = ?", statusID, day).Find(&src).Error
		for _, s := range src {
			rows = append(rows, row{s.DiseaseTypeName, s.Score})
		}
	case *models.AIDiseaseScoring:
		var src []models.AIDiseaseScoring
		_ = db.Where("status_id = ? AND symptom_time = ?", statusID, day).Find(&src).Error
		for _, s := range src {
			rows = append(rows, row{s.DiseaseTypeName, s.Score})
		}
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Score > rows[j].Score })
	out := make([]string, 0, k)
	for i := 0; i < len(rows) && i < k; i++ {
		out = append(out, rows[i].DiseaseTypeName)
	}
	return out
}

func hitAt(top []string, gold string, k int) bool {
	for i := 0; i < len(top) && i < k; i++ {
		if diseaseNameMatch(top[i], gold) {
			return true
		}
	}
	return false
}

func normalizeDiseaseName(s string) string {
	return strings.TrimSpace(s)
}

func diseaseNameMatch(a, b string) bool {
	a, b = strings.TrimSpace(a), strings.TrimSpace(b)
	if a == "" || b == "" {
		return false
	}
	if strings.EqualFold(a, b) {
		return true
	}
	return strings.Contains(a, b) || strings.Contains(b, a)
}

func ListEvalRuns(db *gorm.DB, limit int) ([]models.EvalRun, error) {
	if limit <= 0 || limit > 50 {
		limit = 20
	}
	var rows []models.EvalRun
	err := db.Order("run_id DESC").Limit(limit).Find(&rows).Error
	return rows, err
}

// LatestEvalCompare returns the newest run for screen / dashboard.
func LatestEvalCompare(db *gorm.DB) (*models.EvalRun, error) {
	var run models.EvalRun
	err := db.Order("run_id DESC").First(&run).Error
	if err != nil {
		return nil, err
	}
	return &run, nil
}

// ListEvalCases returns per-case Top-K outcomes for a run (答辩明细).
func ListEvalCases(db *gorm.DB, runID uint, limit int) ([]models.EvalCase, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	var rows []models.EvalCase
	err := db.Where("run_id = ?", runID).Order("case_id ASC").Limit(limit).Find(&rows).Error
	return rows, err
}
