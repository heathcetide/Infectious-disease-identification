// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/llm"
	"github.com/heathcetide/infectious-disease/internal/models"
)

// CheckIn creates today's DailyHealthStatus for the user.
func CheckIn(db *gorm.DB, userID uint, dto CheckInDTO) (*models.DailyHealthStatus, error) {
	now := time.Now()
	day := dateOnly(now)
	var existing models.DailyHealthStatus
	err := db.Where("user_id = ? AND check_in_date = ?", userID, day).First(&existing).Error
	if err == nil {
		return &existing, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	health := dto.Health
	status := models.DailyHealthStatus{
		UserID:       userID,
		CheckInDate:  day,
		CheckInTime:  now,
		IsHealth:     &health,
		Province:     dto.Province,
		City:         dto.City,
		County:       dto.County,
		LocationName: dto.LocationName,
		Longitude:    dto.Longitude,
		Latitude:     dto.Latitude,
	}
	if err := db.Create(&status).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

// hasMonthWorkEnvironment reports whether the user filled work-environment for date's month.
func hasMonthWorkEnvironment(db *gorm.DB, userID uint, date time.Time) bool {
	var n int64
	_ = db.Model(&models.WorkEnvironmentInfo{}).
		Where("user_id = ? AND year_month = ?", userID, date.Format("2006-01")).
		Count(&n).Error
	return n > 0
}

// GetDaily returns check-in and today's disease scores for a user on a date.
func GetDaily(db *gorm.DB, userID uint, date time.Time) (*DailyCheckInVO, error) {
	day := dateOnly(date)
	monthSigned := hasMonthWorkEnvironment(db, userID, day)

	var status models.DailyHealthStatus
	err := db.Where("user_id = ? AND check_in_date = ?", userID, day).First(&status).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &DailyCheckInVO{
			HasSignedToday: false,
			HasMonthSigned: monthSigned,
		}, nil
	}
	if err != nil {
		return nil, err
	}

	mode, _ := GetAIModeDTO(db)
	items := make([]ScoredDiseaseItem, 0)
	var hybrid *HybridMeta
	if mode.IsAI {
		var aiScores []models.AIDiseaseScoring
		if err := db.Where("status_id = ? AND symptom_time = ?", status.StatusID, day).
			Find(&aiScores).Error; err != nil {
			return nil, err
		}
		for _, a := range aiScores {
			var ev []string
			if a.EvidenceJSON != "" {
				_ = json.Unmarshal([]byte(a.EvidenceJSON), &ev)
			}
			items = append(items, ScoredDiseaseItem{
				ScoringID: a.ScoringID, StatusID: a.StatusID, SymptomTime: a.SymptomTime,
				DiseaseTypeID: a.DiseaseTypeID, DiseaseTypeName: a.DiseaseTypeName,
				Score: a.Score, Reason: a.Reason, Evidence: ev,
			})
		}
		disagree := status.ScoringDisagreement != nil && *status.ScoringDisagreement
		hybrid = &HybridMeta{
			IsAI:              true,
			Disagreement:      disagree,
			NeedsHumanConfirm: false, // 职工端不要求确认；分歧仅供疾控复核
			Confirmed:         status.ScoringConfirmedAt != nil,
			Message:           employeeHybridMessage(disagree),
		}
	}
	if len(items) == 0 {
		var scores []models.DiseaseScoring
		if err := db.Where("status_id = ? AND symptom_time = ?", status.StatusID, day).
			Find(&scores).Error; err != nil {
			return nil, err
		}
		for _, s := range scores {
			items = append(items, ScoredDiseaseItem{
				ScoringID: s.ScoringID, StatusID: s.StatusID, SymptomTime: s.SymptomTime,
				DiseaseTypeID: s.DiseaseTypeID, DiseaseTypeName: s.DiseaseTypeName, Score: s.Score,
			})
		}
	}

	return &DailyCheckInVO{
		HasSignedToday:     true,
		HasMonthSigned:     monthSigned,
		Dailyhealthstatus:  &status,
		DiseasescoringList: items,
		Hybrid:             hybrid,
	}, nil
}

func employeeHybridMessage(disagree bool) string {
	_ = disagree
	return "已根据症状完成风险评估"
}

func statusHybridMessage(disagree, confirmed bool) string {
	if !disagree {
		return "AI 与专家权重结论一致"
	}
	if confirmed {
		return "疾控已复核（AI 与专家不一致）"
	}
	return "AI 与专家权重结论不一致，请疾控人员复核"
}

// GetMonth lists days in yyyy-MM with check-in, diagnosis, and lab flags.
func GetMonth(db *gorm.DB, userID uint, yearMonth string) ([]MonthlyHealthDay, error) {
	if len(yearMonth) != 7 {
		return nil, ErrInvalidDate
	}

	start, end, err := monthDateRange(yearMonth)
	if err != nil {
		return nil, err
	}
	var checkins []models.DailyHealthStatus
	if err := db.Where("user_id = ? AND check_in_date >= ? AND check_in_date < ?", userID, start, end).
		Find(&checkins).Error; err != nil {
		return nil, err
	}

	byDate := make(map[string]*MonthlyHealthDay)
	for _, c := range checkins {
		d := c.CheckInDate.Format("2006-01-02")
		byDate[d] = &MonthlyHealthDay{Date: d, Health: c.IsHealth}
	}

	var diagnoses []models.DiagnosisResults
	_ = db.Where("user_id = ? AND submission_time >= ? AND submission_time < ?", userID, start, end).
		Find(&diagnoses).Error
	for _, d := range diagnoses {
		if d.SubmissionTime == nil {
			continue
		}
		ds := d.SubmissionTime.Format("2006-01-02")
		day := byDate[ds]
		if day == nil {
			day = &MonthlyHealthDay{Date: ds}
			byDate[ds] = day
		}
		day.Diagnose = true
	}

	var labs []models.LabTestReport
	_ = db.Where("user_id = ? AND upload_date >= ? AND upload_date < ?", userID, start, end).
		Find(&labs).Error
	for _, l := range labs {
		if l.UploadDate == nil {
			continue
		}
		ds := l.UploadDate.Format("2006-01-02")
		day := byDate[ds]
		if day == nil {
			day = &MonthlyHealthDay{Date: ds}
			byDate[ds] = day
		}
		day.Examine = true
	}

	out := make([]MonthlyHealthDay, 0, len(byDate))
	for _, v := range byDate {
		out = append(out, *v)
	}
	return out, nil
}

// SaveSymptomsAndScore attaches symptoms to today's status, saves tables, and runs scoring.
// Expert mode → weighted disease_scoring. AI mode → LLM (+ expert contrast / disagreement).
func SaveSymptomsAndScore(db *gorm.DB, chat *llm.Client, userID uint, symptoms AllSymptoms) (*HybridScoringResult, error) {
	status, err := todayStatus(db, userID)
	if err != nil {
		return nil, err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		sid := status.StatusID
		general := symptoms.Generalsymptoms.ToModel()
		general.StatusID = sid
		if err = upsertSymptomRow(tx, &general, sid); err != nil {
			return err
		}
		symptoms.Respiratorysymptoms.StatusID = sid
		if err = upsertSymptomRow(tx, &symptoms.Respiratorysymptoms, sid); err != nil {
			return err
		}
		digestive := symptoms.Digestivesymptoms.ToModel()
		digestive.StatusID = sid
		if err = upsertSymptomRow(tx, &digestive, sid); err != nil {
			return err
		}
		symptoms.Circulatorysymptoms.StatusID = sid
		if err = upsertSymptomRow(tx, &symptoms.Circulatorysymptoms, sid); err != nil {
			return err
		}
		symptoms.NeurologicalSymptoms.StatusID = sid
		if err = upsertSymptomRow(tx, &symptoms.NeurologicalSymptoms, sid); err != nil {
			return err
		}
		symptoms.LocalSymptoms.StatusID = sid
		if err = upsertSymptomRow(tx, &symptoms.LocalSymptoms, sid); err != nil {
			return err
		}
		other := symptoms.OtherSymptoms
		if other.WorkLifeStress == "" {
			other.WorkLifeStress = "小"
		}
		other.StatusID = sid
		if err = upsertSymptomRow(tx, &other, sid); err != nil {
			return err
		}
		risk := symptoms.RiskFactorsAndExposure.ToModel()
		risk.StatusID = sid
		if err = upsertSymptomRow(tx, &risk, sid); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// P4: merge vision-classified skin labels into symptoms before scoring.
	if len(symptoms.SymptomPhotoIDs) > 0 {
		var photos []models.SymptomPhoto
		_ = db.Where("photo_id IN ? AND user_id = ?", symptoms.SymptomPhotoIDs, userID).Find(&photos)
		var labels []VisionLabel
		for _, p := range photos {
			_ = db.Model(&p).Update("status_id", status.StatusID)
			if p.LabelsJSON == "" {
				continue
			}
			var ls []VisionLabel
			if json.Unmarshal([]byte(p.LabelsJSON), &ls) == nil {
				labels = append(labels, ls...)
			}
		}
		ApplyVisionLabelsToSymptoms(&symptoms, labels, 0.55)
	}

	expertScores, err := ComputeScores(db, status.StatusID, symptoms)
	if err != nil {
		return nil, err
	}
	expertByName := map[string]float64{}
	for _, s := range expertScores {
		expertByName[s.DiseaseTypeName] = s.Score
	}

	mode, _ := GetAIModeDTO(db)
	if !mode.IsAI {
		items := make([]ScoredDiseaseItem, 0, len(expertScores))
		for _, s := range expertScores {
			items = append(items, ScoredDiseaseItem{
				ScoringID: s.ScoringID, StatusID: s.StatusID, SymptomTime: s.SymptomTime,
				DiseaseTypeID: s.DiseaseTypeID, DiseaseTypeName: s.DiseaseTypeName,
				Score: s.Score, ExpertScore: s.Score,
			})
		}
		_ = db.Model(&models.DailyHealthStatus{}).Where("status_id = ?", status.StatusID).
			Updates(map[string]any{"scoring_disagreement": false, "scoring_confirmed_at": nil}).Error
		return finalizeHighRisk(db, userID, status, &HybridScoringResult{
			DiseasescoringList: items,
			Hybrid:             HybridMeta{IsAI: false, Message: "专家权重模式"},
		}), nil
	}

	llmDegraded := false
	var aiScores []models.DiseaseScoring
	detailByName := map[string]llmDiseaseScore{}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	if chat != nil && chat.Enabled() {
		var llmErr error
		aiScores, detailByName, llmErr = ComputeLLMScores(ctx, db, chat, status.StatusID, symptoms)
		if llmErr != nil {
			log.Printf("ai scoring: llm failed, falling back to ai weights: %v", llmErr)
			llmDegraded = true
		}
	} else {
		log.Printf("ai scoring: llm not configured, using ai_symptom_weighting")
		llmDegraded = true
	}
	if len(aiScores) == 0 {
		aiWeightScores, fbErr := ComputeAIWeightScores(db, status.StatusID, symptoms)
		if fbErr != nil {
			log.Printf("ai scoring: weight fallback failed, returning expert scores: %v", fbErr)
			items := make([]ScoredDiseaseItem, 0, len(expertScores))
			for _, s := range expertScores {
				items = append(items, ScoredDiseaseItem{
					ScoringID: s.ScoringID, StatusID: s.StatusID, SymptomTime: s.SymptomTime,
					DiseaseTypeID: s.DiseaseTypeID, DiseaseTypeName: s.DiseaseTypeName,
					Score: s.Score, ExpertScore: s.Score, Reason: "LLM 不可用，已回退专家分",
				})
			}
			return finalizeHighRisk(db, userID, status, &HybridScoringResult{
				DiseasescoringList: items,
				Hybrid: HybridMeta{
					IsAI: true, LLMDegraded: true,
					Message: "AI 不可用，已展示专家评分",
				},
			}), nil
		}
		aiScores = aiWeightScores
		llmDegraded = true
		// Load reasons written by ComputeAIWeightScores.
		var aiRows []models.AIDiseaseScoring
		_ = db.Where("status_id = ?", status.StatusID).Find(&aiRows)
		for _, r := range aiRows {
			var ev []string
			_ = json.Unmarshal([]byte(r.EvidenceJSON), &ev)
			detailByName[r.DiseaseTypeName] = llmDiseaseScore{
				DiseaseTypeName: r.DiseaseTypeName, Score: r.Score, Reason: r.Reason, Evidence: ev,
			}
		}
	}

	disagree := expertAIDisagree(expertScores, aiScores)
	falseVal := false
	trueVal := true
	flag := &falseVal
	if disagree {
		flag = &trueVal
	}
	_ = db.Model(&models.DailyHealthStatus{}).Where("status_id = ?", status.StatusID).
		Updates(map[string]any{
			"scoring_disagreement": flag,
			// Employee never confirms; CDC/admin reviews disagreement separately.
		}).Error

	items := make([]ScoredDiseaseItem, 0, len(aiScores))
	for _, s := range aiScores {
		item := ScoredDiseaseItem{
			ScoringID: s.ScoringID, StatusID: s.StatusID, SymptomTime: s.SymptomTime,
			DiseaseTypeID: s.DiseaseTypeID, DiseaseTypeName: s.DiseaseTypeName,
			Score: s.Score, ExpertScore: expertByName[s.DiseaseTypeName],
		}
		if d, ok := detailByName[s.DiseaseTypeName]; ok {
			item.Reason = d.Reason
			item.Evidence = d.Evidence
		}
		items = append(items, item)
	}

	msg := "AI 与专家权重结论一致"
	if disagree {
		msg = "AI 与专家权重 Top 结论不一致（仅供疾控复核）"
	} else if llmDegraded {
		msg = "已使用离线 AI 权重（LLM 降级）"
	}

	return finalizeHighRisk(db, userID, status, &HybridScoringResult{
		DiseasescoringList: items,
		Hybrid: HybridMeta{
			IsAI:              true,
			Disagreement:      disagree,
			NeedsHumanConfirm: false, // never push confirmation onto职工端
			Confirmed:         status.ScoringConfirmedAt != nil,
			LLMDegraded:       llmDegraded,
			ExpertTop:         topDiseaseNames(expertScores, 3),
			AITop:             topDiseaseNames(aiScores, 3),
			Message:           msg,
		},
	}), nil
}

func finalizeHighRisk(db *gorm.DB, userID uint, status *models.DailyHealthStatus, result *HybridScoringResult) *HybridScoringResult {
	if result == nil || status == nil {
		return result
	}
	alerts, err := RaiseHighRiskAlerts(db, userID, status.StatusID, result.DiseasescoringList, status.Latitude, status.Longitude)
	if err != nil {
		log.Printf("high-risk alert: %v", err)
		return result
	}
	if len(alerts) == 0 {
		return result
	}
	result.HighRisk = true
	result.RiskAlerts = alerts
	result.NearestHospital = alerts[0].NearestHospital
	return result
}

// ConfirmScoring marks today's hybrid disagreement as employee-confirmed.
func ConfirmScoring(db *gorm.DB, userID uint) (*HybridMeta, error) {
	status, err := todayStatus(db, userID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	if err := db.Model(&models.DailyHealthStatus{}).Where("status_id = ?", status.StatusID).
		Update("scoring_confirmed_at", now).Error; err != nil {
		return nil, err
	}
	disagree := status.ScoringDisagreement != nil && *status.ScoringDisagreement
	return &HybridMeta{
		IsAI:              true,
		Disagreement:      disagree,
		NeedsHumanConfirm: false,
		Confirmed:         true,
		Message:           "已确认风险提示",
	}, nil
}

func upsertSymptomRow[T any](tx *gorm.DB, row *T, statusID uint) error {
	if err := tx.Where("status_id = ?", statusID).Delete(row).Error; err != nil {
		return err
	}
	return tx.Create(row).Error
}

func todayStatus(db *gorm.DB, userID uint) (*models.DailyHealthStatus, error) {
	day := dateOnly(time.Now())
	var status models.DailyHealthStatus
	err := db.Where("user_id = ? AND check_in_date = ?", userID, day).First(&status).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNoDailyCheckIn
	}
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func dateOnly(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// SaveWorkEnvironment upserts monthly work environment for the user.
func SaveWorkEnvironment(db *gorm.DB, userID uint, dto WorkEnvironmentDTO) (*models.WorkEnvironmentInfo, error) {
	yearMonth := time.Now().Format("2006-01")
	if t, ok := parseSubmissionTime(dto.SubmissionTime); ok {
		yearMonth = t.Format("2006-01")
	}

	info := workEnvironmentFromDTO(userID, yearMonth, dto)
	var existing models.WorkEnvironmentInfo
	err := db.Where("user_id = ? AND year_month = ?", userID, yearMonth).First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.Create(&info).Error; err != nil {
			return nil, err
		}
		return &info, nil
	}
	if err != nil {
		return nil, err
	}
	info.EnvironmentID = existing.EnvironmentID
	if err := db.Model(&existing).Updates(&info).Error; err != nil {
		return nil, err
	}
	return &info, nil
}

// GetWorkEnvironment returns work environment for user and yyyy-MM.
func GetWorkEnvironment(db *gorm.DB, userID uint, yearMonth string) (*models.WorkEnvironmentInfo, error) {
	var info models.WorkEnvironmentInfo
	err := db.Where("user_id = ? AND year_month = ?", userID, yearMonth).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &info, err
}

func parseSubmissionTime(s string) (time.Time, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false
	}
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, s, time.Local); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

func workEnvironmentFromDTO(userID uint, yearMonth string, dto WorkEnvironmentDTO) models.WorkEnvironmentInfo {
	submitted := time.Now()
	if t, ok := parseSubmissionTime(dto.SubmissionTime); ok {
		submitted = t
	}
	return models.WorkEnvironmentInfo{
		UserID:                      userID,
		YearMonth:                   yearMonth,
		SubmissionTime:              submitted,
		WorkStationAltitude:         dto.WorkStationAltitude.Int(),
		WorkStationName:             dto.WorkStationName,
		DormitoryBuildingNumber:     dto.DormitoryBuildingNumber.Int(),
		DormitoryRoomNumber:         dto.DormitoryRoomNumber.Int(),
		DormitoryOccupants:          dto.DormitoryOccupants.Int(),
		DormitoryArea:               dto.DormitoryArea.Float64(),
		WindowOpenArea:              dto.WindowOpenArea.Float64(),
		VentilationCondition:        dto.VentilationCondition,
		AvgWorkRoomPeople:           dto.AvgWorkRoomPeople.Int(),
		AvgWorkFaceToFace:           dto.AvgWorkFaceToFace.Int(),
		AvgWorkPhysicalContact:      dto.AvgWorkPhysicalContact.Int(),
		AvgBreakfastRoomPeople:      dto.AvgBreakfastRoomPeople.Int(),
		AvgBreakfastFaceToFace:      dto.AvgBreakfastFaceToFace.Int(),
		AvgBreakfastPhysicalContact: dto.AvgBreakfastPhysicalContact.Int(),
		AvgLunchRoomPeople:          dto.AvgLunchRoomPeople.Int(),
		AvgLunchFaceToFace:          dto.AvgLunchFaceToFace.Int(),
		AvgLunchPhysicalContact:     dto.AvgLunchPhysicalContact.Int(),
		AvgDinnerRoomPeople:         dto.AvgDinnerRoomPeople.Int(),
		AvgDinnerFaceToFace:         dto.AvgDinnerFaceToFace.Int(),
		AvgDinnerPhysicalContact:    dto.AvgDinnerPhysicalContact.Int(),
		CollectiveActivitiesCount:   dto.CollectiveActivitiesCount.Int(),
		AvgActivityRoomPeople:       dto.AvgActivityRoomPeople.Int(),
		AvgActivityFaceToFace:       dto.AvgActivityFaceToFace.Int(),
		AvgActivityPhysicalContact:  dto.AvgActivityPhysicalContact.Int(),
		AvgMonthlyPressure:          dto.AvgMonthlyPressure.Float64(),
		MaxTemperature:              dto.MaxTemperature.Float64(),
		MinTemperature:              dto.MinTemperature.Float64(),
	}
}
