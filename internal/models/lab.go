// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// LabTestReport is a lab test report header for a user.
type LabTestReport struct {
	LabTestReportID uint `json:"labTestReportId" gorm:"primaryKey;autoIncrement"`
	UserID          uint `json:"userId" gorm:"index;not null"`

	UploadDate *time.Time `json:"uploadDate,omitempty" gorm:"type:date"`

	IsVirusAntigenTestDone      *bool  `json:"isVirusAntigenTestDone"`
	IsVirusNucleicAcidTestDone  *bool  `json:"isVirusNucleicAcidTestDone"`
	IsVirusCultureIsolationDone *bool  `json:"isVirusCultureIsolationDone"`
	IsSerologicalTestDone       *bool  `json:"isSerologicalTestDone"`
	PathogenicTestResults       string `json:"pathogenicTestResults,omitempty" gorm:"type:text"`
}

func (LabTestReport) TableName() string { return constants.TableLabTestReports }

// LabTestFile is an uploaded file attached to a lab test report.
type LabTestFile struct {
	FileID          uint `json:"fileId" gorm:"primaryKey;autoIncrement"`
	LabTestReportID uint `json:"labTestReportId" gorm:"index;not null"`

	SpecimenType string `json:"specimenType,omitempty" gorm:"size:100"`
	FileType     string `json:"fileType,omitempty" gorm:"size:50"`
	FileName     string `json:"fileName,omitempty" gorm:"size:255"`
	FilePath     string `json:"filePath,omitempty" gorm:"size:500"`
}

func (LabTestFile) TableName() string { return constants.TableLabTestFiles }
