// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
)

// PageUsers returns a paginated user list for admin (with employee profile fields).
func PageUsers(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Model(&models.User{})

	if role := roleFromUserTypeLabel(q.UserType); role != "" {
		query = query.Where("role = ?", role)
	}
	if q.Name != "" {
		query = query.Where("name LIKE ?", "%"+q.Name+"%")
	}
	if q.PhoneNumber != "" {
		query = query.Where("phone_number LIKE ? OR phone LIKE ?", "%"+q.PhoneNumber+"%", "%"+q.PhoneNumber+"%")
	}
	if q.Keyword != "" {
		like := "%" + q.Keyword + "%"
		query = query.Where("name LIKE ? OR phone_number LIKE ? OR phone LIKE ? OR username LIKE ?", like, like, like, like)
	}

	needProfileJoin := q.Department != "" || q.SpecificOccupation != ""
	if needProfileJoin {
		query = query.Joins("LEFT JOIN employee_profiles ON employee_profiles.user_id = users.id")
		if q.Department != "" {
			query = query.Where("employee_profiles.department LIKE ?", "%"+q.Department+"%")
		}
		if q.SpecificOccupation != "" {
			query = query.Where("employee_profiles.specific_occupation LIKE ?", "%"+q.SpecificOccupation+"%")
		}
	}

	var total int64
	if err := query.Distinct("users.id").Count(&total).Error; err != nil {
		return nil, err
	}
	var users []models.User
	if err := query.Select("users.*").Order("users.id DESC").
		Offset((page - 1) * size).Limit(size).Find(&users).Error; err != nil {
		return nil, err
	}

	list := make([]AdminUserListItem, 0, len(users))
	for _, u := range users {
		item := AdminUserListItem{
			UserID:      strconv.FormatUint(uint64(u.ID), 10),
			UserType:    roleDisplayName(u.Role),
			Name:        u.Name,
			PhoneNumber: firstNonEmpty(u.PhoneNumber, u.Phone),
			IsActived:   u.IsActive(),
		}
		var p models.EmployeeProfile
		if err := db.Where("user_id = ?", u.ID).First(&p).Error; err == nil {
			item.Gender = p.Gender
			item.Age = p.Age
			item.Ethnicity = p.Ethnicity
			item.EducationLevel = p.EducationLevel
			item.Department = p.Department
			item.SpecificOccupation = p.SpecificOccupation
			if p.WorkOnPlateauStartDate != nil {
				t := *p.WorkOnPlateauStartDate
				item.WorkOnPlateauStartDate = []int{t.Year(), int(t.Month()), t.Day()}
			}
		}
		list = append(list, item)
	}
	return response.NewPage(list, total, page, size), nil
}

func roleFromUserTypeLabel(label string) string {
	switch label {
	case "系统管理员":
		return models.UserRoleAdmin
	case "铁路职工":
		return models.UserRoleRailway
	case "疾控中心工作人员":
		return models.UserRoleCDC
	case "专职医护", "专职医护人员":
		return models.UserRoleMedical
	default:
		return ""
	}
}

func roleDisplayName(role string) string {
	switch role {
	case models.UserRoleAdmin:
		return "系统管理员"
	case models.UserRoleMedical:
		return "专职医护人员"
	case models.UserRoleCDC:
		return "疾控中心工作人员"
	case models.UserRoleRailway:
		return "铁路职工"
	default:
		return role
	}
}

