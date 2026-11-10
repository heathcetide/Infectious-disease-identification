// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/llm"
	"github.com/heathcetide/infectious-disease/internal/models"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
)

// VisionLabel is one predicted skin/symptom class (auxiliary feature only).
type VisionLabel struct {
	Label        string  `json:"label"`
	SymptomField string  `json:"symptomField"`
	Confidence   float64 `json:"confidence"`
}

var visionClassCatalog = []VisionLabel{
	{Label: "painful_red_rash", SymptomField: "HasPainfulRedRash"},
	{Label: "skin_ulcer", SymptomField: "HasSkinUlcer"},
	{Label: "itchy_rash", SymptomField: "HasItchyRash"},
	{Label: "petechial_rash", SymptomField: "HasCongestiveOrPetechialRash"},
	{Label: "blood_blister", SymptomField: "HasBloodBlisters"},
	{Label: "cyanosis", SymptomField: "HasCyanosis"},
	{Label: "normal_skin", SymptomField: ""},
	{Label: "other", SymptomField: ""},
}

func fieldForLabel(label string) string {
	for _, c := range visionClassCatalog {
		if strings.EqualFold(c.Label, label) {
			return c.SymptomField
		}
	}
	return ""
}

func isLesionLabel(label string) bool {
	switch strings.ToLower(strings.TrimSpace(label)) {
	case "painful_red_rash", "skin_ulcer", "itchy_rash", "petechial_rash", "blood_blister", "cyanosis":
		return true
	default:
		return false
	}
}

// calibrateVisionLabels soft-priors normal_skin when top lesion confidence is weak
// (common with class-imbalanced YOLO checkpoints that over-fire blood_blister/ulcer).
func calibrateVisionLabels(labels []VisionLabel) []VisionLabel {
	if len(labels) == 0 {
		return labels
	}
	out := make([]VisionLabel, len(labels))
	copy(out, labels)
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Confidence > out[j].Confidence
	})
	top := out[0]
	if !isLesionLabel(top.Label) || top.Confidence >= 0.48 {
		return out
	}
	// Dampen weak lesion scores so a soft normal prior can win.
	for i := range out {
		if isLesionLabel(out[i].Label) && out[i].Confidence < 0.48 {
			out[i].Confidence *= 0.82
		}
	}
	var nsConf float64
	if top.Confidence < 0.35 {
		nsConf = 0.64
	} else {
		nsConf = 0.50 + (0.48-top.Confidence)*0.5
	}
	found := false
	for i := range out {
		if strings.EqualFold(out[i].Label, "normal_skin") {
			if out[i].Confidence < nsConf {
				out[i].Confidence = nsConf
			}
			found = true
			break
		}
	}
	if !found {
		out = append(out, VisionLabel{Label: "normal_skin", SymptomField: "", Confidence: nsConf})
	}
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Confidence > out[j].Confidence
	})
	if len(out) > 3 {
		out = out[:3]
	}
	return out
}

// ClassifySymptomImage runs YOLO if model+bin configured, else multimodal LLM vision.
func ClassifySymptomImage(ctx context.Context, chat *llm.Client, mime string, raw []byte) ([]VisionLabel, string, string, error) {
	if path := os.Getenv("SYMPTOM_YOLO_MODEL"); path != "" {
		// path comes from operator-controlled env, not request input
		if _, err := os.Stat(path); err == nil { // #nosec G703 -- trusted SYMPTOM_YOLO_MODEL env
			labels, yoloRaw, err := classifyWithYOLO(ctx, path, raw)
			if err == nil && len(labels) > 0 {
				return calibrateVisionLabels(labels), "yolo", yoloRaw, nil
			}
		}
	}
	if chat == nil || !chat.Enabled() {
		return nil, "", "", fmt.Errorf("vision classifier unavailable (no YOLO model, LLM disabled)")
	}
	system := `你是皮肤症状图像辅助分类器（不做独立诊断）。从图片判断最可能的皮肤表现。
只输出 JSON：
{"labels":[{"label":"painful_red_rash","confidence":0.8}]}
label 只能是：painful_red_rash, skin_ulcer, itchy_rash, petechial_rash, blood_blister, cyanosis, normal_skin, other。
confidence 0-1；最多 3 条，按置信度降序。不确定给 normal_skin/other。没有明显皮疹/溃疡/血疱时优先 normal_skin。`
	user := "请分类这张皮肤/皮疹/溃疡相关照片。若皮肤外观正常、无明显皮损，请输出 normal_skin。"
	b64 := base64.StdEncoding.EncodeToString(raw)
	out, err := chat.ChatVision(ctx, system, user, mime, b64)
	if err != nil {
		return nil, "llm_vision", out, err
	}
	labels := calibrateVisionLabels(parseVisionLabels(out))
	return labels, "llm_vision", out, nil
}

