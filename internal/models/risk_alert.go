// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// RiskAlert is raised when a check-in disease score exceeds the high-risk threshold (e.g. 80%).
type RiskAlert struct {
	AlertID         uint      `json:"alertId" gorm:"primaryKey;autoIncrement"`
	UserID          uint      `json:"userId,string" gorm:"index;not null"`
	StatusID        uint      `json:"statusId,string" gorm:"index"`
	DiseaseTypeName string    `json:"diseaseTypeName" gorm:"size:100;index"`
	Score           float64   `json:"score"`
	Threshold       float64   `json:"threshold" gorm:"default:80"`
	NearestHospital string    `json:"nearestHospital,omitempty" gorm:"size:200"`
	HospitalID      *uint     `json:"hospitalId,omitempty"`
	Latitude        float64   `json:"latitude,omitempty"`
	Longitude       float64   `json:"longitude,omitempty"`
	Message         string    `json:"message,omitempty" gorm:"size:500"`
	Acknowledged    bool      `json:"acknowledged" gorm:"default:false"`
	CreatedAt       time.Time `json:"createdAt"`
}

func (RiskAlert) TableName() string { return constants.TableRiskAlerts }
