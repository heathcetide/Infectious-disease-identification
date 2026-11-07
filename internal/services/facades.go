// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"io"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	"gorm.io/gorm"
)

// EmployeeService wraps employee domain operations.
type EmployeeService struct{ db *gorm.DB }

func NewEmployeeService(db *gorm.DB) *EmployeeService { return &EmployeeService{db: db} }

func (s *EmployeeService) PhoneLogin(phone string) (*models.User, error) {
	return PhoneLogin(s.db, phone)
}
func (s *EmployeeService) GetProfile(userID uint) (*models.EmployeeProfile, error) {
	return GetEmployeeProfile(s.db, userID)
}
func (s *EmployeeService) UpdateProfile(userID uint, p models.EmployeeProfile) (*models.EmployeeProfile, error) {
	return UpdateEmployeeProfile(s.db, userID, p)
}
func (s *EmployeeService) ChangePassword(userID uint, dto ChangePasswordDTO) error {
	return ChangePassword(s.db, userID, dto)
}
func (s *EmployeeService) SubmitFeedback(userID uint, dto FeedbackDTO) error {
	return SubmitFeedback(s.db, userID, dto)
}
func (s *EmployeeService) NeedRating(userID uint) (*NeedRatingVO, error) {
	return GetIsNeedRating(s.db, userID)
}
func (s *EmployeeService) SubmitSatisfaction(userID uint, dto SubmitSatisfactionDTO) error {
	return SubmitSatisfaction(s.db, userID, dto)
}

// CheckInService wraps daily check-in and symptoms.
type CheckInService struct{ db *gorm.DB }

func NewCheckInService(db *gorm.DB) *CheckInService { return &CheckInService{db: db} }

func (s *CheckInService) CheckIn(userID uint, dto CheckInDTO) (*models.DailyHealthStatus, error) {
	return CheckIn(s.db, userID, dto)
}
func (s *CheckInService) GetDaily(userID uint, date time.Time) (*DailyCheckInVO, error) {
	return GetDaily(s.db, userID, date)
}
func (s *CheckInService) GetMonth(userID uint, yearMonth string) ([]MonthlyHealthDay, error) {
	return GetMonth(s.db, userID, yearMonth)
}
func (s *CheckInService) SaveSymptomsAndScore(userID uint, symptoms AllSymptoms) ([]models.DiseaseScoring, error) {
	return SaveSymptomsAndScore(s.db, userID, symptoms)
}
func (s *CheckInService) SaveWorkEnvironment(userID uint, dto WorkEnvironmentDTO) (*models.WorkEnvironmentInfo, error) {
	return SaveWorkEnvironment(s.db, userID, dto)
}
func (s *CheckInService) GetWorkEnvironment(userID uint, yearMonth string) (*models.WorkEnvironmentInfo, error) {
	return GetWorkEnvironment(s.db, userID, yearMonth)
}

// DiagnosisService wraps diagnosis operations.
type DiagnosisService struct{ db *gorm.DB }

func NewDiagnosisService(db *gorm.DB) *DiagnosisService { return &DiagnosisService{db: db} }

func (s *DiagnosisService) Save(userID uint, userType string, dto AllDiagnosisDTO) error {
	return SaveDiagnosis(s.db, userID, userType, dto)
}
func (s *DiagnosisService) GetByDate(userID uint, date string) (*AllDiagnosisDTO, error) {
	return GetDiagnosisByDate(s.db, userID, date)
}
func (s *DiagnosisService) GetByID(id uint) (*DiagnosisDetailVO, error) {
	return GetDiagnosisByID(s.db, id)
}
func (s *DiagnosisService) UpdateByID(id uint, dto AllDiagnosisDTO) error {
	return UpdateDiagnosisByID(s.db, id, dto)
}
func (s *DiagnosisService) SaveForUser(userID uint, dto AllDiagnosisDTO) error {
	return SaveDiagnosisForUser(s.db, userID, dto)
}
func (s *DiagnosisService) ExportCSV() ([]byte, error) {
	return ExportDiagnosisCSV(s.db)
}

// LabService wraps lab test operations.
type LabService struct{ db *gorm.DB }

