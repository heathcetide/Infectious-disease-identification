// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// SatisfactionSurvey is an admin-published satisfaction questionnaire.
type SatisfactionSurvey struct {
	SurveyID      uint       `json:"surveyID" gorm:"primaryKey;autoIncrement;column:survey_id"`
	SurveyTitle   string     `json:"surveyTitle" gorm:"size:200;not null"`
	SurveyContent string     `json:"surveyContent" gorm:"type:text"`
	IsOpen        bool       `json:"isOpen" gorm:"default:false"`
	CreateTime    time.Time  `json:"createTime"`
	CloseTime     *time.Time `json:"closeTime,omitempty"`
}

func (SatisfactionSurvey) TableName() string { return constants.TableSatisfactionSurveys }

// SatisfactionRating records a user's satisfaction survey response (frontend scoring page).
type SatisfactionRating struct {
	RatingID uint `json:"ratingId" gorm:"primaryKey;autoIncrement"`

	UserID   uint `json:"userId" gorm:"index;not null"`
	SurveyID uint `json:"surveyId" gorm:"index"`

	PhoneNumber       string    `json:"phoneNumber,omitempty" gorm:"size:20"`
	SatisfactionLevel string    `json:"satisfactionLevel" gorm:"size:50;not null"`
	RatingTime        time.Time `json:"ratingTime"`
}

func (SatisfactionRating) TableName() string { return constants.TableSatisfactionRatings }
