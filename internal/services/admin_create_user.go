// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/password"
	"gorm.io/gorm"
)

// AdminCreateUserDTO matches the admin "添加用户" drawer payload.
type AdminCreateUserDTO struct {
	UserType    json.RawMessage `json:"userType"`
	Name        string          `json:"name"`
	PhoneNumber string          `json:"phoneNumber"`
	Gender      string          `json:"gender"`
	Ethnicity   string          `json:"ethnicity"`
	Age         json.RawMessage `json:"age"`
	IDNumber    string          `json:"idNumber"`
	Department  string          `json:"department"`
	EducationLevel string       `json:"educationLevel"`
	SmokingStatus  string       `json:"smokingStatus"`
	DrinkingStatus string       `json:"drinkingStatus"`
	SpecificOccupation   string `json:"specificOccupation"`
	MedicalPersonnelType string `json:"medicalPersonnelType"`
	WorkOnPlateauStartDate string `json:"workOnPlateauStartDate"`

	Height json.RawMessage `json:"height"`
	Weight json.RawMessage `json:"weight"`
	PregnancyWeeks json.RawMessage `json:"pregnancyWeeks"`

	IsPregnant             bool `json:"isPregnant"`
	IsPostpartumInSixWeeks bool `json:"isPostpartumInSixWeeks"`
	HasMedicalHistory      bool `json:"hasMedicalHistory"`
	HasHypertension        bool `json:"hasHypertension"`
	HasDiabetes            bool `json:"hasDiabetes"`
	HasHyperlipidemia      bool `json:"hasHyperlipidemia"`
	HasHyperuricemia       bool `json:"hasHyperuricemia"`
	HasCoronaryHeartDisease bool `json:"hasCoronaryHeartDisease"`
	HasStroke              bool `json:"hasStroke"`
	HasOtherCardiovascularDiseases bool `json:"hasOtherCardiovascularDiseases"`
	HasAsthma              bool `json:"hasAsthma"`
	HasCOPD                bool `json:"hasCOPD"`
	HasPepticUlcer         bool `json:"hasPepticUlcer"`
	HasMalignantTumor      bool `json:"hasMalignantTumor"`
	HasLungCancer          bool `json:"hasLungCancer"`
	HasOtherCancer         bool `json:"hasOtherCancer"`
	HasSevereMentalDisorders bool `json:"hasSevereMentalDisorders"`
	HasTuberculosis        bool `json:"hasTuberculosis"`
	HasHepatitis           bool `json:"hasHepatitis"`
	HasOccupationalDisease bool `json:"hasOccupationalDisease"`
	HasChronicKidneyDisease bool `json:"hasChronicKidneyDisease"`
	HasChronicLiverDisease bool `json:"hasChronicLiverDisease"`
	HasImmunodeficiency    bool `json:"hasImmunodeficiency"`
	HasTyphus              bool `json:"hasTyphus"`
	HasDustExposure        bool `json:"hasDustExposure"`
	HasOtherDiseases       bool `json:"hasOtherDiseases"`
	IsVaccinatedForCOVID   bool `json:"isVaccinatedForCOVID"`
	IsVaccinatedForFlu     bool `json:"isVaccinatedForFlu"`
	IsVaccinatedForPlague  bool `json:"isVaccinatedForPlague"`
	IsVaccinatedForBCG     bool `json:"isVaccinatedForBCG"`
	IsVaccinatedForHepatitis bool `json:"isVaccinatedForHepatitis"`

	// Frontend sometimes uses lowercase "name" suffix.
	EmergencyContactName         string `json:"emergencyContactname"`
	EmergencyContactNameAlt      string `json:"emergencyContactName"`
	EmergencyContactPhoneNumber  string `json:"emergencyContactphoneNumber"`
	EmergencyContactPhoneAlt     string `json:"emergencyContactPhoneNumber"`
	EmergencyContactRelation     string `json:"emergencyContactRelation"`
	OtherCancerName             string `json:"otherCancername"`
	OtherCancerNameAlt          string `json:"otherCancerName"`
	OtherDiseasesName            string `json:"otherDiseasesname"`
	OtherDiseasesNameAlt         string `json:"otherDiseasesName"`
	OtherPositionName            string `json:"otherPositionname"`
	OtherPositionNameAlt         string `json:"otherPositionName"`
	OtherPhoneNumber             string `json:"otherphoneNumber"`
	OtherPhoneNumberAlt          string `json:"otherPhoneNumber"`
}

