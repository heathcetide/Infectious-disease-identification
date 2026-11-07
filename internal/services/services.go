// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import "gorm.io/gorm"

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
}

// New wires service stubs with the given database handle (may be nil in tests).
func New(db *gorm.DB) *Services {
	return &Services{
		Employee:  NewEmployeeService(db),
		CheckIn:   NewCheckInService(db),
		Diagnosis: NewDiagnosisService(db),
		Lab:       NewLabService(db),
		Learning:  NewLearningService(db),
		Admin:     NewAdminService(db),
		CDC:       NewCDCService(db),
		Medical:   NewMedicalService(db),
		Profile:   NewProfileService(db),
	}
}
