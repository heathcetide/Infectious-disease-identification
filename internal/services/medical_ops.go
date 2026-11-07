// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"bytes"
	"encoding/csv"
	"errors"
	"strconv"
	"strings"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// DiagnosisDetailVO wraps aggregate with result row.
type DiagnosisDetailVO struct {
	Result models.DiagnosisResults `json:"result"`
	AllDiagnosisDTO
}

// GetDiagnosisByID loads full diagnosis by result ID.
func GetDiagnosisByID(db *gorm.DB, id uint) (*DiagnosisDetailVO, error) {
	var result models.DiagnosisResults
	if err := db.First(&result, id).Error; err != nil {
		return nil, err
	}
	agg := loadDiagnosisChildren(db, id, &AllDiagnosisDTO{}, result)
	return &DiagnosisDetailVO{Result: result, AllDiagnosisDTO: *agg}, nil
}

func loadDiagnosisChildren(db *gorm.DB, rid uint, out *AllDiagnosisDTO, result models.DiagnosisResults) *AllDiagnosisDTO {
	if out == nil {
		out = &AllDiagnosisDTO{}
	}
	out.DiagnosisResultsEmpDTO = DiagnosisResultsEmpDTO{
		DiseaseType:                result.DiseaseType,
		OtherDiseaseName:           result.OtherDiseaseName,
		PlagueSubtype:              result.PlagueSubtype,
		AnthraxSubtype:             result.AnthraxSubtype,
		DiscoveryMethod:            result.DiscoveryMethod,
		OtherDiscoveryMethodName:   result.OtherDiscoveryMethodName,
		DiseaseOutcome:             result.DiseaseOutcome,
		HospitalName:               result.HospitalName,
		AdmissionSymptomsAndSigns:  result.AdmissionSymptomsAndSigns,
		RegistrationClassification: result.RegistrationClassification,
		OtherRegistrationDetails:   result.OtherRegistrationDetails,
	}
	loadChild := func(dest any) error {
		return db.Where("diagnosis_results_id = ?", rid).First(dest).Error
	}
	_ = loadChild(&out.Diagnosispersonalinfo)
	_ = loadChild(&out.Diagnosisgeneralsymptoms)
	_ = loadChild(&out.Diagnosisrespiratorysymptoms)
	var dig models.DiagnosisDigestiveSymptoms
	if loadChild(&dig) == nil {
		out.Diagnosisdigestivesymptoms = DiagnosisDigestiveDTO{DiagnosisDigestiveSymptoms: dig}
	}
	_ = loadChild(&out.Diagnosiscirculatorysymptoms)
	_ = loadChild(&out.Diagnosisneurologicalsymptoms)
	_ = loadChild(&out.Diagnosislocalsymptoms)
	_ = loadChild(&out.Diagnosisothersymptoms)
	_ = loadChild(&out.Diagnosiscomplications)
	_ = loadChild(&out.Diagnosisexaminations)
	var reportIDs []uint
	if err := db.Model(&models.DiagnosisReport{}).
		Where("diagnosis_results_id = ?", rid).
		Pluck("report_id", &reportIDs).Error; err == nil {
		out.FileIDList = reportIDs
	}
	return out
}

