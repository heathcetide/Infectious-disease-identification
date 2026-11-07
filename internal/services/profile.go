// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// ProfileService handles employee profile completeness and shared profile ops.
type ProfileService struct {
	db *gorm.DB
}

// NewProfileService creates a ProfileService.
func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{db: db}
}

// InformationComplete returns whether the user has filled essential profile fields.
// Gender + ID number are the minimum needed to stop the login "完善信息" prompt;
// department / ethnicity / education remain on the survey form but are not blockers.
func (s *ProfileService) InformationComplete(userID uint) bool {
	if s.db == nil || userID == 0 {
		return false
	}
	var profile models.EmployeeProfile
	if err := s.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return false
	}
	return profile.Gender != "" && profile.IDNumber != ""
}
