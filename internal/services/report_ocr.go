// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/llm"
	"github.com/heathcetide/infectious-disease/internal/models"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
)

// LabMetrics is the structured OCR output for blood / imaging reports.
type LabMetrics struct {
	ReportKind string             `json:"reportKind"` // blood | imaging | other
	Summary    string             `json:"summary"`
	Findings   []string           `json:"findings,omitempty"`
	Indicators map[string]float64 `json:"indicators,omitempty"` // e.g. wbc, crp, esr, neut%, hb, plt
	Flags      []string           `json:"flags,omitempty"`      // elevated_wbc, lung_opacity, ...
}

// ShouldOCRReportType is true for blood / imaging reports in the P3 OCR scope.
func ShouldOCRReportType(reportType string) bool {
	t := reportType
	return strings.Contains(t, "血液") || strings.Contains(t, "血常规") ||
		strings.Contains(t, "影像") || strings.Contains(t, "B超") ||
		strings.EqualFold(t, "blood") || strings.EqualFold(t, "imaging")
}

func shouldOCRReportType(reportType string) bool { return ShouldOCRReportType(reportType) }

// ProcessDiagnosisReportOCR runs OCR → structured metrics for one uploaded report.
func ProcessDiagnosisReportOCR(db *gorm.DB, chat *llm.Client, reportID uint) error {
	if db == nil || reportID == 0 {
		return fmt.Errorf("invalid ocr args")
	}
	var rep models.DiagnosisReport
	if err := db.First(&rep, reportID).Error; err != nil {
		return err
	}
	if !shouldOCRReportType(rep.ReportType) {
		_ = db.Model(&rep).Updates(map[string]any{
			"ocr_status": "skipped",
			"ocr_error":  "report type not in OCR scope",
		}).Error
		return nil
	}
	_ = db.Model(&rep).Updates(map[string]any{"ocr_status": "pending", "ocr_error": ""}).Error

	raw, mime, err := readReportBytes(rep.FilePath, rep.FileName)
	if err != nil {
		_ = db.Model(&rep).Updates(map[string]any{
			"ocr_status": "failed",
			"ocr_error":  truncateRunes(err.Error(), 200),
		}).Error
		return err
	}

	metrics, ocrText, err := ocrReportContent(context.Background(), chat, rep.ReportType, mime, raw)
	if err != nil {
		_ = db.Model(&rep).Updates(map[string]any{
			"ocr_status": "failed",
			"ocr_error":  truncateRunes(err.Error(), 200),
		}).Error
		return err
	}
	mj, _ := json.Marshal(metrics)
	return db.Model(&rep).Updates(map[string]any{
		"ocr_status":   "done",
		"ocr_text":     truncateRunes(ocrText, 4000),
		"metrics_json": string(mj),
		"ocr_error":    "",
	}).Error
}

func readReportBytes(key, fileName string) ([]byte, string, error) {
	r, _, err := pkgstorage.Read(key)
	if err != nil {
		return nil, "", err
	}
	defer func() { _ = r.Close() }()
	raw, err := io.ReadAll(io.LimitReader(r, 8<<20))
	if err != nil {
		return nil, "", err
	}
	ext := strings.ToLower(filepath.Ext(fileName))
	if ext == "" {
		ext = strings.ToLower(filepath.Ext(key))
	}
	mime := mimeFromExt(ext)
	return raw, mime, nil
}

func mimeFromExt(ext string) string {
	switch ext {
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".bmp":
		return "image/bmp"
	case ".pdf":
		return "application/pdf"
	default:
		return "image/jpeg"
	}
}

