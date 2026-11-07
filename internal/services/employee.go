// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"errors"
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

// GetEmployeeProfile loads the extended profile for a user.
func GetEmployeeProfile(db *gorm.DB, userID uint) (*models.EmployeeProfile, error) {
	var profile models.EmployeeProfile
	err := db.Where("user_id = ?", userID).First(&profile).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrProfileNotFound
	}
	return &profile, err
}

// UpdateEmployeeProfile creates or updates the employee profile.
func UpdateEmployeeProfile(db *gorm.DB, userID uint, profile models.EmployeeProfile) (*models.EmployeeProfile, error) {
	profile.UserID = userID
	var existing models.EmployeeProfile
	err := db.Where("user_id = ?", userID).First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Create(&profile).Error; err != nil {
			return nil, err
		}
		return &profile, nil
	}
	if err != nil {
		return nil, err
	}
	profile.ID = existing.ID
	if err := db.Model(&existing).Updates(&profile).Error; err != nil {
		return nil, err
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
	fb := models.UserFeedback{
		UserID:        userID,
		Name:          dto.Name,
		PhoneNumber:   dto.PhoneNumber,
		FeedbackTitle: dto.FeedbackTitle,
		FeedbackText:  dto.FeedbackText,
		FeedbackDate:  time.Now(),
	}
	return db.Create(&fb).Error
}