// UpdateDiagnosisByID replaces child rows and updates result header.
func UpdateDiagnosisByID(db *gorm.DB, id uint, dto AllDiagnosisDTO) error {
	var result models.DiagnosisResults
	if err := db.First(&result, id).Error; err != nil {
		return err
	}
	return db.Transaction(func(tx *gorm.DB) error {
		emp := dto.DiagnosisResultsEmpDTO
		updates := map[string]any{
			"disease_type":                  emp.DiseaseType,
			"other_disease_name":            emp.OtherDiseaseName,
			"plague_subtype":                emp.PlagueSubtype,
			"anthrax_subtype":               emp.AnthraxSubtype,
			"discovery_method":              emp.DiscoveryMethod,
			"other_discovery_method_name":   emp.OtherDiscoveryMethodName,
			"disease_outcome":               emp.DiseaseOutcome,
			"hospital_name":                 emp.HospitalName,
			"admission_symptoms_and_signs":  emp.AdmissionSymptomsAndSigns,
			"registration_classification":   emp.RegistrationClassification,
			"other_registration_details":    emp.OtherRegistrationDetails,
		}
		if t, err := ParseDate(emp.AdmissionDate); err == nil {
			updates["admission_date"] = t
		}
		if t, err := ParseDate(emp.DischargeDate); err == nil {
			updates["discharge_date"] = t
		}
		if emp.DeathDate != "" {
			if t, err := ParseDate(emp.DeathDate); err == nil {
				updates["death_date"] = t
			}
		}
		if err := tx.Model(&result).Updates(updates).Error; err != nil {
			return err
		}
		rid := id
		replaceChild := func(model any, create func() error) error {
			if err := tx.Where("diagnosis_results_id = ?", rid).Delete(model).Error; err != nil {
				return err
			}
			return create()
		}
		p := dto.Diagnosispersonalinfo
		p.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisPersonalInfo{}, func() error { return tx.Create(&p).Error }); err != nil {
			return err
		}
		g := dto.Diagnosisgeneralsymptoms
		g.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisGeneralSymptoms{}, func() error { return tx.Create(&g).Error }); err != nil {
			return err
		}
		rsp := dto.Diagnosisrespiratorysymptoms
		rsp.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisRespiratorySymptoms{}, func() error { return tx.Create(&rsp).Error }); err != nil {
			return err
		}
		d := dto.Diagnosisdigestivesymptoms.ToModel()
		d.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisDigestiveSymptoms{}, func() error { return tx.Create(&d).Error }); err != nil {
			return err
		}
		c := dto.Diagnosiscirculatorysymptoms
		c.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisCirculatorySymptoms{}, func() error { return tx.Create(&c).Error }); err != nil {
			return err
		}
		n := dto.Diagnosisneurologicalsymptoms
		n.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisNeurologicalSymptoms{}, func() error { return tx.Create(&n).Error }); err != nil {
			return err
		}
		l := dto.Diagnosislocalsymptoms
		l.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisLocalSymptoms{}, func() error { return tx.Create(&l).Error }); err != nil {
			return err
		}
		o := dto.Diagnosisothersymptoms
		o.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisOtherSymptoms{}, func() error { return tx.Create(&o).Error }); err != nil {
			return err
		}
		comp := dto.Diagnosiscomplications
		comp.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisComplications{}, func() error { return tx.Create(&comp).Error }); err != nil {
			return err
		}
		ex := dto.Diagnosisexaminations
		ex.DiagnosisResultsID = rid
		if err := replaceChild(&models.DiagnosisExaminations{}, func() error { return tx.Create(&ex).Error }); err != nil {
			return err
		}
		for _, fileID := range dto.FileIDList {
			if err := tx.Model(&models.DiagnosisReport{}).
				Where("report_id = ?", fileID).
				Update("diagnosis_results_id", rid).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// SaveDiagnosisForUser saves a new diagnosis for a target user (medical staff).
func SaveDiagnosisForUser(db *gorm.DB, userID uint, dto AllDiagnosisDTO) error {
	return SaveDiagnosis(db, userID, models.UserRoleMedical, dto)
}

// ExportDiagnosisCSV exports diagnosis results.
func ExportDiagnosisCSV(db *gorm.DB) ([]byte, error) {
	var list []models.DiagnosisResults
	if err := db.Order("diagnosis_results_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"diagnosisResultsId", "userId", "submissionTime", "diseaseType", "hospitalName", "diseaseOutcome"})
	for _, d := range list {
		sub := ""
		if d.SubmissionTime != nil {
			sub = d.SubmissionTime.Format("2006-01-02")
		}
		_ = w.Write([]string{
			strconv.FormatUint(uint64(d.DiagnosisResultsID), 10),
			strconv.FormatUint(uint64(d.UserID), 10),
			sub, d.DiseaseType, d.HospitalName, d.DiseaseOutcome,
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

// GetLabReportByID loads lab report and files.
func GetLabReportByID(db *gorm.DB, id uint) (*LabReportWithFiles, error) {
	var report models.LabTestReport
	if err := db.First(&report, id).Error; err != nil {
		return nil, err
	}
	var files []models.LabTestFile
	if err := db.Where("lab_test_report_id = ?", id).Find(&files).Error; err != nil {
		return nil, err
	}
	return &LabReportWithFiles{Report: report, Files: files}, nil
}

// UpdateLabReportByID updates lab report fields.
func UpdateLabReportByID(db *gorm.DB, id uint, dto LabTestReportDTO) error {
	updates := map[string]any{
		"pathogenic_test_results":        dto.PathogenicTestResults,
		"is_virus_antigen_test_done":       dto.IsVirusAntigenTestDone,
		"is_virus_nucleic_acid_test_done":  dto.IsVirusNucleicAcidTestDone,
		"is_virus_culture_isolation_done":  dto.IsVirusCultureIsolationDone,
		"is_serological_test_done":         dto.IsSerologicalTestDone,
	}
	res := db.Model(&models.LabTestReport{}).Where("lab_test_report_id = ?", id).Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return LinkLabFiles(db, id, dto.LabTestFileIDs)
}

// SaveLabReportForUser creates today's lab report for userID (medical entry).
func SaveLabReportForUser(db *gorm.DB, userID uint, dto LabTestReportDTO) error {
	return SaveLabReport(db, userID, dto)
}

// ExportLabCSV exports lab reports.
func ExportLabCSV(db *gorm.DB) ([]byte, error) {
	var list []models.LabTestReport
	if err := db.Order("lab_test_report_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"labTestReportId", "userId", "uploadDate", "pathogenicTestResults"})
	for _, r := range list {
		up := ""
		if r.UploadDate != nil {
			up = r.UploadDate.Format("2006-01-02")
		}
		_ = w.Write([]string{
			strconv.FormatUint(uint64(r.LabTestReportID), 10),
			strconv.FormatUint(uint64(r.UserID), 10),
			up, r.PathogenicTestResults,
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

// UserLookupDTO search body.
type UserLookupDTO struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	PhoneNumber string `json:"phoneNumber"`
	IDNumber  string `json:"idNumber"`
}

// MedicalUserLookup searches users by identifiers.
func MedicalUserLookup(db *gorm.DB, dto UserLookupDTO) ([]models.User, error) {
	q := db.Model(&models.User{})
	phone := dto.Phone
	if phone == "" {
		phone = dto.PhoneNumber
	}
	if phone != "" {
		like := "%" + phone + "%"
		q = q.Where("phone_number LIKE ? OR phone LIKE ?", like, like)
	}
	if dto.Name != "" {
		q = q.Where("name LIKE ?", "%"+dto.Name+"%")
	}
	if dto.IDNumber != "" {
		var userIDs []uint
		if err := db.Model(&models.EmployeeProfile{}).
			Where("id_number LIKE ?", "%"+dto.IDNumber+"%").
			Pluck("user_id", &userIDs).Error; err != nil {
			return nil, err
		}
		if len(userIDs) == 0 {
			return []models.User{}, nil
		}
		q = q.Where("id IN ?", userIDs)
	}
	var users []models.User
	if err := q.Limit(50).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// MedicalUserDetail combines user and profile.
type MedicalUserDetail struct {
	User    models.User             `json:"user"`
	Profile *models.EmployeeProfile `json:"profile,omitempty"`
}

// GetMedicalUserDetail loads user by id with profile.
func GetMedicalUserDetail(db *gorm.DB, id uint) (*MedicalUserDetail, error) {
	u, err := models.User{}.FindByID(db, id)
	if err != nil {
		return nil, err
	}
	detail := &MedicalUserDetail{User: *u}
	var profile models.EmployeeProfile
	if err := db.Where("user_id = ?", id).First(&profile).Error; err == nil {
		detail.Profile = &profile
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return detail, nil
}

// ParseOptionalUserIDQuery reads userId from query string.
func ParseOptionalUserIDQuery(s string) (uint, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, errors.New("userId required")
	}
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("invalid userId")
	}
	return uint(id), nil
}
