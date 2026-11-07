// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"bytes"
	"encoding/csv"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// CheckInDetailVO aggregates a check-in with scores for CDC review.
type CheckInDetailVO struct {
	Status           models.DailyHealthStatus   `json:"status"`
	User             *models.User               `json:"user,omitempty"`
	Profile          *models.EmployeeProfile    `json:"profile,omitempty"`
	DiseaseScoring   []models.DiseaseScoring    `json:"diseaseScoringList"`
	AIDiseaseScoring []models.AIDiseaseScoring  `json:"aiDiseaseScoringList,omitempty"`
}

// GetCheckInDetail loads check-in by status ID with user and scores.
func GetCheckInDetail(db *gorm.DB, statusID uint) (*CheckInDetailVO, error) {
	var status models.DailyHealthStatus
	if err := db.First(&status, statusID).Error; err != nil {
		return nil, err
	}
	vo := &CheckInDetailVO{Status: status}
	if u, err := (models.User{}).FindByID(db, status.UserID); err == nil {
		vo.User = u
	}
	var profile models.EmployeeProfile
	if err := db.Where("user_id = ?", status.UserID).First(&profile).Error; err == nil {
		vo.Profile = &profile
	}
	day := status.CheckInDate
	_ = db.Where("status_id = ? AND symptom_time = ?", statusID, day).Find(&vo.DiseaseScoring).Error
	_ = db.Where("status_id = ? AND symptom_time = ?", statusID, day).Find(&vo.AIDiseaseScoring).Error
	return vo, nil
}

// CheckInLocationVO is location-focused check-in info.
type CheckInLocationVO struct {
	CheckInDetailVO
}

// GetCheckInLocation is an alias detail for location endpoints.
func GetCheckInLocation(db *gorm.DB, statusID uint) (*CheckInLocationVO, error) {
	d, err := GetCheckInDetail(db, statusID)
	if err != nil {
		return nil, err
	}
	return &CheckInLocationVO{CheckInDetailVO: *d}, nil
}

