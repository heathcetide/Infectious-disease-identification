// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// DailyHealthStatus is a per-day health check-in record for a user.
type DailyHealthStatus struct {
	StatusID uint `json:"statusId" gorm:"primaryKey;autoIncrement"`

	UserID uint `json:"userId" gorm:"index;not null"`

	CheckInDate time.Time `json:"checkInDate" gorm:"type:date;index;not null"`
	// datetime (not TIME) so SQLite/MySQL both scan cleanly into time.Time
	CheckInTime time.Time `json:"checkInTime" gorm:"type:datetime;not null"`

	IsHealth *bool `json:"isHealth"`

	DiseaseTypeName          string `json:"diseaseTypeName,omitempty" gorm:"size:100"`
	DiagnosisDiseaseTypeName string `json:"diagnosisDiseaseTypeName,omitempty" gorm:"size:100"`

	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	Province     string  `json:"province,omitempty" gorm:"size:50"`
	City         string  `json:"city,omitempty" gorm:"size:50"`
	County       string  `json:"county,omitempty" gorm:"size:50"`
	LocationName string  `json:"locationName,omitempty" gorm:"size:200"`

	// Hybrid AI vs expert scoring QA (P1).
	ScoringDisagreement *bool      `json:"scoringDisagreement,omitempty"`
	ScoringConfirmedAt  *time.Time `json:"scoringConfirmedAt,omitempty"`
}

func (DailyHealthStatus) TableName() string { return constants.TableDailyHealthStatuses }
