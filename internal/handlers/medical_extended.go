// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
)

// MedicalGetDiagnosis GET /api/v1/medical/diagnosis/:id
func (h *Handlers) MedicalGetDiagnosis(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.Diagnosis.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// MedicalSaveDiagnosis POST /api/v1/medical/diagnosis
func (h *Handlers) MedicalSaveDiagnosis(c *gin.Context) {
	userIDStr := c.Query("userId")
	if userIDStr == "" {
		var wrap struct {
			UserID uint `json:"userId"`
			services.AllDiagnosisDTO
		}
		if err := c.ShouldBindJSON(&wrap); err != nil {
			respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
			return
		}
		if wrap.UserID == 0 {
			respgin.WriteError(c, response.New(response.CodeBadRequest, "userId required"))
			return
		}
		if err := h.svc.Diagnosis.SaveForUser(wrap.UserID, wrap.AllDiagnosisDTO); err != nil {
			respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
			return
		}
		respgin.Success(c, gin.H{"ok": true})
		return
	}
	uid, err := services.ParseOptionalUserIDQuery(userIDStr)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	var dto services.AllDiagnosisDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Diagnosis.SaveForUser(uid, dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}

// MedicalUpdateDiagnosis PUT /api/v1/medical/diagnosis/:id
func (h *Handlers) MedicalUpdateDiagnosis(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	var dto services.AllDiagnosisDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Diagnosis.UpdateByID(id, dto); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}

// MedicalExportDiagnosis GET /api/v1/medical/diagnosis/export
func (h *Handlers) MedicalExportDiagnosis(c *gin.Context) {
	data, err := h.svc.Diagnosis.ExportCSV()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "diagnosis.csv", data)
}

// MedicalGetLab GET /api/v1/medical/lab/:id
func (h *Handlers) MedicalGetLab(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.Lab.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// MedicalUpdateLab PUT /api/v1/medical/lab/:id
func (h *Handlers) MedicalUpdateLab(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	var dto services.LabTestReportDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Lab.UpdateByID(id, dto); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}

// MedicalCreateLab POST /api/v1/medical/lab
func (h *Handlers) MedicalCreateLab(c *gin.Context) {
	uid, err := services.ParseOptionalUserIDQuery(c.Query("userId"))
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	var dto services.LabTestReportDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Lab.SaveForUser(uid, dto); err != nil {
		if errors.Is(err, services.ErrAlreadySubmitted) {
			respgin.WriteError(c, response.New(response.CodeConflict, err.Error()))
			return
		}
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}

// MedicalExportLab GET /api/v1/medical/lab/export
func (h *Handlers) MedicalExportLab(c *gin.Context) {
	data, err := h.svc.Lab.ExportCSV()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "lab.csv", data)
}

// MedicalUserLookup POST /api/v1/medical/users/lookup
func (h *Handlers) MedicalUserLookup(c *gin.Context) {
	var dto services.UserLookupDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	data, err := h.svc.Medical.LookupUsers(dto)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// MedicalGetUser GET /api/v1/medical/users/:id
func (h *Handlers) MedicalGetUser(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.Medical.GetUserDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}
