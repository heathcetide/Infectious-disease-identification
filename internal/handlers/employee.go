// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/heathcetide/infectious-disease/pkg/common/validate"
)

// EmployeeGetProfile GET /api/v1/employee/profile
func (h *Handlers) EmployeeGetProfile(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	profile, err := h.svc.Employee.GetProfileVO(userID)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, services.FormatEmployeeProfileVO(profile))
}

// EmployeeUpdateProfile PUT /api/v1/employee/profile
func (h *Handlers) EmployeeUpdateProfile(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var req services.UpdateEmployeeProfileDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	out, err := h.svc.Employee.UpdateProfileFields(userID, req.EmployeeProfile, req.Name, req.PhoneNumber)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, out)
}

// EmployeeChangePassword PUT /api/v1/employee/password
func (h *Handlers) EmployeeChangePassword(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var req services.ChangePasswordDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := validate.Validate(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Employee.ChangePassword(userID, req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"message": "password changed"})
}

// EmployeeSubmitFeedback POST /api/v1/employee/feedback
func (h *Handlers) EmployeeSubmitFeedback(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var req services.FeedbackDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Employee.SubmitFeedback(userID, req); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}

// EmployeeNeedRating GET /api/v1/employee/satisfaction/need-rating
func (h *Handlers) EmployeeNeedRating(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	vo, err := h.svc.Employee.NeedRating(userID)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, vo)
}

// EmployeeSubmitSatisfaction POST /api/v1/employee/satisfaction
func (h *Handlers) EmployeeSubmitSatisfaction(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var req services.SubmitSatisfactionDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Employee.SubmitSatisfaction(userID, req); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}