func parseVisionLabels(raw string) []VisionLabel {
	raw = strings.TrimSpace(raw)
	if i := strings.Index(raw, "{"); i >= 0 {
		if j := strings.LastIndex(raw, "}"); j > i {
			raw = raw[i : j+1]
		}
	}
	var wrap struct {
		Labels []struct {
			Label      string  `json:"label"`
			Confidence float64 `json:"confidence"`
		} `json:"labels"`
	}
	if err := json.Unmarshal([]byte(raw), &wrap); err != nil {
		return nil
	}
	out := make([]VisionLabel, 0, len(wrap.Labels))
	for _, l := range wrap.Labels {
		lab := strings.TrimSpace(l.Label)
		if lab == "" {
			continue
		}
		conf := l.Confidence
		if conf < 0 {
			conf = 0
		}
		if conf > 1 {
			conf = 1
		}
		out = append(out, VisionLabel{
			Label: lab, SymptomField: fieldForLabel(lab), Confidence: conf,
		})
	}
	return out
}

// classifyWithYOLO shells out to SYMPTOM_YOLO_BIN (default: python scripts/vision/yolo_infer.py).
// Expected stdout JSON: {"labels":[{"label":"skin_ulcer","confidence":0.81}]}
func classifyWithYOLO(ctx context.Context, modelPath string, raw []byte) ([]VisionLabel, string, error) {
	bin := strings.TrimSpace(os.Getenv("SYMPTOM_YOLO_BIN"))
	if bin == "" {
		bin = "python3"
	}
	script := strings.TrimSpace(os.Getenv("SYMPTOM_YOLO_SCRIPT"))
	if script == "" {
		script = filepath.Join("scripts", "vision", "yolo_infer.py")
	}
	tmp, err := os.CreateTemp("", "symptom-yolo-*.jpg")
	if err != nil {
		return nil, "", err
	}
	tmpPath := tmp.Name()
	defer func() { _ = os.Remove(tmpPath) }()
	if _, err := tmp.Write(raw); err != nil {
		_ = tmp.Close()
		return nil, "", err
	}
	_ = tmp.Close()

	// bin/script/modelPath come from operator-controlled env vars, not request input
	var cmd *exec.Cmd
	if strings.HasSuffix(bin, ".py") || strings.Contains(bin, "yolo_infer") {
		cmd = exec.CommandContext(ctx, "python3", bin, "--model", modelPath, "--image", tmpPath) // #nosec G204 G702 -- trusted YOLO env
	} else if strings.Contains(bin, "python") {
		cmd = exec.CommandContext(ctx, bin, script, "--model", modelPath, "--image", tmpPath) // #nosec G204 G702 -- trusted YOLO env
	} else {
		cmd = exec.CommandContext(ctx, bin, "--model", modelPath, "--image", tmpPath) // #nosec G204 G702 -- trusted YOLO env
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, stdout.String() + stderr.String(), fmt.Errorf("yolo run: %w: %s", err, truncateRunes(stderr.String(), 400))
	}
	out := strings.TrimSpace(stdout.String())
	labels := parseVisionLabels(out)
	if len(labels) == 0 {
		return nil, out, fmt.Errorf("yolo returned no labels")
	}
	return labels, out, nil
}