// ExportCheckInsCSV exports all daily health rows.
func ExportCheckInsCSV(db *gorm.DB) ([]byte, error) {
	var list []models.DailyHealthStatus
	if err := db.Order("status_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"statusId", "userId", "checkInDate", "isHealth", "province", "city", "county", "locationName", "longitude", "latitude"})
	for _, s := range list {
		health := ""
		if s.IsHealth != nil {
			health = strconv.FormatBool(*s.IsHealth)
		}
		_ = w.Write([]string{
			strconv.FormatUint(uint64(s.StatusID), 10),
			strconv.FormatUint(uint64(s.UserID), 10),
			s.CheckInDate.Format("2006-01-02"),
			health, s.Province, s.City, s.County, s.LocationName,
			strconv.FormatFloat(s.Longitude, 'f', -1, 64),
			strconv.FormatFloat(s.Latitude, 'f', -1, 64),
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

// CDCStatisticsData matches Java StatisticsDataVO core fields.
type CDCStatisticsData struct {
	UserNumber        int64 `json:"userNumber"`
	CheckInNumber     int64 `json:"checkInNumber"`
	UserHealthNumber  int64 `json:"userHealthNumber"`
	UserDiseaseNumber int64 `json:"userDiseaseNumber"`
	CheckInToday      int64 `json:"checkInToday"`
}

// QueryCDCStatisticsData loads dashboard statistics.
func QueryCDCStatisticsData(db *gorm.DB) (*CDCStatisticsData, error) {
	todayStart, todayEnd := todayDateRange()
	var out CDCStatisticsData
	if err := db.Model(&models.User{}).Where("status = ?", models.UserStatusActive).Count(&out.UserNumber).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).Count(&out.CheckInNumber).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ?", todayStart, todayEnd).
		Count(&out.CheckInToday).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ? AND is_health = ?", todayStart, todayEnd, true).
		Count(&out.UserHealthNumber).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ? AND is_health = ?", todayStart, todayEnd, false).
		Count(&out.UserDiseaseNumber).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

// DiseaseNameCount is a disease label with count/score sum.
type DiseaseNameCount struct {
	DiseaseTypeName string  `json:"diseaseTypeName"`
	Count           int64   `json:"count"`
	TotalScore      float64 `json:"totalScore"`
}

// CDCDiseaseToday aggregates today's disease scoring rows.
func CDCDiseaseToday(db *gorm.DB) ([]DiseaseNameCount, error) {
	todayStart, todayEnd := todayDateRange()
	var statusIDs []uint
	if err := db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ?", todayStart, todayEnd).
		Pluck("status_id", &statusIDs).Error; err != nil {
		return nil, err
	}
	if len(statusIDs) == 0 {
		return []DiseaseNameCount{}, nil
	}
	type row struct {
		Name  string
		Cnt   int64
		Score float64
	}
	var rows []row
	err := db.Model(&models.DiseaseScoring{}).
		Select("disease_type_name as name, count(*) as cnt, sum(score) as score").
		Where("status_id IN ?", statusIDs).
		Group("disease_type_name").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]DiseaseNameCount, 0, len(rows))
	for _, r := range rows {
		out = append(out, DiseaseNameCount{DiseaseTypeName: r.Name, Count: r.Cnt, TotalScore: r.Score})
	}
	return out, nil
}

// UserStation is a map point for today's check-ins.
type UserStation struct {
	UserID       uint    `json:"userId"`
	Name         string  `json:"name"`
	PhoneNumber  string  `json:"phoneNumber"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	LocationName string  `json:"locationName"`
}

// CDCUserStations returns today's check-in locations with user info.
func CDCUserStations(db *gorm.DB) ([]UserStation, error) {
	todayStart, todayEnd := todayDateRange()
	var statuses []models.DailyHealthStatus
	if err := db.Where("check_in_date >= ? AND check_in_date < ?", todayStart, todayEnd).
		Find(&statuses).Error; err != nil {
		return nil, err
	}
	out := make([]UserStation, 0, len(statuses))
	for _, s := range statuses {
		st := UserStation{
			UserID: s.UserID, Longitude: s.Longitude, Latitude: s.Latitude, LocationName: s.LocationName,
		}
		if u, err := (models.User{}).FindByID(db, s.UserID); err == nil {
			st.Name = u.Name
			st.PhoneNumber = userPhoneField(u)
		}
		out = append(out, st)
	}
	return out, nil
}

func userPhoneField(u *models.User) string {
	if u.PhoneNumber != "" {
		return u.PhoneNumber
	}
	return u.Phone
}

// CheckInSearchItem is a search result row.
type CheckInSearchItem struct {
	StatusID    uint      `json:"statusId"`
	UserID      uint      `json:"userId"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	CheckInDate time.Time `json:"checkInDate"`
	IsHealth    *bool     `json:"isHealth"`
}

// CDCCheckInSearch finds recent check-ins matching q (phone or name).
func CDCCheckInSearch(db *gorm.DB, q string, limit int) ([]CheckInSearchItem, error) {
	q = strings.TrimSpace(q)
	if limit < 1 {
		limit = 20
	}
	query := db.Model(&models.DailyHealthStatus{}).Order("status_id DESC").Limit(limit)
	if q != "" {
		like := "%" + q + "%"
		var userIDs []uint
		_ = db.Model(&models.User{}).
			Where("name LIKE ? OR phone_number LIKE ? OR phone LIKE ?", like, like, like).
			Pluck("id", &userIDs)
		if len(userIDs) > 0 {
			query = query.Where("user_id IN ?", userIDs)
		} else {
			return []CheckInSearchItem{}, nil
		}
	}
	var statuses []models.DailyHealthStatus
	if err := query.Find(&statuses).Error; err != nil {
		return nil, err
	}
	out := make([]CheckInSearchItem, 0, len(statuses))
	for _, s := range statuses {
		item := CheckInSearchItem{
			StatusID: s.StatusID, UserID: s.UserID, CheckInDate: s.CheckInDate, IsHealth: s.IsHealth,
		}
		if u, err := (models.User{}).FindByID(db, s.UserID); err == nil {
			item.Name = u.Name
			item.PhoneNumber = userPhoneField(u)
		}
		out = append(out, item)
	}
	return out, nil
}

// CheckInDailySeriesRequest is the POST body for daily series.
type CheckInDailySeriesRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// CheckInDailyNumber is count per day.
type CheckInDailyNumber struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// CDCCheckInDailySeries counts check-ins per day in range.
func CDCCheckInDailySeries(db *gorm.DB, req CheckInDailySeriesRequest) ([]CheckInDailyNumber, error) {
	start, err := ParseDate(req.StartDate)
	if err != nil {
		return nil, ErrInvalidDate
	}
	end, err := ParseDate(req.EndDate)
	if err != nil {
		return nil, ErrInvalidDate
	}
	end = end.AddDate(0, 0, 1)
	var list []CheckInDailyNumber
	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		dayEnd := d.AddDate(0, 0, 1)
		var cnt int64
		if err := db.Model(&models.DailyHealthStatus{}).
			Where("check_in_date >= ? AND check_in_date < ?", d, dayEnd).
			Count(&cnt).Error; err != nil {
			return nil, err
		}
		list = append(list, CheckInDailyNumber{Date: d.Format("2006-01-02"), Count: cnt})
	}
	return list, nil
}

// ImportantUserInfo highlights users with unhealthy check-in today.
type ImportantUserInfo struct {
	UserID      uint   `json:"userId"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Reason      string `json:"reason"`
}

// CDCImportantUsers lists users flagged today.
func CDCImportantUsers(db *gorm.DB) ([]ImportantUserInfo, error) {
	todayStart, todayEnd := todayDateRange()
	var statuses []models.DailyHealthStatus
	if err := db.Where("check_in_date >= ? AND check_in_date < ? AND is_health = ?",
		todayStart, todayEnd, false).Find(&statuses).Error; err != nil {
		return nil, err
	}
	out := make([]ImportantUserInfo, 0, len(statuses))
	for _, s := range statuses {
		info := ImportantUserInfo{UserID: s.UserID, Reason: "unhealthy check-in today"}
		if u, err := (models.User{}).FindByID(db, s.UserID); err == nil {
			info.Name = u.Name
			info.PhoneNumber = userPhoneField(u)
		}
		out = append(out, info)
	}
	return out, nil
}

// DiseaseStatisticsQuery filters aggregated scores.
type DiseaseStatisticsQuery struct {
	DiseaseTypeName string
	StartDate       string
	EndDate         string
}

// CDCDiseaseStatistics aggregates disease scores in a date range.
func CDCDiseaseStatistics(db *gorm.DB, q DiseaseStatisticsQuery) ([]DiseaseNameCount, error) {
	query := db.Model(&models.DiseaseScoring{})
	if q.DiseaseTypeName != "" {
		query = query.Where("disease_type_name = ?", q.DiseaseTypeName)
	}
	if q.StartDate != "" {
		start, err := ParseDate(q.StartDate)
		if err != nil {
			return nil, ErrInvalidDate
		}
		query = query.Where("symptom_time >= ?", start)
	}
	if q.EndDate != "" {
		end, err := ParseDate(q.EndDate)
		if err != nil {
			return nil, ErrInvalidDate
		}
		query = query.Where("symptom_time < ?", end.AddDate(0, 0, 1))
	}
	type row struct {
		Name  string
		Cnt   int64
		Score float64
	}
	var rows []row
	err := query.Select("disease_type_name as name, count(*) as cnt, sum(score) as score").
		Group("disease_type_name").Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]DiseaseNameCount, 0, len(rows))
	for _, r := range rows {
		out = append(out, DiseaseNameCount{DiseaseTypeName: r.Name, Count: r.Cnt, TotalScore: r.Score})
	}
	return out, nil
}

// GetWorkEnvironmentByID loads work environment row.
func GetWorkEnvironmentByID(db *gorm.DB, id uint) (*models.WorkEnvironmentInfo, error) {
	var info models.WorkEnvironmentInfo
	if err := db.First(&info, id).Error; err != nil {
		return nil, err
	}
	return &info, nil
}

// ExportWorkEnvironmentsCSV exports work environment rows.
func ExportWorkEnvironmentsCSV(db *gorm.DB) ([]byte, error) {
	var list []models.WorkEnvironmentInfo
	if err := db.Order("environment_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"environmentId", "userId", "yearMonth", "workStationName"})
	for _, e := range list {
		_ = w.Write([]string{
			strconv.FormatUint(uint64(e.EnvironmentID), 10),
			strconv.FormatUint(uint64(e.UserID), 10),
			e.YearMonth, e.WorkStationName,
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

// LearningMaterialCreateDTO for POST/PUT materials (multipart or JSON).
// Legacy form field "filePath" is accepted as an alias for Link (CDN URL).
type LearningMaterialCreateDTO struct {
	DiseaseTypeName string `json:"diseaseTypeName" form:"diseaseTypeName"`
	Title           string `json:"title" form:"title"`
	MaterialType    string `json:"materialType" form:"materialType"`
	Link            string `json:"link" form:"link"`
	FilePath        string `json:"filePath" form:"filePath"` // deprecated alias → Link
	SequenceNumber  int    `json:"sequenceNumber" form:"sequenceNumber"`
}

func resolveMaterialLink(dto *LearningMaterialCreateDTO) string {
	link := strings.TrimSpace(dto.Link)
	if link == "" {
		link = strings.TrimSpace(dto.FilePath)
	}
	return link
}

// CreateLearningMaterial inserts a material (aligned with Java saveMaterial).
func CreateLearningMaterial(db *gorm.DB, dto LearningMaterialCreateDTO) (*models.DiseaseLearningMaterial, error) {
	deleted := false
	now := time.Now()
	link := resolveMaterialLink(&dto)

	var seq int
	_ = db.Model(&models.DiseaseLearningMaterial{}).Select("COALESCE(MAX(sequence_number),0)").Scan(&seq)
	if dto.SequenceNumber > 0 {
		seq = dto.SequenceNumber
	} else {
		seq++
	}

	diseaseID := 0
	if name := strings.TrimSpace(dto.DiseaseTypeName); name != "" {
		var dt models.DiseaseType
		if err := db.Where("disease_type_name = ?", name).First(&dt).Error; err == nil {
			diseaseID = int(dt.DiseaseTypeID)
		}
	}

	m := models.DiseaseLearningMaterial{
		DiseaseTypeID:   diseaseID,
		DiseaseTypeName: dto.DiseaseTypeName,
		Title:           dto.Title,
		MaterialType:    dto.MaterialType,
		Link:            link,
		SequenceNumber:  seq,
		PublishDate:     &now,
		IsDeleted:       &deleted,
	}
	if err := db.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// GetLearningMaterialByID loads one material.
func GetLearningMaterialByID(db *gorm.DB, id uint) (*models.DiseaseLearningMaterial, error) {
	var m models.DiseaseLearningMaterial
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// UpdateLearningMaterial updates fields.
func UpdateLearningMaterial(db *gorm.DB, id uint, dto LearningMaterialCreateDTO) (*models.DiseaseLearningMaterial, error) {
	m, err := GetLearningMaterialByID(db, id)
	if err != nil {
		return nil, err
	}
	updates := map[string]any{}
	if dto.DiseaseTypeName != "" {
		updates["disease_type_name"] = dto.DiseaseTypeName
		var dt models.DiseaseType
		if err := db.Where("disease_type_name = ?", dto.DiseaseTypeName).First(&dt).Error; err == nil {
			updates["disease_type_id"] = int(dt.DiseaseTypeID)
		}
	}
	if dto.Title != "" {
		updates["title"] = dto.Title
	}
	if dto.MaterialType != "" {
		updates["material_type"] = dto.MaterialType
	}
	if link := resolveMaterialLink(&dto); link != "" {
		updates["link"] = link
	}
	if dto.SequenceNumber > 0 {
		updates["sequence_number"] = dto.SequenceNumber
	}
	if len(updates) > 0 {
		if err := db.Model(m).Updates(updates).Error; err != nil {
			return nil, err
		}
	}
	return GetLearningMaterialByID(db, id)
}

// SetLearningMaterialStatus sets IsDeleted (enable/disable).
func SetLearningMaterialStatus(db *gorm.DB, id uint, disabled bool) error {
	m, err := GetLearningMaterialByID(db, id)
	if err != nil {
		return err
	}
	m.IsDeleted = &disabled
	if disabled {
		now := time.Now()
		m.DeletedDate = &now
	}
	return db.Model(m).Updates(map[string]any{
		"is_deleted":   disabled,
		"deleted_date": m.DeletedDate,
	}).Error
}

// MaterialSequenceDirection body.
type MaterialSequenceDirection struct {
	Direction string `json:"direction"` // up|down
}

// MoveLearningMaterialSequence swaps sequence with neighbor in same disease group.
func MoveLearningMaterialSequence(db *gorm.DB, id uint, direction string) error {
	cur, err := GetLearningMaterialByID(db, id)
	if err != nil {
		return err
	}
	var neighbor models.DiseaseLearningMaterial
	q := db.Where("disease_type_name = ?", cur.DiseaseTypeName)
	if strings.EqualFold(direction, "up") {
		q = q.Where("sequence_number < ?", cur.SequenceNumber).Order("sequence_number DESC")
	} else {
		q = q.Where("sequence_number > ?", cur.SequenceNumber).Order("sequence_number ASC")
	}
	if err := q.First(&neighbor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return db.Transaction(func(tx *gorm.DB) error {
		a, b := cur.SequenceNumber, neighbor.SequenceNumber
		if err := tx.Model(cur).Update("sequence_number", b).Error; err != nil {
			return err
		}
		return tx.Model(&neighbor).Update("sequence_number", a).Error
	})
}