func parseFlexibleInt(raw json.RawMessage) int {
	if len(raw) == 0 || string(raw) == "null" {
		return 0
	}
	var n int
	if err := json.Unmarshal(raw, &n); err == nil {
		return n
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		s = strings.TrimSpace(s)
		if s == "" {
			return 0
		}
		v, _ := strconv.Atoi(s)
		return v
	}
	return 0
}

func mapUserTypeToRole(raw json.RawMessage) (string, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return "", fmt.Errorf("请选择用户类型")
	}
	var n int
	if err := json.Unmarshal(raw, &n); err == nil {
		switch n {
		case 1:
			return models.UserRoleAdmin, nil
		case 2:
			return models.UserRoleRailway, nil
		case 3:
			return models.UserRoleCDC, nil
		case 4:
			return models.UserRoleMedical, nil
		default:
			return "", fmt.Errorf("无效的用户类型: %d", n)
		}
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		s = strings.TrimSpace(s)
		switch s {
		case "系统管理员", "admin", "1":
			return models.UserRoleAdmin, nil
		case "铁路职工", "railway", "2":
			return models.UserRoleRailway, nil
		case "疾控中心工作人员", "cdc", "3":
			return models.UserRoleCDC, nil
		case "专职医护", "专职医护人员", "medical", "4":
			return models.UserRoleMedical, nil
		}
	}
	return "", fmt.Errorf("无效的用户类型")
}

// InitialPasswordForDisplay exposes the same rule used when creating users.
func InitialPasswordForDisplay(idNumber, phone string) string {
	return initialPasswordFromID(idNumber, phone)
}

// initialPasswordFromID uses the last 6 digits of the ID card when available,
// otherwise falls back to the last 6 digits of the phone number.
func initialPasswordFromID(idNumber, phone string) string {
	digits := make([]rune, 0, len(idNumber))
	for _, r := range idNumber {
		if r >= '0' && r <= '9' {
			digits = append(digits, r)
		}
	}
	if len(digits) >= 6 {
		return string(digits[len(digits)-6:])
	}
	phoneDigits := make([]rune, 0, len(phone))
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			phoneDigits = append(phoneDigits, r)
		}
	}
	if len(phoneDigits) >= 6 {
		return string(phoneDigits[len(phoneDigits)-6:])
	}
	return defaultResetPassword
}

func parseDatePtr(s string) *time.Time {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	for _, layout := range []string{"2006-01-02", time.RFC3339, "2006/01/02"} {
		if t, err := time.ParseInLocation(layout, s, time.Local); err == nil {
			return &t
		}
	}
	return nil
}

