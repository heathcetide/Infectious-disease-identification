// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"time"

	"github.com/heathcetide/infectious-disease/pkg/common/response"
	"gorm.io/gorm"
)

// MedicalPageDiagnosis lists diagnosis records for medical staff.
func MedicalPageDiagnosis(db *gorm.DB, q PageQuery) (*response.Page, error) {
	return PageDiagnosisResults(db, q)
}

// MedicalPageLabReports lists lab reports for medical staff.
func MedicalPageLabReports(db *gorm.DB, q PageQuery) (*response.Page, error) {
	return PageLabReports(db, q)
}

// MedicalPageHospitals lists hospitals for medical staff.
func MedicalPageHospitals(db *gorm.DB, q PageQuery) (*response.Page, error) {
	return PageHospitals(db, q)
}

// MedicalCheckinStats returns check-in stats for a given day (default today).
func MedicalCheckinStats(db *gorm.DB, day time.Time) (int64, error) {
	if day.IsZero() {
		day = time.Now()
	}
	return CheckinStats(db, day)
}