func ocrReportContent(ctx context.Context, chat *llm.Client, reportType, mime string, raw []byte) (LabMetrics, string, error) {
	kind := "other"
	if strings.Contains(reportType, "血液") || strings.Contains(reportType, "血常规") {
		kind = "blood"
	} else if strings.Contains(reportType, "影像") || strings.Contains(reportType, "B超") {
		kind = "imaging"
	}

	// Textual PDF fallback: pull printable runs (no YOLO / no PDF lib).
	if mime == "application/pdf" {
		text := extractPrintablePDFText(raw)
		if strings.TrimSpace(text) == "" {
			return LabMetrics{ReportKind: kind, Summary: "PDF 无可提取文本"}, "", fmt.Errorf("pdf has no extractable text; upload image photo of report")
		}
		metrics := parseLabMetricsFromText(kind, text)
		if chat != nil && chat.Enabled() {
			enriched, err := structureMetricsWithLLM(ctx, chat, kind, text)
			if err == nil {
				metrics = enriched
			}
		}
		return metrics, text, nil
	}

	if chat == nil || !chat.Enabled() {
		return LabMetrics{}, "", fmt.Errorf("llm not configured for vision OCR")
	}

	system := `你是医学报告 OCR 助手。从检验/影像报告图片中提取结构化指标。
只输出 JSON，不要 Markdown：
{"reportKind":"blood|imaging","summary":"一句话中文摘要","findings":["..."],"indicators":{"wbc":6.5,"crp":12,"esr":20,"neut_pct":70,"hb":130,"plt":200,"lymph_pct":25},"flags":["elevated_wbc","elevated_crp","lung_opacity"]}
规则：
1. indicators 只填能读到的数值（白细胞用 wbc，单位 10^9/L；CRP mg/L；ESR mm/h；血红蛋白 hb g/L；血小板 plt 10^9/L）。
2. flags 用英文短码：elevated_wbc, low_wbc, elevated_crp, elevated_esr, anemia, thrombocytopenia, lung_opacity, consolidation, pleural_effusion, lymphadenopathy。
3. 读不到的指标不要编造。`
	user := fmt.Sprintf("报告类型：%s。请 OCR 并结构化。", reportType)
	b64 := base64.StdEncoding.EncodeToString(raw)
	out, err := chat.ChatVision(ctx, system, user, mime, b64)
	if err != nil {
		return LabMetrics{}, "", err
	}
	metrics, parseErr := parseLabMetricsJSON(out)
	if parseErr != nil {
		// Keep raw OCR text even if JSON parse fails.
		fallback := parseLabMetricsFromText(kind, out)
		fallback.Summary = truncateRunes(out, 120)
		return fallback, out, nil
	}
	if metrics.ReportKind == "" {
		metrics.ReportKind = kind
	}
	return metrics, out, nil
}

func structureMetricsWithLLM(ctx context.Context, chat *llm.Client, kind, text string) (LabMetrics, error) {
	system := `从检验/影像报告文本提取 JSON：
{"reportKind":"blood|imaging","summary":"...","findings":[],"indicators":{"wbc":0},"flags":[]}
只输出 JSON。`
	user := "报告类型倾向：" + kind + "\n文本：\n" + truncateRunes(text, 3500)
	raw, err := chat.Chat(ctx, system, user)
	if err != nil {
		return LabMetrics{}, err
	}
	return parseLabMetricsJSON(raw)
}

func parseLabMetricsJSON(raw string) (LabMetrics, error) {
	raw = strings.TrimSpace(raw)
	if i := strings.Index(raw, "{"); i >= 0 {
		if j := strings.LastIndex(raw, "}"); j > i {
			raw = raw[i : j+1]
		}
	}
	var m LabMetrics
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		return LabMetrics{}, err
	}
	if m.Indicators == nil {
		m.Indicators = map[string]float64{}
	}
	return m, nil
}

func parseLabMetricsFromText(kind, text string) LabMetrics {
	m := LabMetrics{ReportKind: kind, Indicators: map[string]float64{}, Summary: truncateRunes(text, 80)}
	lower := strings.ToLower(text)
	if strings.Contains(text, "白细胞") || strings.Contains(lower, "wbc") {
		m.Flags = append(m.Flags, "mentions_wbc")
	}
	if strings.Contains(text, "C反应蛋白") || strings.Contains(lower, "crp") {
		m.Flags = append(m.Flags, "mentions_crp")
	}
	if strings.Contains(text, "阴影") || strings.Contains(text, "渗出") || strings.Contains(text, "实变") {
		m.Flags = append(m.Flags, "lung_opacity")
	}
	return m
}

func extractPrintablePDFText(raw []byte) string {
	var b strings.Builder
	run := 0
	for _, c := range raw {
		if c >= 32 && c < 127 || c >= 0xC0 {
			b.WriteByte(c)
			run++
		} else if run > 0 {
			b.WriteByte(' ')
			run = 0
		}
	}
	s := b.String()
	// Keep Chinese-ish by also scanning runes from UTF-8 attempts — best-effort.
	var out strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.Is(unicode.Han, r) ||
			r == '.' || r == '%' || r == '/' || r == '-' || r == ' ' || r == ':' || r == '：' {
			out.WriteRune(r)
		}
	}
	return strings.Join(strings.Fields(out.String()), " ")
}

