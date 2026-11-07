// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/heathcetide/infectious-disease/pkg/common/validate"
	"github.com/gin-gonic/gin"
)

// MedicalDiagnosisManager POST /api/v1/medical/diagnosis/page
func (h *Handlers) MedicalDiagnosisManager(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.Medical.PageDiagnosis(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// MedicalLabManager POST /api/v1/medical/lab/page
func (h *Handlers) MedicalLabManager(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.Medical.PageLabReports(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// MedicalChangePassword PUT /api/v1/medical/password
func (h *Handlers) MedicalChangePassword(c *gin.Context) {
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
	if err := h.svc.Medical.ChangePassword(userID, req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"message": "password changed"})
}
