// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// EmployeeProfile holds extended employee/medical profile data split from Java User.
type EmployeeProfile struct {
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint `json:"userId" gorm:"uniqueIndex;not null"`

	Gender                   string     `json:"gender,omitempty" gorm:"size:10"`
	IsPregnant               bool       `json:"isPregnant" gorm:"default:false"`
	PregnancyWeeks           int        `json:"pregnancyWeeks,omitempty"`
	IDNumber                 string     `json:"idNumber,omitempty" gorm:"size:32"`
	BirthDate                *time.Time `json:"birthDate,omitempty" gorm:"type:date"`
	Age                      int        `json:"age,omitempty"`
	Ethnicity                string     `json:"ethnicity,omitempty" gorm:"size:50"`
	EducationLevel           string     `json:"educationLevel,omitempty" gorm:"size:50"`
	WorkOnPlateauStartDate   *time.Time `json:"workOnPlateauStartDate,omitempty" gorm:"type:date"`
	Department               string     `json:"department,omitempty" gorm:"size:100"`
	SpecificOccupation       string     `json:"specificOccupation,omitempty" gorm:"size:100"`
	MedicalPersonnelType     string     `json:"medicalPersonnelType,omitempty" gorm:"size:50"`
	OtherPositionName        string     `json:"otherPositionName,omitempty" gorm:"size:100"`
	OtherPhoneNumber         string     `json:"otherPhoneNumber,omitempty" gorm:"size:20"`
	EmergencyContactName     string     `json:"emergencyContactName,omitempty" gorm:"size:100"`
	EmergencyContactPhone    string     `json:"emergencyContactPhoneNumber,omitempty" gorm:"size:20"`
	EmergencyContactRelation string     `json:"emergencyContactRelation,omitempty" gorm:"size:50"`

	HasMedicalHistory              bool   `json:"hasMedicalHistory" gorm:"default:false"`
	HasHypertension                bool   `json:"hasHypertension" gorm:"default:false"`
	HasDiabetes                    bool   `json:"hasDiabetes" gorm:"default:false"`
	HasHyperlipidemia              bool   `json:"hasHyperlipidemia" gorm:"default:false"`
	HasHyperuricemia               bool   `json:"hasHyperuricemia" gorm:"default:false"`
	HasCoronaryHeartDisease        bool   `json:"hasCoronaryHeartDisease" gorm:"default:false"`
	HasStroke                      bool   `json:"hasStroke" gorm:"default:false"`
	HasOtherCardiovascularDiseases bool   `json:"hasOtherCardiovascularDiseases" gorm:"default:false"`
	HasAsthma                      bool   `json:"hasAsthma" gorm:"default:false"`
	HasCOPD                        bool   `json:"hasCOPD" gorm:"default:false"`
	HasPepticUlcer                 bool   `json:"hasPepticUlcer" gorm:"default:false"`
	HasMalignantTumor              bool   `json:"hasMalignantTumor" gorm:"default:false"`
	HasLungCancer                  bool   `json:"hasLungCancer" gorm:"default:false"`
	HasOtherCancer                 bool   `json:"hasOtherCancer" gorm:"default:false"`
	OtherCancerName                string `json:"otherCancerName,omitempty" gorm:"size:100"`
	HasSevereMentalDisorders       bool   `json:"hasSevereMentalDisorders" gorm:"default:false"`
	HasTuberculosis                bool   `json:"hasTuberculosis" gorm:"default:false"`
	HasHepatitis                   bool   `json:"hasHepatitis" gorm:"default:false"`
	HasOccupationalDisease         bool   `json:"hasOccupationalDisease" gorm:"default:false"`
	HasChronicKidneyDisease        bool   `json:"hasChronicKidneyDisease" gorm:"default:false"`
	HasChronicLiverDisease         bool   `json:"hasChronicLiverDisease" gorm:"default:false"`
	HasImmunodeficiency            bool   `json:"hasImmunodeficiency" gorm:"default:false"`
	HasTyphus                      bool   `json:"hasTyphus" gorm:"default:false"`
	IsPostpartumInSixWeeks         bool   `json:"isPostpartumInSixWeeks" gorm:"default:false"`
	HasDustExposure                bool   `json:"hasDustExposure" gorm:"default:false"`
	HasOtherDiseases               bool   `json:"hasOtherDiseases" gorm:"default:false"`
	OtherDiseasesName              string `json:"otherDiseasesName,omitempty" gorm:"size:200"`

	SmokingStatus  string `json:"smokingStatus,omitempty" gorm:"size:50"`
	DrinkingStatus string `json:"drinkingStatus,omitempty" gorm:"size:50"`
	Height         int    `json:"height,omitempty"`
	Weight         int    `json:"weight,omitempty"`

	IsVaccinatedForCOVID     bool `json:"isVaccinatedForCOVID" gorm:"default:false"`
	IsVaccinatedForFlu       bool `json:"isVaccinatedForFlu" gorm:"default:false"`
	IsVaccinatedForPlague    bool `json:"isVaccinatedForPlague" gorm:"default:false"`
	IsVaccinatedForBCG       bool `json:"isVaccinatedForBCG" gorm:"default:false"`
	IsVaccinatedForHepatitis bool `json:"isVaccinatedForHepatitis" gorm:"default:false"`
}

func (EmployeeProfile) TableName() string { return constants.TableEmployeeProfiles }
