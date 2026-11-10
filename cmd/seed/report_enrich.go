// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/heathcetide/infectious-disease/internal/models"
)

var reportDiseases = []string{
	"流感", "流感", "流感",
	"新型冠状病毒感染", "新型冠状病毒感染",
	"感染性腹泻", "感染性腹泻",
	"结核病",
	"斑疹伤寒",
	"疟疾（蚊媒传染病）",
	"登革热（蚊媒传染病）",
	"炭疽",
	"鼠疫",
	"森林脑炎（蜱媒传染病）",
	"发热伴血小板减少综合征（蜱媒传染病）",
	"流行性出血热",
}

// seedReportEnrichment boosts unhealthy check-ins, disease_scoring and diagnoses
// so disease-report / accuracy / screen charts look populated.
func seedReportEnrichment(db *gorm.DB) error {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	silent := db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})

	// Normalize legacy name so accuracy chart keys match.
	_ = silent.Model(&models.DailyHealthStatus{}).
		Where("disease_type_name = ?", "流行性热出血").
		Update("disease_type_name", "流行性出血热").Error
	_ = silent.Model(&models.DailyHealthStatus{}).
		Where("diagnosis_disease_type_name = ?", "流行性热出血").
		Update("diagnosis_disease_type_name", "流行性出血热").Error
	_ = silent.Model(&models.DiseaseScoring{}).
		Where("disease_type_name = ?", "流行性热出血").
		Update("disease_type_name", "流行性出血热").Error
	_ = silent.Model(&models.DiagnosisResults{}).
		Where("disease_type = ?", "流行性热出血").
		Update("disease_type", "流行性出血热").Error

	today := time.Date(2026, 11, 8, 0, 0, 0, 0, time.Local)
	dayEnd := today.AddDate(0, 0, 1)

	// Flip more of today's healthy check-ins to unhealthy (~45% unhealthy target).
	var healthyToday []models.DailyHealthStatus
	if err := silent.Where("check_in_date >= ? AND check_in_date < ? AND is_health = ?", today, dayEnd, true).
		Limit(2500).Find(&healthyToday).Error; err != nil {
		return err
	}
	flipN := len(healthyToday) * 28 / 100
	if flipN > len(healthyToday) {
		flipN = len(healthyToday)
	}
	flipped := 0
	for i := 0; i < flipN; i++ {
		rec := healthyToday[i]
		disease := reportDiseases[rng.Intn(len(reportDiseases))]
		h := false
		if err := silent.Model(&rec).Updates(map[string]any{
			"is_health":                   h,
			"disease_type_name":           disease,
			"diagnosis_disease_type_name": disease,
		}).Error; err != nil {
			return err
		}
		flipped++
	}

	// Backfill disease_scoring for recent unhealthy check-ins that lack scores.
	start := today.AddDate(0, 0, -30)
	var unhealthy []models.DailyHealthStatus
	if err := silent.Where("check_in_date >= ? AND check_in_date < ? AND is_health = ? AND disease_type_name <> ''",
		start, dayEnd, false).
		Order("status_id DESC").Limit(8000).Find(&unhealthy).Error; err != nil {
		return err
	}
	scoreBatch := make([]models.DiseaseScoring, 0, 2000)
	addedScores := 0
	for _, st := range unhealthy {
		var exist int64
		_ = silent.Model(&models.DiseaseScoring{}).Where("status_id = ?", st.StatusID).Count(&exist)
		if exist > 0 {
			continue
		}
		day := st.CheckInDate
		// Primary predicted disease + 1-2 competitors for richer charts.
		names := []string{st.DiseaseTypeName}
		for len(names) < 3 {
			cand := reportDiseases[rng.Intn(len(reportDiseases))]
			dup := false
			for _, n := range names {
				if n == cand {
					dup = true
					break
				}
			}
			if !dup {
				names = append(names, cand)
			}
		}
		for i, name := range names {
			score := 70 + rng.Float64()*25
			if i > 0 {
				score = 20 + rng.Float64()*45
			}
			d := day
			scoreBatch = append(scoreBatch, models.DiseaseScoring{
				StatusID:        st.StatusID,
				SymptomTime:     &d,
				DiseaseTypeName: name,
				Score:           score,
			})
			addedScores++
		}
		if len(scoreBatch) >= 1500 {
			if err := silent.CreateInBatches(&scoreBatch, 400).Error; err != nil {
				return err
			}
			scoreBatch = scoreBatch[:0]
		}
	}
	if len(scoreBatch) > 0 {
		if err := silent.CreateInBatches(&scoreBatch, 400).Error; err != nil {
			return err
		}
	}

	// Extra hospital diagnoses across last 45 days for accuracy + hospital source charts.
	var workers []models.User
	_ = silent.Where("role = ?", models.UserRoleRailway).Order("id ASC").Limit(400).Find(&workers)
	var medical models.User
	_ = silent.Where("role = ?", models.UserRoleMedical).First(&medical)
	diagAdded := 0
	for i := 0; i < 220 && i < len(workers); i++ {
		u := workers[rng.Intn(len(workers))]
		disease := reportDiseases[rng.Intn(len(reportDiseases))]
		day := today.AddDate(0, 0, -(rng.Intn(40)))
		admit := day
		dis := day.AddDate(0, 0, 2+rng.Intn(10))
		sub := day.Add(time.Duration(8+rng.Intn(10)) * time.Hour)
		rec := models.DiagnosisResults{
			UserID:                     u.ID,
			SubmissionTime:             &sub,
			SubmissionUserID:           medical.ID,
			SubmissionUserName:         medical.Name,
			SubmissionUserType:         "专职医护人员",
			DiseaseType:                disease,
			DiscoveryMethod:            []string{"主动监测", "健康打卡预警", "门诊就诊"}[rng.Intn(3)],
			DiseaseOutcome:             []string{"治愈", "好转", "未愈"}[rng.Intn(3)],
			HospitalName:               "示例铁路医院",
			AdmissionDate:              &admit,
			DischargeDate:              &dis,
			AdmissionSymptomsAndSigns:  "发热、乏力",
			RegistrationClassification: "临床诊断病例",
		}
		if err := silent.Create(&rec).Error; err != nil {
			continue
		}
		diagAdded++
	}

	var unhealthyToday, scoringN, diagN int64
	_ = silent.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ? AND is_health = ?", today, dayEnd, false).
		Count(&unhealthyToday)
	_ = silent.Model(&models.DiseaseScoring{}).Count(&scoringN)
	_ = silent.Model(&models.DiagnosisResults{}).Count(&diagN)
	fmt.Printf("report enrich: flipped=%d scores+=%d diagnoses+=%d unhealthy_today=%d scoring_total=%d diagnosis_total=%d\n",
		flipped, addedScores, diagAdded, unhealthyToday, scoringN, diagN)
	return nil
}