// loadVisionLabelsForStatus loads YOLO/VLM labels attached to today's check-in photos.
func loadVisionLabelsForStatus(db *gorm.DB, userID, statusID uint) []VisionLabel {
	if db == nil || (userID == 0 && statusID == 0) {
		return nil
	}
	q := db.Model(&models.SymptomPhoto{})
	switch {
	case statusID > 0 && userID > 0:
		q = q.Where("(status_id = ? OR (status_id = 0 AND user_id = ?))", statusID, userID)
	case statusID > 0:
		q = q.Where("status_id = ?", statusID)
	default:
		q = q.Where("user_id = ?", userID).Order("photo_id DESC").Limit(5)
	}
	var photos []models.SymptomPhoto
	if err := q.Order("photo_id DESC").Limit(8).Find(&photos).Error; err != nil {
		return nil
	}
	var out []VisionLabel
	seen := map[string]struct{}{}
	for _, p := range photos {
		if p.LabelsJSON == "" {
			continue
		}
		var ls []VisionLabel
		if json.Unmarshal([]byte(p.LabelsJSON), &ls) != nil {
			continue
		}
		for _, l := range ls {
			key := l.Label
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			if l.SymptomField == "" {
				l.SymptomField = fieldForLabel(l.Label)
			}
			out = append(out, l)
		}
	}
	return out
}

// ExportVisionDatasetManifest lists active samples for YOLO retrain (ultralytics classify layout).
func ExportVisionDatasetManifest(db *gorm.DB) ([]map[string]any, error) {
	var rows []models.VisionSample
	if err := db.Where("is_active = ?", true).Order("sample_id ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		sf := r.SymptomField
		if sf == "" {
			sf = fieldForLabel(r.ClassLabel)
		}
		out = append(out, map[string]any{
			"sampleId":     r.SampleID,
			"classLabel":   r.ClassLabel,
			"symptomField": sf,
			"split":        r.Split,
			"filePath":     r.FilePath,
			"fileName":     r.FileName,
			"publicUrl":    pkgstorage.PublicURL(r.FilePath),
			"bboxJson":     r.BBoxJSON,
			"source":       r.Source,
			"notes":        r.Notes,
		})
	}
	return out, nil
}

// PromoteSymptomPhotoToSample copies a classified check-in photo into the training library.
func PromoteSymptomPhotoToSample(db *gorm.DB, photoID, createdBy uint, classLabel, split, notes string) (*models.VisionSample, error) {
	var photo models.SymptomPhoto
	if err := db.First(&photo, photoID).Error; err != nil {
		return nil, err
	}
	if classLabel == "" {
		classLabel = photo.TopLabel
	}
	if classLabel == "" {
		return nil, fmt.Errorf("classLabel required")
	}
	if split == "" {
		split = "train"
	}
	sample := &models.VisionSample{
		FileName: photo.FileName, FilePath: photo.FilePath,
		ClassLabel: classLabel, SymptomField: fieldForLabel(classLabel),
		Split: split, Source: "checkin", Notes: notes, CreatedBy: createdBy,
	}
	if err := CreateVisionSample(db, sample); err != nil {
		return nil, err
	}
	return sample, nil
}

// SaveAndClassifySymptomPhoto stores upload and classifies.
func SaveAndClassifySymptomPhoto(db *gorm.DB, chat *llm.Client, userID, statusID uint, key, fileName string) (*models.SymptomPhoto, []VisionLabel, error) {
	r, _, err := pkgstorage.Read(key)
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = r.Close() }()
	raw, err := io.ReadAll(io.LimitReader(r, 8<<20))
	if err != nil {
		return nil, nil, err
	}
	mime := mimeFromExt(strings.ToLower(filepath.Ext(fileName)))
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	labels, classifier, rawResp, classErr := ClassifySymptomImage(ctx, chat, mime, raw)
	lj, _ := json.Marshal(labels)
	top, conf := "", 0.0
	if len(labels) > 0 {
		top, conf = labels[0].Label, labels[0].Confidence
	}
	photo := &models.SymptomPhoto{
		UserID: userID, StatusID: statusID,
		FileName: fileName, FilePath: key, UploadDate: time.Now(),
		LabelsJSON: string(lj), TopLabel: top, Confidence: conf,
		Classifier: classifier, RawResponse: truncateRunes(rawResp, 4000),
	}
	if classErr != nil && photo.Classifier == "" {
		photo.Classifier = "failed"
		photo.RawResponse = classErr.Error()
	}
	if err := db.Create(photo).Error; err != nil {
		return nil, labels, err
	}
	return photo, labels, classErr
}

