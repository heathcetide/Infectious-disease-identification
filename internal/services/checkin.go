// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"errors"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
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

	var scores []models.DiseaseScoring
	if err := db.Where("status_id = ? AND symptom_time = ?", status.StatusID, day).
		Find(&scores).Error; err != nil {
		return nil, err
	}

	return &DailyCheckInVO{
		HasSignedToday:     true,
		HasMonthSigned:     monthSigned,
		Dailyhealthstatus:  &status,
		DiseasescoringList: scores,
	}, nil
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
func SaveSymptomsAndScore(db *gorm.DB, userID uint, symptoms AllSymptoms) ([]models.DiseaseScoring, error) {
	status, err := todayStatus(db, userID)
	if err != nil {
		return nil, err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		sid := status.StatusID
		general := symptoms.Generalsymptoms.ToModel()
		general.StatusID = sid
		if err := upsertSymptomRow(tx, &general, sid); err != nil {
			return err
		}
		symptoms.Respiratorysymptoms.StatusID = sid
		if err := upsertSymptomRow(tx, &symptoms.Respiratorysymptoms, sid); err != nil {
			return err
		}
		digestive := symptoms.Digestivesymptoms.ToModel()
		digestive.StatusID = sid
		if err := upsertSymptomRow(tx, &digestive, sid); err != nil {
			return err
		}
		symptoms.Circulatorysymptoms.StatusID = sid
		if err := upsertSymptomRow(tx, &symptoms.Circulatorysymptoms, sid); err != nil {
			return err
		}
		symptoms.NeurologicalSymptoms.StatusID = sid
		if err := upsertSymptomRow(tx, &symptoms.NeurologicalSymptoms, sid); err != nil {
			return err
		}
		symptoms.LocalSymptoms.StatusID = sid
		if err := upsertSymptomRow(tx, &symptoms.LocalSymptoms, sid); err != nil {
			return err
		}
		other := symptoms.OtherSymptoms
		if other.WorkLifeStress == "" {
			other.WorkLifeStress = "小"
		}
		other.StatusID = sid
		if err := upsertSymptomRow(tx, &other, sid); err != nil {
			return err
		}
		risk := symptoms.RiskFactorsAndExposure.ToModel()
		risk.StatusID = sid
		if err := upsertSymptomRow(tx, &risk, sid); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return ComputeScores(db, status.StatusID, symptoms)
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
		if err := db.Create(&info).Error; err != nil {
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
