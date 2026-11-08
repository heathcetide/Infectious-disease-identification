// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/password"
	"gorm.io/gorm"
)

// PhoneLogin finds an active user by Phone or PhoneNumber.
func PhoneLogin(db *gorm.DB, phone string) (*models.User, error) {
	var user models.User
	err := db.Where("(phone = ? OR phone_number = ?) AND status = ?", phone, phone, models.UserStatusActive).
		First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// EmployeeProfileVO is the mini-app profile payload including User.name / phone.
type EmployeeProfileVO struct {
	models.EmployeeProfile
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// GetEmployeeProfile loads the extended profile for a user.
func GetEmployeeProfile(db *gorm.DB, userID uint) (*models.EmployeeProfile, error) {
	var profile models.EmployeeProfile
	err := db.Where("user_id = ?", userID).First(&profile).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrProfileNotFound
	}
	return &profile, err
}

// GetEmployeeProfileVO loads profile and merges display fields from users.
// When no profile row exists yet, returns an empty profile VO with name/phone from users.
func GetEmployeeProfileVO(db *gorm.DB, userID uint) (*EmployeeProfileVO, error) {
	vo := &EmployeeProfileVO{EmployeeProfile: models.EmployeeProfile{UserID: userID}}
	profile, err := GetEmployeeProfile(db, userID)
	if err != nil && !errors.Is(err, ErrProfileNotFound) {
		return nil, err
	}
	if profile != nil {
		vo.EmployeeProfile = *profile
	}
	user, uerr := models.User{}.FindByID(db, userID)
	if uerr == nil && user != nil {
		vo.Name = user.Name
		if user.PhoneNumber != "" {
			vo.PhoneNumber = user.PhoneNumber
		} else {
			vo.PhoneNumber = user.Phone
		}
	}
	return vo, nil
}

// FormatEmployeeProfileVO rewrites date fields as YYYY-MM-DD for mini-app display.
func FormatEmployeeProfileVO(vo *EmployeeProfileVO) map[string]any {
	if vo == nil {
		return map[string]any{}
	}
	raw, err := json.Marshal(vo)
	if err != nil {
		return map[string]any{}
	}
	out := map[string]any{}
	if err := json.Unmarshal(raw, &out); err != nil {
		return map[string]any{}
	}
	if vo.WorkOnPlateauStartDate != nil {
		out["workOnPlateauStartDate"] = vo.WorkOnPlateauStartDate.Format("2006-01-02")
	} else {
		out["workOnPlateauStartDate"] = ""
	}
	if vo.BirthDate != nil {
		out["birthDate"] = vo.BirthDate.Format("2006-01-02")
	}
	return out
}

// UpdateEmployeeProfileDTO accepts profile fields plus optional User.name / phone.
type UpdateEmployeeProfileDTO struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	models.EmployeeProfile
}

// UpdateEmployeeProfile creates or updates the employee profile.
// Uses Save so false/0 values from the mini-app are persisted (Updates skips zeros).
// When Name is non-empty, also updates users.name for display after login.
func UpdateEmployeeProfile(db *gorm.DB, userID uint, profile models.EmployeeProfile) (*models.EmployeeProfile, error) {
	return UpdateEmployeeProfileWithName(db, userID, profile, "")
}

// UpdateEmployeeProfileWithName same as UpdateEmployeeProfile, optionally syncing users.name.
func UpdateEmployeeProfileWithName(db *gorm.DB, userID uint, profile models.EmployeeProfile, name string) (*models.EmployeeProfile, error) {
	return UpdateEmployeeProfileFields(db, userID, profile, name, "")
}

// UpdateEmployeeProfileFields syncs profile plus optional users.name / phone_number.
func UpdateEmployeeProfileFields(db *gorm.DB, userID uint, profile models.EmployeeProfile, name, phoneNumber string) (*models.EmployeeProfile, error) {
	profile.UserID = userID
	var existing models.EmployeeProfile
	err := db.Where("user_id = ?", userID).First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Create(&profile).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		profile.ID = existing.ID
		if err := db.Save(&profile).Error; err != nil {
			return nil, err
		}
	}
	updates := map[string]interface{}{}
	if name != "" {
		updates["name"] = name
	}
	if phoneNumber != "" {
		updates["phone_number"] = phoneNumber
		updates["phone"] = phoneNumber
	}
	if len(updates) > 0 {
		_ = db.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
	}
	return &profile, nil
}

// ChangePassword verifies the old password and stores a new hash.
func ChangePassword(db *gorm.DB, userID uint, dto ChangePasswordDTO) error {
	if dto.NewPassword != dto.ConfirmPassword {
		return ErrPasswordMismatch
	}
	user, err := models.User{}.FindByID(db, userID)
	if err != nil {
		return err
	}
	if !password.Verify(dto.OldPassword, user.Password) {
		return ErrInvalidPassword
	}
	hashed, err := password.Hash(dto.NewPassword, nil)
	if err != nil {
		return err
	}
	return models.User{}.UpdatePassword(db, userID, hashed)
}

// SubmitFeedback records user feedback.
func SubmitFeedback(db *gorm.DB, userID uint, dto FeedbackDTO) error {
	name := strings.TrimSpace(dto.Name)
	phone := strings.TrimSpace(dto.PhoneNumber)
	if name == "" || phone == "" {
		var u models.User
		if err := db.Select("name", "phone_number").Where("user_id = ?", userID).First(&u).Error; err == nil {
			if name == "" {
				name = u.Name
			}
			if phone == "" {
				phone = u.PhoneNumber
			}
		}
	}
	fb := models.UserFeedback{
		UserID:        userID,
		Name:          name,
		PhoneNumber:   phone,
		FeedbackTitle: dto.FeedbackTitle,
		FeedbackText:  dto.FeedbackText,
		FeedbackDate:  time.Now(),
	}
	return db.Create(&fb).Error
}