// ApplyVisionLabelsToSymptoms flips matching general-symptom booleans when confidence >= thr.
func ApplyVisionLabelsToSymptoms(symptoms *AllSymptoms, labels []VisionLabel, thr float64) {
	if symptoms == nil || thr <= 0 {
		thr = 0.55
	}
	for _, l := range labels {
		if l.Confidence < thr || l.SymptomField == "" {
			continue
		}
		switch l.SymptomField {
		case "HasPainfulRedRash":
			symptoms.Generalsymptoms.HasPainfulRedRash = true
		case "HasSkinUlcer":
			symptoms.Generalsymptoms.HasSkinUlcer = true
		case "HasItchyRash":
			symptoms.Generalsymptoms.HasItchyRash = true
		case "HasCongestiveOrPetechialRash":
			symptoms.Generalsymptoms.HasCongestiveOrPetechialRash = true
		case "HasBloodBlisters":
			symptoms.Generalsymptoms.HasBloodBlisters = true
		case "HasCyanosis":
			symptoms.Generalsymptoms.HasCyanosis = true
		}
	}
}

// CreateVisionSampleFromUpload registers a labeled training image.
func CreateVisionSample(db *gorm.DB, sample *models.VisionSample) error {
	if sample == nil || sample.FilePath == "" || sample.ClassLabel == "" {
		return fmt.Errorf("filePath and classLabel required")
	}
	if sample.SymptomField == "" {
		sample.SymptomField = fieldForLabel(sample.ClassLabel)
	}
	if sample.Split == "" {
		sample.Split = "train"
	}
	now := time.Now()
	sample.CreatedAt = now
	sample.UpdatedAt = now
	sample.IsActive = true
	return db.Create(sample).Error
}

// PageVisionSamples lists dataset rows.
func PageVisionSamples(db *gorm.DB, classLabel, split string, page, size int) ([]models.VisionSample, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 200 {
		size = 20
	}
	q := db.Model(&models.VisionSample{}).Where("is_active = ?", true)
	if classLabel != "" {
		q = q.Where("class_label = ?", classLabel)
	}
	if split != "" {
		q = q.Where("split = ?", split)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []models.VisionSample
	err := q.Order("sample_id DESC").Offset((page - 1) * size).Limit(size).Find(&rows).Error
	if err != nil {
		return nil, 0, err
	}
	// Imported / legacy rows often omit symptom_field — fill from class catalog for API/UI.
	for i := range rows {
		if rows[i].SymptomField == "" {
			rows[i].SymptomField = fieldForLabel(rows[i].ClassLabel)
		}
	}
	return rows, total, nil
}

// ExportVisionDatasetSummary returns class counts for YOLO retrain prep.
func ExportVisionDatasetSummary(db *gorm.DB) (map[string]any, error) {
	type row struct {
		ClassLabel string
		Cnt        int64
	}
	var rows []row
	if err := db.Model(&models.VisionSample{}).
		Select("class_label, count(*) as cnt").
		Where("is_active = ?", true).
		Group("class_label").
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	byClass := map[string]int64{}
	var total int64
	for _, r := range rows {
		byClass[r.ClassLabel] = r.Cnt
		total += r.Cnt
	}
	classes := make([]string, 0, len(visionClassCatalog))
	for _, c := range visionClassCatalog {
		classes = append(classes, c.Label)
	}
	return map[string]any{
		"total":      total,
		"byClass":    byClass,
		"classNames": classes,
		"note":       "导出文件列表后可用 ultralytics 分类/检测训练；推理时设置 SYMPTOM_YOLO_MODEL + SYMPTOM_YOLO_BIN",
	}, nil
}
