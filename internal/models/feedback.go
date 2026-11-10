// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// UserFeedback is free-form user feedback submitted to the system.
type UserFeedback struct {
	FeedbackID uint `json:"feedbackId" gorm:"primaryKey;autoIncrement"`
	UserID     uint `json:"userId" gorm:"index"`

	Name          string    `json:"name" gorm:"size:100"`
	PhoneNumber   string    `json:"phoneNumber" gorm:"size:20"`
	FeedbackTitle string    `json:"feedbackTitle" gorm:"size:200"`
	FeedbackText  string    `json:"feedbackText" gorm:"type:text"`
	FeedbackDate  time.Time `json:"feedbackDate"`
}

func (UserFeedback) TableName() string { return constants.TableUserFeedback }