// GetAdminUserDetail returns user + employee profile for the admin drawers.
func GetAdminUserDetail(db *gorm.DB, userID uint) (map[string]any, error) {
	vo, err := GetEmployeeProfileVO(db, userID)
	if err != nil {
		return nil, err
	}
	out := FormatEmployeeProfileVO(vo)
	u, uerr := models.User{}.FindByID(db, userID)
	if uerr != nil {
		return nil, uerr
	}
	idStr := strconv.FormatUint(uint64(u.ID), 10)
	out["userId"] = idStr
	out["id"] = idStr
	out["userType"] = roleDisplayName(u.Role)
	out["role"] = u.Role
	out["name"] = firstNonEmpty(vo.Name, u.Name)
	out["phoneNumber"] = firstNonEmpty(vo.PhoneNumber, u.PhoneNumber, u.Phone)
	// Legacy admin form field aliases (typo casing from old Java UI).
	if v, ok := out["otherPositionName"]; ok {
		out["otherPositionname"] = v
	}
	if v, ok := out["otherPhoneNumber"]; ok {
		out["otherphoneNumber"] = v
	}
	if v, ok := out["emergencyContactName"]; ok {
		out["emergencyContactname"] = v
	}
	if v, ok := out["otherDiseasesName"]; ok {
		out["otherDiseasesname"] = v
	}
	if v, ok := out["otherCancerName"]; ok {
		out["OtherCancername"] = v
	}
	return out, nil
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
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
// When diseaseTypeName is set, returns the full disease weight set (no pagination)
// because the admin drawer needs every slider row at once.
func PageSymptomWeights(db *gorm.DB, q PageQuery, diseaseTypeName string) (*response.Page, error) {
	query := db.Model(&models.SymptomWeighting{})
	if diseaseTypeName != "" {
		query = query.Where("disease_type_name = ?", diseaseTypeName)
		var list []models.SymptomWeighting
		if err := query.Order("symptom_weighting_id ASC").Find(&list).Error; err != nil {
			return nil, err
		}
		return response.NewPage(list, int64(len(list)), 1, len(list)), nil
	}
	page, size := q.Normalize()
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

// ListSymptomWeightsByDisease returns all weight rows for one disease as a flat list.
func ListSymptomWeightsByDisease(db *gorm.DB, diseaseTypeName string) ([]models.SymptomWeighting, error) {
	var list []models.SymptomWeighting
	err := db.Where("disease_type_name = ?", diseaseTypeName).
		Order("symptom_weighting_id ASC").Find(&list).Error
	return list, err
}

// UpdateSymptomWeight updates weight score for one row.
func UpdateSymptomWeight(db *gorm.DB, dto SymptomWeightUpdateDTO) error {
	return db.Model(&models.SymptomWeighting{}).
		Where("symptom_weighting_id = ?", dto.SymptomWeightingID).
		Update("weight_score", dto.WeightScore).Error
}

// FeedbackListItem is the admin feedback table row (Java-compat date array).
type FeedbackListItem struct {
	FeedbackID    uint   `json:"feedbackId"`
	UserID        uint   `json:"userId"`
	Name          string `json:"name"`
	PhoneNumber   string `json:"phoneNumber"`
	FeedbackTitle string `json:"feedbackTitle"`
	FeedbackText  string `json:"feedbackText"`
	FeedbackDate  []int  `json:"feedbackDate"`
}

// PageFeedback lists user feedback entries.
func PageFeedback(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Model(&models.UserFeedback{})
	if q.Name != "" {
		query = query.Where("name LIKE ?", "%"+q.Name+"%")
	}
	if q.PhoneNumber != "" {
		query = query.Where("phone_number LIKE ?", "%"+q.PhoneNumber+"%")
	}
	if q.FeedbackTitle != "" {
		query = query.Where("feedback_title LIKE ?", "%"+q.FeedbackTitle+"%")
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.UserFeedback
	if err := query.Order("feedback_date DESC").
		Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]FeedbackListItem, 0, len(list))
	for _, fb := range list {
		t := fb.FeedbackDate
		out = append(out, FeedbackListItem{
			FeedbackID:    fb.FeedbackID,
			UserID:        fb.UserID,
			Name:          fb.Name,
			PhoneNumber:   fb.PhoneNumber,
			FeedbackTitle: fb.FeedbackTitle,
			FeedbackText:  fb.FeedbackText,
			FeedbackDate: []int{
				t.Year(), int(t.Month()), t.Day(),
				t.Hour(), t.Minute(), t.Second(),
			},
		})
	}
	return response.NewPage(out, total, page, size), nil
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

// DiagnosisListItem joins user/profile for diagnosis table.
type DiagnosisListItem struct {
	DiagnosisResultsID uint   `json:"diagnosisResultsId"`
	UserID             uint   `json:"userId"`
	UserType           string `json:"userType"`
	Name               string `json:"name"`
	PhoneNumber        string `json:"phoneNumber"`
	Gender             string `json:"gender"`
	Age                int    `json:"age"`
	Department         string `json:"department"`
	DiseaseType        string `json:"diseaseType"`
	DiscoveryMethod    string `json:"discoveryMethod"`
	HospitalName       string `json:"hospitalName"`
	DiseaseOutcome     string `json:"diseaseOutcome"`
	AdmissionDate      []int  `json:"admissionDate"`
	DischargeDate      []int  `json:"dischargeDate"`
}

// LabReportListItem joins user/profile for lab table.
type LabReportListItem struct {
	LabTestReportID             uint   `json:"labTestReportId"`
	UserID                      uint   `json:"userId"`
	UserType                    string `json:"userType"`
	Name                        string `json:"name"`
	PhoneNumber                 string `json:"phoneNumber"`
	Gender                      string `json:"gender"`
	Age                         int    `json:"age"`
	Department                  string `json:"department"`
	IsVirusAntigenTestDone      *bool  `json:"isVirusAntigenTestDone"`
	IsVirusNucleicAcidTestDone  *bool  `json:"isVirusNucleicAcidTestDone"`
	IsVirusCultureIsolationDone *bool  `json:"isVirusCultureIsolationDone"`
	IsSerologicalTestDone       *bool  `json:"isSerologicalTestDone"`
	PathogenicTestResults       string `json:"pathogenicTestResults"`
	UploadDate                  []int  `json:"uploadDate"`
}

// PageDiagnosisResults lists diagnosis submissions with user fields.
func PageDiagnosisResults(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Table("diagnosis_results AS d").
		Joins("LEFT JOIN users u ON u.id = d.user_id").
		Joins("LEFT JOIN employee_profiles p ON p.user_id = d.user_id")
	if q.Name != "" {
		query = query.Where("u.name LIKE ?", "%"+q.Name+"%")
	}
	if q.PhoneNumber != "" {
		like := "%" + q.PhoneNumber + "%"
		query = query.Where("u.phone_number LIKE ? OR u.phone LIKE ?", like, like)
	}
	if q.Department != "" {
		query = query.Where("p.department LIKE ?", "%"+q.Department+"%")
	}
	if q.BeginDate != "" {
		if t, err := ParseDate(q.BeginDate); err == nil {
			query = query.Where("d.submission_time >= ?", t)
		}
	}
	if q.EndDate != "" {
		if t, err := ParseDate(q.EndDate); err == nil {
			query = query.Where("d.submission_time < ?", t.AddDate(0, 0, 1))
		}
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	type row struct {
		models.DiagnosisResults
		Name        string `gorm:"column:name"`
		PhoneNumber string `gorm:"column:phone_number"`
		Phone       string `gorm:"column:phone"`
		Role        string `gorm:"column:role"`
		Gender      string `gorm:"column:gender"`
		Age         int    `gorm:"column:age"`
		Department  string `gorm:"column:department"`
	}
	var rows []row
	if err := query.Select(`d.*, u.name, u.phone_number, u.phone, u.role, p.gender, p.age, p.department`).
		Order("d.diagnosis_results_id DESC").
		Offset((page - 1) * size).Limit(size).
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	list := make([]DiagnosisListItem, 0, len(rows))
	for _, r := range rows {
		phone := r.PhoneNumber
		if phone == "" {
			phone = r.Phone
		}
		item := DiagnosisListItem{
			DiagnosisResultsID: r.DiagnosisResultsID,
			UserID:             r.UserID,
			UserType:           roleDisplayName(r.Role),
			Name:               r.Name,
			PhoneNumber:        phone,
			Gender:             r.Gender,
			Age:                r.Age,
			Department:         r.Department,
			DiseaseType:        r.DiseaseType,
			DiscoveryMethod:    r.DiscoveryMethod,
			HospitalName:       r.HospitalName,
			DiseaseOutcome:     r.DiseaseOutcome,
		}
		if r.AdmissionDate != nil {
			item.AdmissionDate = dateParts(*r.AdmissionDate)
		}
		if r.DischargeDate != nil {
			item.DischargeDate = dateParts(*r.DischargeDate)
		}
		list = append(list, item)
	}
	return response.NewPage(list, total, page, size), nil
}

// PageLabReports lists lab reports with user fields.
func PageLabReports(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Table("lab_test_reports AS l").
		Joins("LEFT JOIN users u ON u.id = l.user_id").
		Joins("LEFT JOIN employee_profiles p ON p.user_id = l.user_id")
	if q.Name != "" {
		query = query.Where("u.name LIKE ?", "%"+q.Name+"%")
	}
	if q.PhoneNumber != "" {
		like := "%" + q.PhoneNumber + "%"
		query = query.Where("u.phone_number LIKE ? OR u.phone LIKE ?", like, like)
	}
	if q.Department != "" {
		query = query.Where("p.department LIKE ?", "%"+q.Department+"%")
	}
	if q.BeginDate != "" {
		if t, err := ParseDate(q.BeginDate); err == nil {
			query = query.Where("l.upload_date >= ?", t)
		}
	}
	if q.EndDate != "" {
		if t, err := ParseDate(q.EndDate); err == nil {
			query = query.Where("l.upload_date < ?", t.AddDate(0, 0, 1))
		}
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	type row struct {
		models.LabTestReport
		Name        string `gorm:"column:name"`
		PhoneNumber string `gorm:"column:phone_number"`
		Phone       string `gorm:"column:phone"`
		Role        string `gorm:"column:role"`
		Gender      string `gorm:"column:gender"`
		Age         int    `gorm:"column:age"`
		Department  string `gorm:"column:department"`
	}
	var rows []row
	if err := query.Select(`l.*, u.name, u.phone_number, u.phone, u.role, p.gender, p.age, p.department`).
		Order("l.lab_test_report_id DESC").
		Offset((page - 1) * size).Limit(size).
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	list := make([]LabReportListItem, 0, len(rows))
	for _, r := range rows {
		phone := r.PhoneNumber
		if phone == "" {
			phone = r.Phone
		}
		item := LabReportListItem{
			LabTestReportID:             r.LabTestReportID,
			UserID:                      r.UserID,
			UserType:                    roleDisplayName(r.Role),
			Name:                        r.Name,
			PhoneNumber:                 phone,
			Gender:                      r.Gender,
			Age:                         r.Age,
			Department:                  r.Department,
			IsVirusAntigenTestDone:      r.IsVirusAntigenTestDone,
			IsVirusNucleicAcidTestDone:  r.IsVirusNucleicAcidTestDone,
			IsVirusCultureIsolationDone: r.IsVirusCultureIsolationDone,
			IsSerologicalTestDone:       r.IsSerologicalTestDone,
			PathogenicTestResults:       r.PathogenicTestResults,
		}
		if r.UploadDate != nil {
			item.UploadDate = dateParts(*r.UploadDate)
		}
		list = append(list, item)
	}
	return response.NewPage(list, total, page, size), nil
}
