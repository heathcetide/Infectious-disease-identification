// Copyright (c) 2026 heathcetide. All rights reserved.

package models

// AllModels returns every GORM model for AutoMigrate.
func AllModels() []any {
	return []any{
		&User{},
		&EmployeeProfile{},
		&DailyHealthStatus{},
		&GeneralSymptoms{},
		&RespiratorySymptoms{},
		&DigestiveSymptoms{},
		&CirculatorySymptoms{},
		&NeurologicalSymptoms{},
		&LocalSymptoms{},
		&OtherSymptoms{},
		&RiskFactorsAndExposure{},
		&DiseaseScoring{},
		&SymptomWeighting{},
		&AIDiseaseScoring{},
		&AISymptomWeighting{},
		&DiseaseType{},
		&DiseaseSubtype{},
		&SystemSettings{},
		&WorkEnvironmentInfo{},
		&DiagnosisResults{},
		&DiagnosisReport{},
		&DiagnosisPersonalInfo{},
		&DiagnosisComplications{},
		&DiagnosisExaminations{},
		&DiagnosisGeneralSymptoms{},
		&DiagnosisRespiratorySymptoms{},
		&DiagnosisDigestiveSymptoms{},
		&DiagnosisCirculatorySymptoms{},
		&DiagnosisNeurologicalSymptoms{},
		&DiagnosisLocalSymptoms{},
		&DiagnosisOtherSymptoms{},
		&LabTestReport{},
		&LabTestFile{},
		&DiseaseLearningMaterial{},
		&MaterialLearningRecord{},
		&Hospital{},
		&UserFeedback{},
		&SatisfactionRating{},
		&SatisfactionSurvey{},
	}
}
