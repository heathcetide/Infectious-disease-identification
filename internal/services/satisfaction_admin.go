// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"errors"
	"math"
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
)

// SatisfactionSurveyCreateDTO is create/update body for admin surveys.
type SatisfactionSurveyCreateDTO struct {
	SurveyID      uint   `json:"surveyID"`
	SurveyTitle   string `json:"surveyTitle"`
	SurveyContent string `json:"surveyContent"`
}

// SatisfactionSurveyRow is list item for admin satisfaction table.
type SatisfactionSurveyRow struct {
	SurveyID      uint      `json:"surveyID"`
	SurveyTitle   string    `json:"surveyTitle"`
	SurveyContent string    `json:"surveyContent"`
	IsOpen        bool      `json:"isOpen"`
	Satisfaction  *float64  `json:"satisfaction"`
	CreateTime    time.Time `json:"createTime"`
}

// SatisfactionStatsVO is per-survey rating counts.
type SatisfactionStatsVO struct {
	VeryHappy   int64 `json:"veryHappy"`
	Happy       int64 `json:"happy"`
	Normal      int64 `json:"normal"`
	NoHappy     int64 `json:"noHappy"`
	VeryNoHappy int64 `json:"veryNoHappy"`
}

var satisfactionLevelScore = map[string]float64{
	"非常满意":  5,
	"满意":    4,
	"一般":    3,
	"不满意":   2,
	"非常不满意": 1,
}

func satisfactionAvg(db *gorm.DB, surveyID uint) *float64 {
	var ratings []models.SatisfactionRating
	if err := db.Where("survey_id = ?", surveyID).Find(&ratings).Error; err != nil || len(ratings) == 0 {
		return nil
	}
	var sum float64
	var n int
	for _, r := range ratings {
		if sc, ok := satisfactionLevelScore[r.SatisfactionLevel]; ok {
			sum += sc
			n++
		}
	}
	if n == 0 {
		return nil
	}
	v := sum / float64(n) / 5.0
	v = math.Round(v*1e4) / 1e4
	return &v
}

// CreateSatisfactionSurvey inserts a new survey (closed by default).
func CreateSatisfactionSurvey(db *gorm.DB, dto SatisfactionSurveyCreateDTO) (*models.SatisfactionSurvey, error) {
	row := models.SatisfactionSurvey{
		SurveyTitle:   dto.SurveyTitle,
		SurveyContent: dto.SurveyContent,
		IsOpen:        false,
		CreateTime:    time.Now(),
	}
	if err := db.Create(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}

// UpdateSatisfactionSurvey updates title/content.
func UpdateSatisfactionSurvey(db *gorm.DB, id uint, dto SatisfactionSurveyCreateDTO) (*models.SatisfactionSurvey, error) {
	var row models.SatisfactionSurvey
	if err := db.First(&row, id).Error; err != nil {
		return nil, err
	}
	updates := map[string]any{}
	if dto.SurveyTitle != "" {
		updates["survey_title"] = dto.SurveyTitle
	}
	if dto.SurveyContent != "" {
		updates["survey_content"] = dto.SurveyContent
	}
	if len(updates) > 0 {
		if err := db.Model(&row).Updates(updates).Error; err != nil {
			return nil, err
		}
	}
	if err := db.First(&row, id).Error; err != nil {
		return nil, err
	}
	return &row, nil
}

func setSurveyOpen(db *gorm.DB, id uint, open bool) error {
	now := time.Now()
	updates := map[string]any{"is_open": open}
	if open {
		updates["close_time"] = nil
	} else {
		updates["close_time"] = now
	}
	res := db.Model(&models.SatisfactionSurvey{}).Where("survey_id = ?", id).Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// OpenSatisfactionSurvey marks a survey open.
func OpenSatisfactionSurvey(db *gorm.DB, id uint) error {
	return setSurveyOpen(db, id, true)
}

// CloseSatisfactionSurvey marks a survey closed.
func CloseSatisfactionSurvey(db *gorm.DB, id uint) error {
	return setSurveyOpen(db, id, false)
}

// GetSatisfactionStats counts ratings by level for one survey.
func GetSatisfactionStats(db *gorm.DB, surveyID uint) (SatisfactionStatsVO, error) {
	out := SatisfactionStatsVO{}
	type row struct {
		Level string
		Cnt   int64
	}
	var rows []row
	err := db.Model(&models.SatisfactionRating{}).
		Select("satisfaction_level as level, count(*) as cnt").
		Where("survey_id = ?", surveyID).
		Group("satisfaction_level").
		Scan(&rows).Error
	if err != nil {
		return out, err
	}
	for _, r := range rows {
		switch r.Level {
		case "非常满意":
			out.VeryHappy = r.Cnt
		case "满意":
			out.Happy = r.Cnt
		case "一般":
			out.Normal = r.Cnt
		case "不满意":
			out.NoHappy = r.Cnt
		case "非常不满意":
			out.VeryNoHappy = r.Cnt
		}
	}
	return out, nil
}

// PageSatisfactionSurveys lists surveys for admin.
func PageSatisfactionSurveys(db *gorm.DB, page, size int) (*response.AdminPage, error) {
	query := db.Model(&models.SatisfactionSurvey{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var raw []models.SatisfactionSurvey
	if err := query.Order("survey_id DESC").Offset((page - 1) * size).Limit(size).Find(&raw).Error; err != nil {
		return nil, err
	}
	list := make([]SatisfactionSurveyRow, 0, len(raw))
	for _, s := range raw {
		list = append(list, SatisfactionSurveyRow{
			SurveyID:      s.SurveyID,
			SurveyTitle:   s.SurveyTitle,
			SurveyContent: s.SurveyContent,
			IsOpen:        s.IsOpen,
			Satisfaction:  satisfactionAvg(db, s.SurveyID),
			CreateTime:    s.CreateTime,
		})
	}
	return response.NewAdminPage(list, total, page, size), nil
}

// PageSatisfactionRatings lists ratings for a survey.
func PageSatisfactionRatings(db *gorm.DB, surveyID uint, page, size int) (*response.AdminPage, error) {
	query := db.Model(&models.SatisfactionRating{})
	if surveyID > 0 {
		query = query.Where("survey_id = ?", surveyID)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.SatisfactionRating
	if err := query.Order("rating_id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	return response.NewAdminPage(list, total, page, size), nil
}

// GetOpenSatisfactionSurvey returns the single open survey if any.
func GetOpenSatisfactionSurvey(db *gorm.DB) (*models.SatisfactionSurvey, error) {
	var s models.SatisfactionSurvey
	err := db.Where("is_open = ?", true).Order("survey_id DESC").First(&s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &s, err
}