// loadRecentLabMetrics returns OCR metrics for the user from recent diagnosis reports.
func loadRecentLabMetrics(db *gorm.DB, userID uint) []LabMetrics {
	if db == nil || userID == 0 {
		return nil
	}
	since := time.Now().AddDate(0, 0, -45)
	var reports []models.DiagnosisReport
	err := db.Table("diagnosis_reports AS r").
		Select("r.*").
		Joins("JOIN diagnosis_results d ON d.diagnosis_results_id = r.diagnosis_results_id").
		Where("d.user_id = ? AND r.ocr_status = ? AND r.upload_date >= ?", userID, "done", since).
		Order("r.upload_date DESC").
		Limit(8).
		Find(&reports).Error
	if err != nil {
		return nil
	}
	out := make([]LabMetrics, 0, len(reports))
	for _, r := range reports {
		if strings.TrimSpace(r.MetricsJSON) == "" {
			continue
		}
		var m LabMetrics
		if json.Unmarshal([]byte(r.MetricsJSON), &m) == nil {
			out = append(out, m)
		}
	}
	return out
}

func formatLabMetricsForLLM(metrics []LabMetrics) string {
	if len(metrics) == 0 {
		return "近期检验/影像 OCR：无\n"
	}
	var b strings.Builder
	b.WriteString("近期检验/影像 OCR（参与排序）：\n")
	for i, m := range metrics {
		if i >= 5 {
			break
		}
		fmt.Fprintf(&b, "- [%s] %s", m.ReportKind, m.Summary)
		if len(m.Indicators) > 0 {
			parts := make([]string, 0, len(m.Indicators))
			for k, v := range m.Indicators {
				parts = append(parts, fmt.Sprintf("%s=%.2f", k, v))
			}
			b.WriteString("；指标 " + strings.Join(parts, ", "))
		}
		if len(m.Flags) > 0 {
			b.WriteString("；flags " + strings.Join(m.Flags, ","))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func labOCRBonus(disease string, metrics []LabMetrics) (float64, string) {
	if len(metrics) == 0 {
		return 0, ""
	}
	var bonus float64
	var reasons []string
	add := func(pts float64, why string) {
		if pts == 0 {
			return
		}
		bonus += pts
		reasons = append(reasons, why)
	}
	hasFlag := func(f string) bool {
		for _, m := range metrics {
			for _, x := range m.Flags {
				if strings.EqualFold(x, f) {
					return true
				}
			}
			if v, ok := m.Indicators[strings.ToLower(f)]; ok && v > 0 {
				_ = v
			}
		}
		return false
	}
	wbcHigh, crpHigh, lung := false, false, false
	for _, m := range metrics {
		if wbc, ok := m.Indicators["wbc"]; ok && wbc >= 10 {
			wbcHigh = true
		}
		if crp, ok := m.Indicators["crp"]; ok && crp >= 10 {
			crpHigh = true
		}
		for _, f := range m.Flags {
			lf := strings.ToLower(f)
			if lf == "elevated_wbc" {
				wbcHigh = true
			}
			if lf == "elevated_crp" {
				crpHigh = true
			}
			if lf == "lung_opacity" || lf == "consolidation" || lf == "pleural_effusion" {
				lung = true
			}
		}
	}
	_ = hasFlag
	switch {
	case strings.Contains(disease, "流感"), strings.Contains(disease, "新型冠状"), strings.Contains(disease, "结核"):
		if lung {
			add(6, "影像提示肺部异常")
		}
		if wbcHigh || crpHigh {
			add(3, "血常规/CRP 炎症指标升高")
		}
	case strings.Contains(disease, "感染性腹泻"), strings.Contains(disease, "伤寒"):
		if wbcHigh || crpHigh {
			add(2, "炎症指标支持感染")
		}
	case strings.Contains(disease, "鼠疫"), strings.Contains(disease, "炭疽"):
		if wbcHigh {
			add(2, "白细胞升高")
		}
	}
	why := strings.Join(reasons, "；")
	if len([]rune(why)) > 60 {
		why = string([]rune(why)[:60]) + "…"
	}
	return bonus, why
}

// EnqueueDiagnosisReportOCR fires OCR in background (best-effort).
func EnqueueDiagnosisReportOCR(db *gorm.DB, chat *llm.Client, reportID uint) {
	if reportID == 0 {
		return
	}
	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("ocr panic report=%d: %v", reportID, rec)
			}
		}()
		if err := ProcessDiagnosisReportOCR(db, chat, reportID); err != nil {
			log.Printf("ocr report=%d: %v", reportID, err)
		}
	}()
}
