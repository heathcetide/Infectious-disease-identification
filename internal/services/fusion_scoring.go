// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// ScoringContext aggregates multi-source signals for P2/P3 fusion priors.
type ScoringContext struct {
	Month            time.Month
	AltitudeM        int
	Province         string
	City             string
	County           string
	LocationName     string
	AvgPressureKPa   float64
	MaxTempC         float64
	MinTempC         float64
	HasEnv           bool
	HasHypertension  bool
	HasDiabetes      bool
	HasCOPD          bool
	HasAsthma        bool
	HasStroke        bool
	HasCoronary      bool
	HasMedicalHist   bool
	LabMetrics       []LabMetrics
	VisionLabels     []VisionLabel // P4 skin/rash YOLO|VLM assist (not standalone diagnosis)
	Notes            []string
}

func loadScoringContext(db *gorm.DB, userID, statusID uint) ScoringContext {
	ctx := ScoringContext{Month: time.Now().Month()}
	var status models.DailyHealthStatus
	if err := db.First(&status, statusID).Error; err == nil {
		ctx.Province = status.Province
		ctx.City = status.City
		ctx.County = status.County
		ctx.LocationName = status.LocationName
		ctx.AltitudeM = estimateAltitudeM(status)
	}
	ym := time.Now().Format("2006-01")
	var env models.WorkEnvironmentInfo
	if err := db.Where("user_id = ? AND year_month = ?", userID, ym).First(&env).Error; err == nil {
		ctx.HasEnv = true
		ctx.AvgPressureKPa = env.AvgMonthlyPressure
		ctx.MaxTempC = env.MaxTemperature
		ctx.MinTempC = env.MinTemperature
	}
	var profile models.EmployeeProfile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err == nil {
		ctx.HasHypertension = profile.HasHypertension
		ctx.HasDiabetes = profile.HasDiabetes
		ctx.HasCOPD = profile.HasCOPD
		ctx.HasAsthma = profile.HasAsthma
		ctx.HasStroke = profile.HasStroke
		ctx.HasCoronary = profile.HasCoronaryHeartDisease
		ctx.HasMedicalHist = profile.HasMedicalHistory ||
			ctx.HasHypertension || ctx.HasDiabetes || ctx.HasCOPD || ctx.HasAsthma ||
			ctx.HasStroke || ctx.HasCoronary
	}
	ctx.LabMetrics = loadRecentLabMetrics(db, userID)
	ctx.VisionLabels = loadVisionLabelsForStatus(db, userID, statusID)
	return ctx
}

func estimateAltitudeM(st models.DailyHealthStatus) int {
	name := st.Province + st.City + st.County + st.LocationName
	switch {
	case strings.Contains(name, "唐古拉"), strings.Contains(name, "风火山"):
		return 5000
	case strings.Contains(name, "可可西里"), strings.Contains(name, "五道梁"), strings.Contains(name, "沱沱河"),
		strings.Contains(name, "安多"), strings.Contains(name, "那曲"), strings.Contains(name, "当雄"):
		return 4500
	case strings.Contains(name, "格尔木"), strings.Contains(name, "昆仑"), strings.Contains(name, "羊八井"):
		return 3800
	case strings.Contains(st.Province, "西藏"):
		return 3600
	case strings.Contains(st.Province, "青海"):
		return 3000
	case strings.Contains(st.Province, "甘肃") && (strings.Contains(st.City, "甘南") || strings.Contains(st.City, "张掖") || strings.Contains(st.City, "酒泉")):
		return 2200
	case strings.Contains(st.Province, "四川") && (strings.Contains(st.City, "阿坝") || strings.Contains(st.City, "甘孜")):
		return 3000
	case strings.Contains(st.Province, "云南") && strings.Contains(st.City, "迪庆"):
		return 3200
	default:
		return 800
	}
}