// AdminCreateStaffUser creates a user + employee profile from the admin drawer.
func AdminCreateStaffUser(db *gorm.DB, dto AdminCreateUserDTO) (*models.User, error) {
	name := strings.TrimSpace(dto.Name)
	phone := strings.TrimSpace(dto.PhoneNumber)
	if name == "" {
		return nil, fmt.Errorf("姓名不能为空")
	}
	if phone == "" {
		return nil, fmt.Errorf("手机号不能为空")
	}
	role, err := mapUserTypeToRole(dto.UserType)
	if err != nil {
		return nil, err
	}

	var existing models.User
	if err := db.Where("phone_number = ? OR phone = ? OR username = ?", phone, phone, phone).
		First(&existing).Error; err == nil {
		return nil, fmt.Errorf("手机号已存在")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	plainPwd := initialPasswordFromID(strings.TrimSpace(dto.IDNumber), phone)
	hash, err := password.Hash(plainPwd, nil)
	if err != nil {
		return nil, err
	}

	email := fmt.Sprintf("%s@local.dev", phone)
	user := models.User{
		Username:    phone,
		Email:       email,
		Name:        name,
		PhoneNumber: phone,
		Phone:       phone,
		Password:    hash,
		Role:        role,
		Status:      models.UserStatusActive,
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	profile := models.EmployeeProfile{
		UserID:                         user.ID,
		Gender:                         dto.Gender,
		IsPregnant:                     dto.IsPregnant,
		PregnancyWeeks:                 parseFlexibleInt(dto.PregnancyWeeks),
		IDNumber:                       strings.TrimSpace(dto.IDNumber),
		Age:                            parseFlexibleInt(dto.Age),
		Ethnicity:                      dto.Ethnicity,
		EducationLevel:                 dto.EducationLevel,
		WorkOnPlateauStartDate:         parseDatePtr(dto.WorkOnPlateauStartDate),
		Department:                     dto.Department,
		SpecificOccupation:             dto.SpecificOccupation,
		MedicalPersonnelType:           dto.MedicalPersonnelType,
		OtherPositionName:              firstNonEmpty(dto.OtherPositionName, dto.OtherPositionNameAlt),
		OtherPhoneNumber:               firstNonEmpty(dto.OtherPhoneNumber, dto.OtherPhoneNumberAlt),
		EmergencyContactName:           firstNonEmpty(dto.EmergencyContactName, dto.EmergencyContactNameAlt),
		EmergencyContactPhone:          firstNonEmpty(dto.EmergencyContactPhoneNumber, dto.EmergencyContactPhoneAlt),
		EmergencyContactRelation:       dto.EmergencyContactRelation,
		HasMedicalHistory:              dto.HasMedicalHistory,
		HasHypertension:                dto.HasHypertension,
		HasDiabetes:                    dto.HasDiabetes,
		HasHyperlipidemia:              dto.HasHyperlipidemia,
		HasHyperuricemia:               dto.HasHyperuricemia,
		HasCoronaryHeartDisease:        dto.HasCoronaryHeartDisease,
		HasStroke:                      dto.HasStroke,
		HasOtherCardiovascularDiseases: dto.HasOtherCardiovascularDiseases,
		HasAsthma:                      dto.HasAsthma,
		HasCOPD:                        dto.HasCOPD,
		HasPepticUlcer:                 dto.HasPepticUlcer,
		HasMalignantTumor:              dto.HasMalignantTumor,
		HasLungCancer:                  dto.HasLungCancer,
		HasOtherCancer:                 dto.HasOtherCancer,
		OtherCancerName:                firstNonEmpty(dto.OtherCancerName, dto.OtherCancerNameAlt),
		HasSevereMentalDisorders:       dto.HasSevereMentalDisorders,
		HasTuberculosis:                dto.HasTuberculosis,
		HasHepatitis:                   dto.HasHepatitis,
		HasOccupationalDisease:         dto.HasOccupationalDisease,
		HasChronicKidneyDisease:        dto.HasChronicKidneyDisease,
		HasChronicLiverDisease:         dto.HasChronicLiverDisease,
		HasImmunodeficiency:            dto.HasImmunodeficiency,
		HasTyphus:                      dto.HasTyphus,
		IsPostpartumInSixWeeks:         dto.IsPostpartumInSixWeeks,
		HasDustExposure:                dto.HasDustExposure,
		HasOtherDiseases:               dto.HasOtherDiseases,
		OtherDiseasesName:              firstNonEmpty(dto.OtherDiseasesName, dto.OtherDiseasesNameAlt),
		SmokingStatus:                  dto.SmokingStatus,
		DrinkingStatus:                 dto.DrinkingStatus,
		Height:                         parseFlexibleInt(dto.Height),
		Weight:                         parseFlexibleInt(dto.Weight),
		IsVaccinatedForCOVID:           dto.IsVaccinatedForCOVID,
		IsVaccinatedForFlu:             dto.IsVaccinatedForFlu,
		IsVaccinatedForPlague:          dto.IsVaccinatedForPlague,
		IsVaccinatedForBCG:             dto.IsVaccinatedForBCG,
		IsVaccinatedForHepatitis:       dto.IsVaccinatedForHepatitis,
	}
	if err := db.Create(&profile).Error; err != nil {
		_ = db.Delete(&user).Error
		return nil, err
	}
	return &user, nil
}
