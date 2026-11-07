// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/internal/types"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DiagnosisSubmitResult POST /api/v1/diagnosis/result
func (h *Handlers) DiagnosisSubmitResult(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var dto services.AllDiagnosisDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	userType := h.userRole(c, userID)
	if err := h.svc.Diagnosis.Save(userID, userType, dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}

// DiagnosisGetReport GET /api/v1/diagnosis/report
func (h *Handlers) DiagnosisGetReport(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var q types.DateQuery
	_ = c.ShouldBindQuery(&q)
	if q.Date == "" {
		q.Date = time.Now().Format("2006-01-02")
	}
	out, err := h.svc.Diagnosis.GetByDate(userID, q.Date)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, out)
}

// DiagnosisUploadFile POST /api/v1/diagnosis/file
func (h *Handlers) DiagnosisUploadFile(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	key, name, err := h.uploadMultipartFile(c, "diagnosis")
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	resultID := uint(0)
	if v := c.PostForm("diagnosisResultsId"); v != "" {
		if id, parseErr := strconv.ParseUint(v, 10, 64); parseErr == nil {
			resultID = uint(id)
		}
	}
	if resultID == 0 && h.db != nil {
		var result models.DiagnosisResults
		day := time.Now()
		err := h.db.Where("user_id = ? AND submission_time = ?", userID, day.Format("2006-01-02")).First(&result).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result = models.DiagnosisResults{UserID: userID, SubmissionTime: &day}
			_ = h.db.Create(&result).Error
		} else if err == nil {
			resultID = result.DiagnosisResultsID
		}
		if resultID == 0 {
			resultID = result.DiagnosisResultsID
		}
	}
	meta := models.DiagnosisReport{
		DiagnosisResultsID: resultID,
		FileName:           name,
		FilePath:           key,
		UploadDate:         time.Now(),
	}
	if h.db != nil {
		if err := h.db.Create(&meta).Error; err != nil {
			respgin.WriteError(c, response.Err(response.CodeInternal))
			return
		}
	}
	respgin.Created(c, gin.H{
		"report":    meta,
		"key":       key,
		"publicUrl": pkgstorage.PublicURL(key),
	})
}

// InfectDiagnosisFile GET /api/v1/infect-files/diagnosis/:id
func (h *Handlers) InfectDiagnosisFile(c *gin.Context) {
	var req types.IDPathRequest
	if err := c.ShouldBindUri(&req); err != nil {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	if h.db == nil {
		respgin.WriteError(c, response.Err(response.CodeNotFound))
		return
	}
	var meta models.DiagnosisReport
	if err := h.db.First(&meta, req.ID).Error; err != nil {
		respgin.WriteError(c, response.Err(response.CodeNotFound))
		return
	}
	h.serveStorageKey(c, meta.FilePath, meta.FileName)
}

func (h *Handlers) userRole(c *gin.Context, userID uint) string {
	if h.db == nil {
		return models.UserRoleRailway
	}
	user, err := models.User{}.FindByID(h.db.WithContext(c.Request.Context()), userID)
	if err != nil {
		return models.UserRoleRailway
	}
	return user.Role
}