func NewLabService(db *gorm.DB) *LabService { return &LabService{db: db} }

func (s *LabService) SaveReport(userID uint, dto LabTestReportDTO) error {
	return SaveLabReport(s.db, userID, dto)
}
func (s *LabService) GetByDate(userID uint, date string) (*LabReportWithFiles, error) {
	return GetLabReportByDate(s.db, userID, date)
}
func (s *LabService) LinkFiles(reportID uint, fileIDs []uint) error {
	return LinkLabFiles(s.db, reportID, fileIDs)
}
func (s *LabService) GetByID(id uint) (*LabReportWithFiles, error) {
	return GetLabReportByID(s.db, id)
}
func (s *LabService) UpdateByID(id uint, dto LabTestReportDTO) error {
	return UpdateLabReportByID(s.db, id, dto)
}
func (s *LabService) SaveForUser(userID uint, dto LabTestReportDTO) error {
	return SaveLabReportForUser(s.db, userID, dto)
}
func (s *LabService) ExportCSV() ([]byte, error) {
	return ExportLabCSV(s.db)
}

// LearningService wraps learning materials.
type LearningService struct{ db *gorm.DB }

func NewLearningService(db *gorm.DB) *LearningService { return &LearningService{db: db} }

func (s *LearningService) PageSelect(dto LearningMaterialsPageDTO) (*response.Page, error) {
	return PageSelectLearningMaterials(s.db, dto)
}

// AdminService wraps admin CRUD and stats.
type AdminService struct{ db *gorm.DB }

func NewAdminService(db *gorm.DB) *AdminService { return &AdminService{db: db} }

func (s *AdminService) PageUsers(q PageQuery) (*response.Page, error) {
	return PageUsers(s.db, q)
}
func (s *AdminService) PageHospitals(q PageQuery) (*response.Page, error) {
	return PageHospitals(s.db, q)
}
func (s *AdminService) ListDiseaseTypes() ([]models.DiseaseType, error) {
	return ListDiseaseTypes(s.db)
}
func (s *AdminService) PageSymptomWeights(q PageQuery, disease string) (*response.Page, error) {
	return PageSymptomWeights(s.db, q, disease)
}
func (s *AdminService) UpdateSymptomWeight(dto SymptomWeightUpdateDTO) error {
	return UpdateSymptomWeight(s.db, dto)
}
func (s *AdminService) PageFeedback(q PageQuery) (*response.Page, error) {
	return PageFeedback(s.db, q)
}
func (s *AdminService) DashboardCounts() (*DashboardCounts, error) {
	return GetDashboardCounts(s.db)
}
func (s *AdminService) PageDiagnosis(q PageQuery) (*response.Page, error) {
	return PageDiagnosisResults(s.db, q)
}
func (s *AdminService) PageLabReports(q PageQuery) (*response.Page, error) {
	return PageLabReports(s.db, q)
}
func (s *AdminService) GetAIMode() (AIModeDTO, error) {
	return GetAIModeDTO(s.db)
}
func (s *AdminService) SetAIMode(isAI bool) (AIModeDTO, error) {
	_, err := SetSystemAIMode(s.db, isAI)
	if err != nil {
		return AIModeDTO{}, err
	}
	return GetAIModeDTO(s.db)
}
func (s *AdminService) PageDiseases(q PageQuery) (*response.Page, error) {
	return PageDiseaseTypes(s.db, q)
}
func (s *AdminService) GetHospital(id uint) (*models.Hospital, error) {
	return GetHospitalByID(s.db, id)
}
func (s *AdminService) CreateHospital(h models.Hospital) (*models.Hospital, error) {
	return CreateHospital(s.db, h)
}
func (s *AdminService) UpdateHospital(id uint, h models.Hospital) (*models.Hospital, error) {
	return UpdateHospital(s.db, id, h)
}
func (s *AdminService) DeleteHospital(id uint) error {
	return DeleteHospital(s.db, id)
}
func (s *AdminService) SetUserActive(id uint, active bool) error {
	return SetUserActive(s.db, id, active)
}
func (s *AdminService) ResetUserPassword(id uint) error {
	return ResetUserPassword(s.db, id)
}
func (s *AdminService) UpdateWeightsBatch(items []SymptomWeightUpdateDTO) error {
	return UpdateSymptomWeightsBatch(s.db, items)
}
func (s *AdminService) UpdateWeightsBatchWithDiagnosis(items []SymptomWeightUpdateDTO) error {
	return UpdateSymptomWeightsBatchWithDiagnosis(s.db, items)
}
func (s *AdminService) PageSatisfactionSurveys(page, size int) (*response.Page, error) {
	ap, err := PageSatisfactionSurveys(s.db, page, size)
	if err != nil {
		return nil, err
	}
	return &response.Page{
		List: ap.List, Total: ap.Total, Page: ap.Page, Size: ap.Size, TotalPage: ap.TotalPage,
	}, nil
}
func (s *AdminService) GetAutoAIWeightUpdater() (bool, error) {
	return GetAutoAIWeightUpdaterEnabled(s.db)
}
func (s *AdminService) SetAutoAIWeightUpdater(enabled bool) error {
	return SetAutoAIWeightUpdaterEnabled(s.db, enabled)
}
func (s *AdminService) GetFeedback(id uint) (*models.UserFeedback, error) {
	return GetFeedbackByID(s.db, id)
}
func (s *AdminService) ExportUsersCSV() ([]byte, error) { return ExportUsersCSV(s.db) }
func (s *AdminService) ImportUsersCSV(r io.Reader) (int, error) {
	return ImportUsersCSV(s.db, r)
}
func (s *AdminService) ExportHospitalsCSV() ([]byte, error) { return ExportHospitalsCSV(s.db) }
func (s *AdminService) ImportHospitalsCSV(r io.Reader) (int, error) {
	return ImportHospitalsCSV(s.db, r)
}
func (s *AdminService) ExportDiseasesCSV() ([]byte, error) { return ExportDiseasesCSV(s.db) }
func (s *AdminService) ExportWeightsCSV(disease string) ([]byte, error) {
	return ExportSymptomWeightsCSV(s.db, disease)
}
func (s *AdminService) ImportWeightsCSV(r io.Reader) (int, error) {
	return ImportSymptomWeightsCSV(s.db, r)
}
func (s *AdminService) ExportFeedbackCSV() ([]byte, error) { return ExportFeedbackCSV(s.db) }

