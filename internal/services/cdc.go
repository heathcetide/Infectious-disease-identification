// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
)

// CDCPageLearningMaterials delegates to learning material pagination.
func CDCPageLearningMaterials(db *gorm.DB, dto LearningMaterialsPageDTO) (*response.Page, error) {
	return PageSelectLearningMaterials(db, dto)
}

// CDCPageDiagnosis lists diagnosis cases for CDC staff.
func CDCPageDiagnosis(db *gorm.DB, q PageQuery) (*response.Page, error) {
	return PageDiagnosisResults(db, q)
}

// CDCPageLabReports lists lab reports for CDC staff.
func CDCPageLabReports(db *gorm.DB, q PageQuery) (*response.Page, error) {
	return PageLabReports(db, q)
}

// CDCCheckinStats returns today's check-in count.
func CDCCheckinStats(db *gorm.DB) (int64, error) {
	return CheckinStats(db, dateOnlyNow())
}

// CDCListDiseaseTypes returns disease type catalog.
func CDCListDiseaseTypes(db *gorm.DB) ([]models.DiseaseType, error) {
	return ListDiseaseTypes(db)
}

func dateOnlyNow() (t time.Time) { return dateOnly(time.Now()) }
