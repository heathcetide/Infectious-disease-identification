// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/models"
)

// LabReportWithFiles is a lab report plus attached files.
type LabReportWithFiles struct {
	Report models.LabTestReport `json:"report"`
	Files  []models.LabTestFile `json:"files"`
}

// SaveLabReport creates today's lab report and links file IDs.
func SaveLabReport(db *gorm.DB, userID uint, dto LabTestReportDTO) error {
	today := dateOnly(time.Now())
	var count int64
	if err := db.Model(&models.LabTestReport{}).
		Where("user_id = ? AND upload_date = ?", userID, today).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrAlreadySubmitted
	}

	return db.Transaction(func(tx *gorm.DB) error {
		antigen := dto.IsVirusAntigenTestDone
		nucleic := dto.IsVirusNucleicAcidTestDone
		culture := dto.IsVirusCultureIsolationDone
		sero := dto.IsSerologicalTestDone
		report := models.LabTestReport{
			UserID:                      userID,
			UploadDate:                  &today,
			PathogenicTestResults:       dto.PathogenicTestResults,
			IsVirusAntigenTestDone:      &antigen,
			IsVirusNucleicAcidTestDone:  &nucleic,
			IsVirusCultureIsolationDone: &culture,
			IsSerologicalTestDone:       &sero,
		}
		if err := tx.Create(&report).Error; err != nil {
			return err
		}
		for _, fileID := range dto.LabTestFileIDs {
			if err := tx.Model(&models.LabTestFile{}).
				Where("file_id = ?", fileID).
				Update("lab_test_report_id", report.LabTestReportID).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetLabReportByDate loads lab report and files for a user on a date.
func GetLabReportByDate(db *gorm.DB, userID uint, dateStr string) (*LabReportWithFiles, error) {
	day, err := ParseDate(dateStr)
	if err != nil {
		return nil, ErrInvalidDate
	}
	var report models.LabTestReport
	err = db.Where("user_id = ? AND upload_date = ?", userID, day).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var files []models.LabTestFile
	if err := db.Where("lab_test_report_id = ?", report.LabTestReportID).Find(&files).Error; err != nil {
		return nil, err
	}
	return &LabReportWithFiles{Report: report, Files: files}, nil
}

// LinkLabFiles attaches uploaded files to an existing lab report.
func LinkLabFiles(db *gorm.DB, labTestReportID uint, fileIDs []uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, id := range fileIDs {
			if err := tx.Model(&models.LabTestFile{}).
				Where("file_id = ?", id).
				Update("lab_test_report_id", labTestReportID).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
