// Copyright (c) 2026 heathcetide. All rights reserved.

package types

// Infect auth (camelCase JSON for mobile clients)

// LoginResponse is the infect-style login success payload.
type LoginResponse struct {
	AccessToken         string          `json:"accessToken"`
	RefreshToken        string          `json:"refreshToken"`
	ExpiresIn           int64           `json:"expiresIn"`
	TokenType           string          `json:"tokenType"`
	User                InfectUserBrief `json:"user"`
	InformationComplete bool            `json:"informationComplete"`
	// Admin frontend aliases
	Token    string `json:"token"`
	UserType string `json:"userType"`
	UserName string `json:"userName"`
}

// RoleDisplayName maps internal role keys to admin UI labels (Chinese).
func RoleDisplayName(role string) string {
	switch role {
	case "admin":
		return "系统管理员"
	case "medical":
		return "专职医护人员"
	case "cdc":
		return "疾控中心工作人员"
	case "railway":
		return "铁路职工"
	case "user":
		return "普通用户"
	default:
		return role
	}
}

// InfectUserBrief is the user object returned on login.
type InfectUserBrief struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

// EmployeePasswordRequest is PUT /api/v1/employee/password.
type EmployeePasswordRequest struct {
	OldPassword     string `json:"oldPassword" binding:"required" validate:"required,min=6,max=128"`
	NewPassword     string `json:"newPassword" binding:"required" validate:"required,min=6,max=128"`
	ConfirmPassword string `json:"confirmPassword" binding:"required" validate:"required,min=6,max=128"`
}

// FeedbackRequest is POST /api/v1/employee/feedback.
type FeedbackRequest struct {
	FeedbackDate  string `json:"feedbackDate"`
	FeedbackText  string `json:"feedbackText"`
	FeedbackTitle string `json:"feedbackTitle"`
	Name          string `json:"name"`
	PhoneNumber   string `json:"phoneNumber"`
}

// SatisfactionRequest is POST /api/v1/employee/satisfaction.
type SatisfactionRequest struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment,omitempty"`
	Month   string `json:"month,omitempty"`
}

// DateQuery is a common ?date= query param.
type DateQuery struct {
	Date string `form:"date"`
}

// MonthQuery is a common ?month= query param (yyyy-MM).
type MonthQuery struct {
	Month string `form:"month"`
}

// CheckInRequest is POST /api/v1/checkin/.
type CheckInRequest struct {
	Health       bool    `json:"health"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	Province     string  `json:"province,omitempty"`
	City         string  `json:"city,omitempty"`
	County       string  `json:"county,omitempty"`
	LocationName string  `json:"locationName,omitempty"`
}

// PageBodyRequest is POST .../page style pagination.
type PageBodyRequest struct {
	PageRequest
	Keyword string `json:"keyword"`
}

// IDPathRequest is a URI id that is not necessarily a user id.
type IDPathRequest struct {
	ID uint `uri:"id" binding:"required"`
}

// FilenamePathRequest is learning file download.
type FilenamePathRequest struct {
	Filename string `uri:"filename" binding:"required"`
}

// AIModeSettings is admin settings/ai-mode payload.
type AIModeSettings struct {
	Enabled bool   `json:"enabled"`
	Mode    string `json:"mode,omitempty"`
}
