// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// GetIsNeedRating returns whether the user should complete a satisfaction survey.
func GetIsNeedRating(db *gorm.DB, userID uint) (*NeedRatingVO, error) {
	var count int64
	since := time.Now().AddDate(0, -1, 0)
	if err := db.Model(&models.SatisfactionRating{}).
		Where("user_id = ? AND rating_time >= ?", userID, since).
		Count(&count).Error; err != nil {
		return nil, err
	}
	return &NeedRatingVO{NeedRating: count == 0}, nil
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
