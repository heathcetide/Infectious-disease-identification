// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/models"
)

// GetIsNeedRating returns whether the user should complete a satisfaction survey,
// plus currently open surveys for the mini-app picker.
func GetIsNeedRating(db *gorm.DB, userID uint) (*NeedRatingVO, error) {
	var count int64
	since := time.Now().AddDate(0, -1, 0)
	if err := db.Model(&models.SatisfactionRating{}).
		Where("user_id = ? AND rating_time >= ?", userID, since).
		Count(&count).Error; err != nil {
		return nil, err
	}

	var open []models.SatisfactionSurvey
	if err := db.Where("is_open = ?", true).
		Order("create_time DESC").
		Find(&open).Error; err != nil {
		return nil, err
	}
	briefs := make([]SatisfactionSurveyBrief, 0, len(open))
	for _, s := range open {
		briefs = append(briefs, SatisfactionSurveyBrief{
			SurveyID:      s.SurveyID,
			SurveyTitle:   s.SurveyTitle,
			SurveyContent: s.SurveyContent,
			IsOpen:        s.IsOpen,
		})
	}
	vo := &NeedRatingVO{
		NeedRating: count == 0 && len(briefs) > 0,
		Surveys:    briefs,
	}
	if len(briefs) > 0 {
		vo.SurveyID = briefs[0].SurveyID
		vo.SurveyTitle = briefs[0].SurveyTitle
		vo.SurveyContent = briefs[0].SurveyContent
	}
	return vo, nil
}

// SubmitSatisfaction stores a satisfaction rating.
func SubmitSatisfaction(db *gorm.DB, userID uint, dto SubmitSatisfactionDTO) error {
	ratingTime := time.Now()
	if dto.RatingTime != "" {
		if t, err := time.Parse(time.RFC3339, dto.RatingTime); err == nil {
			ratingTime = t
		}
	}
	row := models.SatisfactionRating{
		UserID:            userID,
		SurveyID:          dto.SurveyID,
		PhoneNumber:       dto.PhoneNumber,
		SatisfactionLevel: dto.SatisfactionLevel,
		RatingTime:        ratingTime,
	}
	if dto.RatingID > 0 {
		row.RatingID = dto.RatingID
	}
	return db.Create(&row).Error
}
