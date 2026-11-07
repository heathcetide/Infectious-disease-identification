// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	"gorm.io/gorm"
)

// PageUsers returns a paginated user list for admin.
func PageUsers(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	users, total, err := models.User{}.List(db, page, size, models.ListFilter{Keyword: q.Keyword})
	if err != nil {
		return nil, err
	}
	return response.NewPage(users, total, page, size), nil
}

// PageHospitals lists hospitals with optional keyword on name.
func PageHospitals(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Model(&models.Hospital{})
	if q.Keyword != "" {
		query = query.Where("hospital_name LIKE ?", "%"+q.Keyword+"%")
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.Hospital
	if err := query.Order("hospital_id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	return response.NewPage(list, total, page, size), nil
}

// ListDiseaseTypes returns all disease types.
func ListDiseaseTypes(db *gorm.DB) ([]models.DiseaseType, error) {
	var list []models.DiseaseType
	err := db.Order("disease_type_id ASC").Find(&list).Error
	return list, err
}

// PageSymptomWeights lists symptom weighting rows.
func PageSymptomWeights(db *gorm.DB, q PageQuery, diseaseTypeName string) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Model(&models.SymptomWeighting{})
	if diseaseTypeName != "" {
		query = query.Where("disease_type_name = ?", diseaseTypeName)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.SymptomWeighting
	if err := query.Order("symptom_weighting_id ASC").
		Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	return response.NewPage(list, total, page, size), nil
}

// UpdateSymptomWeight updates weight score for one row.
func UpdateSymptomWeight(db *gorm.DB, dto SymptomWeightUpdateDTO) error {
	return db.Model(&models.SymptomWeighting{}).
		Where("symptom_weighting_id = ?", dto.SymptomWeightingID).
		Update("weight_score", dto.WeightScore).Error
}

// PageFeedback lists user feedback entries.
func PageFeedback(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Model(&models.UserFeedback{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.UserFeedback
	if err := query.Order("feedback_date DESC").
		Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	return response.NewPage(list, total, page, size), nil
}

// CheckinStats returns today's check-in count.
func CheckinStats(db *gorm.DB, day time.Time) (int64, error) {
	d := dateOnly(day)
	var count int64
	err := db.Model(&models.DailyHealthStatus{}).Where("check_in_date = ?", d).Count(&count).Error
	return count, err
}

// GetDashboardCounts returns aggregate counts for the admin dashboard.
func GetDashboardCounts(db *gorm.DB) (*DashboardCounts, error) {
	todayStart, todayEnd := todayDateRange()
	var out DashboardCounts
	if err := db.Model(&models.User{}).Count(&out.UserCount).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ?", todayStart, todayEnd).
		Count(&out.CheckInTodayCount).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DiagnosisResults{}).
		Where("submission_time >= ? AND submission_time < ?", todayStart, todayEnd).
		Count(&out.DiagnosisTodayCount).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.LabTestReport{}).
		Where("upload_date >= ? AND upload_date < ?", todayStart, todayEnd).
		Count(&out.LabReportTodayCount).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.UserFeedback{}).Count(&out.FeedbackCount).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

// PageDiagnosisResults lists diagnosis submissions for admin review.
func PageDiagnosisResults(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Model(&models.DiagnosisResults{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.DiagnosisResults
	if err := query.Order("diagnosis_results_id DESC").
		Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	return response.NewPage(list, total, page, size), nil
}

// PageLabReports lists lab reports for admin review.
func PageLabReports(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Model(&models.LabTestReport{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.LabTestReport
	if err := query.Order("lab_test_report_id DESC").
		Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	return response.NewPage(list, total, page, size), nil
}
