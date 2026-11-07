// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"github.com/heathcetide/infectious-disease/internal/constants"
)

// Hospital is a hospital reference record.
type Hospital struct {
	HospitalID          uint   `json:"hospitalId" gorm:"primaryKey;autoIncrement"`
	HospitalName        string `json:"hospitalName" gorm:"size:200;not null"`
	HospitalPhoneNumber string `json:"hospitalPhoneNumber,omitempty" gorm:"size:30"`
	Address             string `json:"address,omitempty" gorm:"size:500"`
}

func (Hospital) TableName() string { return constants.TableHospitals }