// contextBonus returns additive prior points for a disease under multi-source rules.
func contextBonus(disease string, ctx ScoringContext, active map[string]struct{}) (float64, string) {
	var bonus float64
	var reasons []string
	m := ctx.Month
	summer := m >= time.May && m <= time.October
	winter := m == time.November || m == time.December || m == time.January || m == time.February

	add := func(pts float64, why string) {
		if pts == 0 {
			return
		}
		bonus += pts
		reasons = append(reasons, why)
	}

	switch {
	case strings.Contains(disease, "登革热"), strings.Contains(disease, "疟疾"):
		if summer {
			add(8, "夏秋蚊媒季节先验")
		} else {
			add(-4, "非蚊媒活跃季")
		}
		if ctx.MinTempC > 0 && ctx.MinTempC < 12 {
			add(-3, "低温不利于蚊媒")
		}
		if ctx.AltitudeM > 2800 {
			add(-6, "高海拔蚊媒风险偏低")
		}
	case strings.Contains(disease, "森林脑炎"), strings.Contains(disease, "发热伴血小板减少"):
		if m >= time.April && m <= time.September {
			add(7, "蜱媒活跃季先验")
		}
		if _, ok := active["HasTickBite"]; ok {
			add(6, "有蜱叮咬暴露")
		}
		if ctx.AltitudeM > 2000 && ctx.AltitudeM < 4200 {
			add(3, "中高海拔林区蜱媒先验")
		}
	case strings.Contains(disease, "流感"), strings.Contains(disease, "新型冠状"):
		if winter {
			add(6, "冬春季呼吸道传染病先验")
		}
		if ctx.AvgPressureKPa > 0 && ctx.AvgPressureKPa < 80 {
			add(3, "低压环境呼吸道负担")
		}
	case strings.Contains(disease, "感染性腹泻"):
		if summer {
			add(4, "夏秋肠道传染病先验")
		}
		if ctx.MaxTempC >= 28 {
			add(3, "高温利于肠道传播")
		}
	case strings.Contains(disease, "鼠疫"):
		if m >= time.May && m <= time.October {
			add(5, "鼠疫自然疫源活跃季")
		}
		if ctx.AltitudeM > 3000 {
			add(4, "高原疫源地海拔先验")
		}
		if _, wild := active["HasWildActivityHistory"]; wild {
			add(5, "野外活动暴露")
		}
	case strings.Contains(disease, "炭疽"):
		if m >= time.May && m <= time.September {
			add(3, "炭疽夏秋先验")
		}
	case strings.Contains(disease, "结核"):
		if ctx.HasCOPD || ctx.HasAsthma {
			add(4, "慢阻肺/哮喘基础病史")
		}
		if ctx.AltitudeM > 3500 {
			add(2, "高海拔慢性呼吸道负担")
		}
	case strings.Contains(disease, "流行性出血热"):
		if m >= time.October || m <= time.January {
			add(4, "出血热秋冬高峰先验")
		}
	case strings.Contains(disease, "斑疹伤寒"):
		if winter {
			add(3, "冬春斑疹伤寒先验")
		}
	}

	// Plateau reaction / respiratory burden for any disease when altitude high + low pressure.
	if ctx.AltitudeM >= 4000 {
		add(1.5, "极高海拔综合风险")
	}
	if ctx.HasEnv && ctx.AvgPressureKPa > 0 && ctx.AvgPressureKPa < 70 {
		add(2, "本月气压偏低")
	}
	if ctx.HasMedicalHist && (strings.Contains(disease, "流感") || strings.Contains(disease, "新型冠状") || strings.Contains(disease, "结核")) {
		if ctx.HasDiabetes || ctx.HasCoronary || ctx.HasStroke {
			add(2, "慢病基础增加呼吸道重症风险")
		}
	}

	why := strings.Join(reasons, "；")
	if len([]rune(why)) > 80 {
		why = string([]rune(why)[:80]) + "…"
	}
	return bonus, why
}

func applyContextFusion(scores []models.DiseaseScoring, ctx ScoringContext, active map[string]struct{}) map[string]string {
	notes := map[string]string{}
	for i := range scores {
		bonus, why := contextBonus(scores[i].DiseaseTypeName, ctx, active)
		labBonus, labWhy := labOCRBonus(scores[i].DiseaseTypeName, ctx.LabMetrics)
		if labBonus != 0 {
			bonus += labBonus
			if labWhy != "" {
				if why != "" {
					why = why + "；" + labWhy
				} else {
					why = labWhy
				}
			}
		}
		if bonus != 0 {
			scores[i].Score += bonus
			if scores[i].Score < 0 {
				scores[i].Score = 0
			}
			if scores[i].Score > 100 {
				scores[i].Score = 100
			}
			notes[scores[i].DiseaseTypeName] = why
		}
	}
	return notes
}

func formatContextForLLM(ctx ScoringContext) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("当前月份：%d 月\n", int(ctx.Month)))
	b.WriteString(fmt.Sprintf("打卡驻地：%s %s %s %s\n", ctx.Province, ctx.City, ctx.County, ctx.LocationName))
	b.WriteString(fmt.Sprintf("估算海拔：约 %d 米\n", ctx.AltitudeM))
	if ctx.HasEnv {
		b.WriteString(fmt.Sprintf("本月平均气压：%.1f kPa；最高温 %.1f℃；最低温 %.1f℃\n",
			ctx.AvgPressureKPa, ctx.MaxTempC, ctx.MinTempC))
	} else {
		b.WriteString("本月气压气温：未登记\n")
	}
	hist := []string{}
	if ctx.HasHypertension {
		hist = append(hist, "高血压")
	}
	if ctx.HasDiabetes {
		hist = append(hist, "糖尿病")
	}
	if ctx.HasCOPD {
		hist = append(hist, "慢阻肺")
	}
	if ctx.HasAsthma {
		hist = append(hist, "哮喘")
	}
	if ctx.HasCoronary {
		hist = append(hist, "冠心病")
	}
	if ctx.HasStroke {
		hist = append(hist, "脑卒中史")
	}
	if len(hist) == 0 {
		b.WriteString("既往病史：无明显记录\n")
	} else {
		b.WriteString("既往病史：" + strings.Join(hist, "、") + "\n")
	}
	b.WriteString(formatLabMetricsForLLM(ctx.LabMetrics))
	b.WriteString(formatVisionLabelsForLLM(ctx.VisionLabels))
	b.WriteString("请结合季节/海拔/气压气温/暴露/病史/检验影像 OCR/皮肤视觉辅助特征调整分数，并在 reason 中点明用到的先验；视觉标签仅为辅助，不可单独作为诊断依据。\n")
	return b.String()
}

func formatVisionLabelsForLLM(labels []VisionLabel) string {
	if len(labels) == 0 {
		return "皮肤视觉辅助：无\n"
	}
	parts := make([]string, 0, len(labels))
	for _, l := range labels {
		if l.Label == "" || l.Label == "normal_skin" || l.Label == "other" {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s(%.2f→%s)", l.Label, l.Confidence, l.SymptomField))
	}
	if len(parts) == 0 {
		return "皮肤视觉辅助：未检出明确皮疹/溃疡特征\n"
	}
	return "皮肤视觉辅助（YOLO/VLM，辅助特征）：" + strings.Join(parts, "；") + "\n"
}
