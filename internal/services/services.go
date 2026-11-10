// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/llm"
)

// Services holds domain service stubs for infect APIs.
type Services struct {
	Employee  *EmployeeService
	CheckIn   *CheckInService
	Diagnosis *DiagnosisService
	Lab       *LabService
	Learning  *LearningService
	Admin     *AdminService
	CDC       *CDCService
	Medical   *MedicalService
	Profile   *ProfileService
	LLM       *llm.Client
}

// New wires service stubs with the given database handle (may be nil in tests).
func New(db *gorm.DB, chat *llm.Client) *Services {
	return &Services{
		Employee:  NewEmployeeService(db),
		CheckIn:   NewCheckInService(db, chat),
		Diagnosis: NewDiagnosisService(db),
		Lab:       NewLabService(db),
		Learning:  NewLearningService(db),
		Admin:     NewAdminService(db),
		CDC:       NewCDCService(db),
		Medical:   NewMedicalService(db),
		Profile:   NewProfileService(db),
		LLM:       chat,
	}
}
