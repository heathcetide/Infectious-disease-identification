// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"errors"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// SaveDiagnosis persists diagnosis aggregate and links report file IDs.
func SaveDiagnosis(db *gorm.DB, userID uint, userType string, dto AllDiagnosisDTO) error {
	return db.Transaction(func(tx *gorm.DB) error {
		now := dateOnly(time.Now())
		emp := dto.DiagnosisResultsEmpDTO

		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}

		result := models.DiagnosisResults{
			UserID:                     userID,
			SubmissionTime:             &now,
			SubmissionUserID:           userID,
			SubmissionUserName:         user.Name,
			SubmissionUserType:         userType,
			DiseaseType:                emp.DiseaseType,
			OtherDiseaseName:           emp.OtherDiseaseName,
			PlagueSubtype:              emp.PlagueSubtype,
			AnthraxSubtype:             emp.AnthraxSubtype,
			DiscoveryMethod:            emp.DiscoveryMethod,
			OtherDiscoveryMethodName:   emp.OtherDiscoveryMethodName,
			DiseaseOutcome:             emp.DiseaseOutcome,
			HospitalName:               emp.HospitalName,
			AdmissionSymptomsAndSigns:  emp.AdmissionSymptomsAndSigns,
			RegistrationClassification: emp.RegistrationClassification,
			OtherRegistrationDetails:   emp.OtherRegistrationDetails,
		}
		if t, err := ParseDate(emp.AdmissionDate); err == nil {
			result.AdmissionDate = &t
		}
		if t, err := ParseDate(emp.DischargeDate); err == nil {
			result.DischargeDate = &t
		}
		if emp.DeathDate != "" {
			if t, err := ParseDate(emp.DeathDate); err == nil {
				result.DeathDate = &t
			}
		}

		if err := tx.Create(&result).Error; err != nil {
			return err
		}
		rid := result.DiagnosisResultsID

		personal := dto.Diagnosispersonalinfo.ToModel()
		personal.DiagnosisResultsID = rid
		if err := tx.Create(&personal).Error; err != nil {
			return err
		}

		general := dto.Diagnosisgeneralsymptoms.ToModel()
		general.DiagnosisResultsID = rid
		if err := tx.Create(&general).Error; err != nil {
			return err
		}

		resp := dto.Diagnosisrespiratorysymptoms
		resp.DiagnosisResultsID = rid
		if err := tx.Create(&resp).Error; err != nil {
			return err
		}

		digestive := dto.Diagnosisdigestivesymptoms.ToModel()
		digestive.DiagnosisResultsID = rid
		if err := tx.Create(&digestive).Error; err != nil {
			return err
		}

		circ := dto.Diagnosiscirculatorysymptoms
		circ.DiagnosisResultsID = rid
		if err := tx.Create(&circ).Error; err != nil {
			return err
		}

		neuro := dto.Diagnosisneurologicalsymptoms
		neuro.DiagnosisResultsID = rid
		if err := tx.Create(&neuro).Error; err != nil {
			return err
		}

		local := dto.Diagnosislocalsymptoms
		local.DiagnosisResultsID = rid
		if err := tx.Create(&local).Error; err != nil {
			return err
		}

		other := dto.Diagnosisothersymptoms
		other.DiagnosisResultsID = rid
		if err := tx.Create(&other).Error; err != nil {
			return err
		}

		comp := dto.Diagnosiscomplications
		comp.DiagnosisResultsID = rid
		if err := tx.Create(&comp).Error; err != nil {
			return err
		}

		exam := dto.Diagnosisexaminations
		exam.DiagnosisResultsID = rid
		if err := tx.Create(&exam).Error; err != nil {
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

// GetDiagnosisByDate loads the full diagnosis aggregate for a user on a date.
func GetDiagnosisByDate(db *gorm.DB, userID uint, dateStr string) (*AllDiagnosisDTO, error) {
	day, err := ParseDate(dateStr)
	if err != nil {
		return nil, ErrInvalidDate
	}

	var result models.DiagnosisResults
	err = db.Where("user_id = ? AND submission_time = ?", userID, day).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &AllDiagnosisDTO{}, nil
	}
	if err != nil {
		return nil, err
	}
	rid := result.DiagnosisResultsID

	out := &AllDiagnosisDTO{
		DiagnosisResultsEmpDTO: DiagnosisResultsEmpDTO{
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
		},
	}
	if result.AdmissionDate != nil {
		out.DiagnosisResultsEmpDTO.AdmissionDate = result.AdmissionDate.Format("2006-01-02")
	}
	if result.DischargeDate != nil {
		out.DiagnosisResultsEmpDTO.DischargeDate = result.DischargeDate.Format("2006-01-02")
	}
	if result.DeathDate != nil {
		out.DiagnosisResultsEmpDTO.DeathDate = result.DeathDate.Format("2006-01-02")
	}

	loadChild := func(dest any) error {
		return db.Where("diagnosis_results_id = ?", rid).First(dest).Error
	}
	var pers models.DiagnosisPersonalInfo
	if loadChild(&pers) == nil {
		out.Diagnosispersonalinfo = DiagnosisPersonalFromModel(pers)
	}
	var gen models.DiagnosisGeneralSymptoms
	if loadChild(&gen) == nil {
		out.Diagnosisgeneralsymptoms = DiagnosisGeneralFromModel(gen)
	}
	_ = loadChild(&out.Diagnosisrespiratorysymptoms)
	var dig models.DiagnosisDigestiveSymptoms
	if loadChild(&dig) == nil {
		out.Diagnosisdigestivesymptoms = DiagnosisDigestiveFromModel(dig)
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
	return out, nil
}