// CDCService wraps CDC-facing list endpoints.
type CDCService struct{ db *gorm.DB }

func NewCDCService(db *gorm.DB) *CDCService { return &CDCService{db: db} }

func (s *CDCService) PageLearning(dto LearningMaterialsPageDTO) (*response.Page, error) {
	return CDCPageLearningMaterials(s.db, dto)
}
func (s *CDCService) PageDiagnosis(q PageQuery) (*response.Page, error) {
	return CDCPageDiagnosis(s.db, q)
}
func (s *CDCService) PageLabReports(q PageQuery) (*response.Page, error) {
	return CDCPageLabReports(s.db, q)
}
func (s *CDCService) CheckinStats() (int64, error) { return CDCCheckinStats(s.db) }
func (s *CDCService) ListDiseaseTypes() ([]models.DiseaseType, error) {
	return CDCListDiseaseTypes(s.db)
}
func (s *CDCService) PageCheckIns(q PageQuery) (*response.Page, error) {
	if s.db == nil {
		return response.NewPage([]models.DailyHealthStatus{}, 0, 1, 10), nil
	}
	page, size := q.Normalize()
	query := s.db.Model(&models.DailyHealthStatus{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.DailyHealthStatus
	if err := query.Order("status_id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	return response.NewPage(list, total, page, size), nil
}
func (s *CDCService) PageWorkEnvironments(q PageQuery) (*response.Page, error) {
	if s.db == nil {
		return response.NewPage([]models.WorkEnvironmentInfo{}, 0, 1, 10), nil
	}
	page, size := q.Normalize()
	query := s.db.Model(&models.WorkEnvironmentInfo{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.WorkEnvironmentInfo
	if err := query.Order("environment_id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	return response.NewPage(list, total, page, size), nil
}
func (s *CDCService) GetCheckInDetail(statusID uint) (*CheckInDetailVO, error) {
	return GetCheckInDetail(s.db, statusID)
}
func (s *CDCService) GetCheckInLocation(statusID uint) (*CheckInLocationVO, error) {
	return GetCheckInLocation(s.db, statusID)
}
func (s *CDCService) ExportCheckInsCSV() ([]byte, error) { return ExportCheckInsCSV(s.db) }
func (s *CDCService) StatisticsData() (*CDCStatisticsData, error) {
	return QueryCDCStatisticsData(s.db)
}
func (s *CDCService) DiseaseToday() ([]DiseaseNameCount, error) { return CDCDiseaseToday(s.db) }
func (s *CDCService) UserStations() ([]UserStation, error)       { return CDCUserStations(s.db) }
func (s *CDCService) CheckInSearch(q string, limit int) ([]CheckInSearchItem, error) {
	return CDCCheckInSearch(s.db, q, limit)
}
func (s *CDCService) CheckInDailySeries(req CheckInDailySeriesRequest) ([]CheckInDailyNumber, error) {
	return CDCCheckInDailySeries(s.db, req)
}
func (s *CDCService) ImportantUsers() ([]ImportantUserInfo, error) {
	return CDCImportantUsers(s.db)
}
func (s *CDCService) DiseaseStatistics(q DiseaseStatisticsQuery) ([]DiseaseNameCount, error) {
	return CDCDiseaseStatistics(s.db, q)
}
func (s *CDCService) GetWorkEnvironment(id uint) (*models.WorkEnvironmentInfo, error) {
	return GetWorkEnvironmentByID(s.db, id)
}
func (s *CDCService) ExportWorkEnvironmentsCSV() ([]byte, error) {
	return ExportWorkEnvironmentsCSV(s.db)
}
func (s *CDCService) CreateMaterial(dto LearningMaterialCreateDTO) (*models.DiseaseLearningMaterial, error) {
	return CreateLearningMaterial(s.db, dto)
}
func (s *CDCService) GetMaterial(id uint) (*models.DiseaseLearningMaterial, error) {
	return GetLearningMaterialByID(s.db, id)
}
func (s *CDCService) UpdateMaterial(id uint, dto LearningMaterialCreateDTO) (*models.DiseaseLearningMaterial, error) {
	return UpdateLearningMaterial(s.db, id, dto)
}
func (s *CDCService) SetMaterialStatus(id uint, disabled bool) error {
	return SetLearningMaterialStatus(s.db, id, disabled)
}
func (s *CDCService) MoveMaterialSequence(id uint, direction string) error {
	return MoveLearningMaterialSequence(s.db, id, direction)
}
func (s *CDCService) DiagnosticAccuracyStrike(start, end string) (DiagnosticAccuracyStrikeVO, error) {
	return DiagnosticAccuracyStrike(s.db, start, end)
}
func (s *CDCService) SetPredictorRange(v float64) error {
	return SetPredictorRange(s.db, v)
}

// MedicalService wraps medical staff endpoints.
type MedicalService struct{ db *gorm.DB }

func NewMedicalService(db *gorm.DB) *MedicalService { return &MedicalService{db: db} }

func (s *MedicalService) PageDiagnosis(q PageQuery) (*response.Page, error) {
	return MedicalPageDiagnosis(s.db, q)
}
func (s *MedicalService) PageLabReports(q PageQuery) (*response.Page, error) {
	return MedicalPageLabReports(s.db, q)
}
func (s *MedicalService) PageHospitals(q PageQuery) (*response.Page, error) {
	return MedicalPageHospitals(s.db, q)
}
func (s *MedicalService) CheckinStats(day time.Time) (int64, error) {
	return MedicalCheckinStats(s.db, day)
}
func (s *MedicalService) ChangePassword(userID uint, dto ChangePasswordDTO) error {
	return ChangePassword(s.db, userID, dto)
}
func (s *MedicalService) LookupUsers(dto UserLookupDTO) ([]models.User, error) {
	return MedicalUserLookup(s.db, dto)
}
func (s *MedicalService) GetUserDetail(id uint) (*MedicalUserDetail, error) {
	return GetMedicalUserDetail(s.db, id)
}
