// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// DiagnosticAccuracyStrikeVO matches admin accuracy.vue card + chart payload.
type DiagnosticAccuracyStrikeVO struct {
	ConfirmedNumber int64                        `json:"confirmedNumber"`
	DiagnosedNumber int64                        `json:"diagnosedNumber"`
	List            []DiagnosticAccuracyDiseaseRow `json:"list"`
}

// DiagnosticAccuracyDiseaseRow is per-disease predict vs confirmed counts.
type DiagnosticAccuracyDiseaseRow struct {
	DiseaseTypeName           string `json:"diseaseTypeName"`
	PredictDiagnoseNumber     int64  `json:"predictDiagnoseNumber"`
	ActuallyDiagnoseNumber    int64  `json:"actuallyDiagnoseNumber"`
}

// defaultAccuracyDiseaseNames matches admin chart labels.
var defaultAccuracyDiseaseNames = []string{
	"新型冠状病毒感染",
	"流感",
	"鼠疫",
	"感染性腹泻",
	"炭疽",
	"结核病",
	"登革热（蚊媒传染病）",
	"疟疾（蚊媒传染病）",
	"森林脑炎（蜱媒传染病）",
	"发热伴血小板减少综合征（蜱媒传染病）",
	"斑疹伤寒",
	"流行性热出血",
}

// DiagnosticAccuracyStrike computes strike stats between startDate and endDate (yyyy-MM-dd).
func DiagnosticAccuracyStrike(db *gorm.DB, startDate, endDate string) (DiagnosticAccuracyStrikeVO, error) {
	out := DiagnosticAccuracyStrikeVO{List: make([]DiagnosticAccuracyDiseaseRow, 0, len(defaultAccuracyDiseaseNames))}
	if db == nil {
		for _, name := range defaultAccuracyDiseaseNames {
			out.List = append(out.List, DiagnosticAccuracyDiseaseRow{DiseaseTypeName: name})
		}
		return out, nil
	}
	start, end, err := parseDateRange(startDate, endDate)
	if err != nil {
		for _, name := range defaultAccuracyDiseaseNames {
			out.List = append(out.List, DiagnosticAccuracyDiseaseRow{DiseaseTypeName: name})
		}
		return out, nil
	}

	threshold := defaultPredictorRange(db)

	var confirmed int64
	_ = db.Model(&models.DiagnosisResults{}).
		Where("submission_time >= ? AND submission_time <= ?", start, end).
		Count(&confirmed).Error
	out.ConfirmedNumber = confirmed

	predictByName := map[string]int64{}
	type dsRow struct {
		Name string
		Cnt  int64
	}
	var dsRows []dsRow
	_ = db.Model(&models.DiseaseScoring{}).
		Select("disease_type_name as name, count(*) as cnt").
		Where("symptom_time >= ? AND symptom_time <= ? AND score >= ?", start, end, threshold).
		Group("disease_type_name").
		Scan(&dsRows).Error
	var diagnosed int64
	for _, r := range dsRows {
		predictByName[r.Name] = r.Cnt
		diagnosed += r.Cnt
	}
	out.DiagnosedNumber = diagnosed

	actualByName := map[string]int64{}
	type drRow struct {
		Name string
		Cnt  int64
	}
	var drRows []drRow
	_ = db.Model(&models.DiagnosisResults{}).
		Select("disease_type as name, count(*) as cnt").
		Where("submission_time >= ? AND submission_time <= ?", start, end).
		Group("disease_type").
		Scan(&drRows).Error
	for _, r := range drRows {
		actualByName[r.Name] = r.Cnt
	}

	for _, name := range defaultAccuracyDiseaseNames {
		out.List = append(out.List, DiagnosticAccuracyDiseaseRow{
			DiseaseTypeName:        name,
			PredictDiagnoseNumber:  predictByName[name],
			ActuallyDiagnoseNumber: actualByName[name],
		})
	}
	return out, nil
}

func parseDateRange(startDate, endDate string) (time.Time, time.Time, error) {
	start, err := time.ParseInLocation("2006-01-02", startDate, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	end, err := time.ParseInLocation("2006-01-02", endDate, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	end = end.Add(24*time.Hour - time.Nanosecond)
	return start, end, nil
}

func defaultPredictorRange(db *gorm.DB) float64 {
	s, err := GetSystemSettings(db)
	if err != nil || s.PredictorRange == nil {
		return 1.0
	}
	return *s.PredictorRange
}

// SetPredictorRange stores the scoring threshold for predicted diagnoses.
func SetPredictorRange(db *gorm.DB, rangeVal float64) error {
	s, err := GetSystemSettings(db)
	if err != nil {
		return err
	}
	s.PredictorRange = &rangeVal
	return db.Model(s).Update("predictor_range", rangeVal).Error
}
