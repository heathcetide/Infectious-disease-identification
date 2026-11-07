// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"strconv"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/internal/types"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
)

// LabSubmitReport POST /api/v1/lab/report
func (h *Handlers) LabSubmitReport(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var dto services.LabTestReportDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Lab.SaveReport(userID, dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}

// LabGetReport GET /api/v1/lab/report
func (h *Handlers) LabGetReport(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var q types.DateQuery
	_ = c.ShouldBindQuery(&q)
	if q.Date == "" {
		q.Date = time.Now().Format("2006-01-02")
	}
	out, err := h.svc.Lab.GetByDate(userID, q.Date)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, out)
}

// LabUploadFile POST /api/v1/lab/file
func (h *Handlers) LabUploadFile(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	key, name, err := h.uploadMultipartFile(c, "lab")
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	reportID := uint(0)
	if v := c.PostForm("labTestReportId"); v != "" {
		if id, parseErr := strconv.ParseUint(v, 10, 64); parseErr == nil {
			reportID = uint(id)
		}
	}
	if reportID == 0 {
		dto := services.LabTestReportDTO{}
		_ = h.svc.Lab.SaveReport(userID, dto)
		if h.db != nil {
			var report models.LabTestReport
			if err := h.db.Where("user_id = ?", userID).Order("lab_test_report_id DESC").First(&report).Error; err == nil {
				reportID = report.LabTestReportID
			}
		}
	}
	meta := models.LabTestFile{
		LabTestReportID: reportID,
		FileName:        name,
		FilePath:        key,
	}
	if h.db != nil {
		if err := h.db.Create(&meta).Error; err != nil {
			respgin.WriteError(c, response.Err(response.CodeInternal))
			return
		}
	}
	respgin.Created(c, gin.H{
		"file":      meta,
		"key":       key,
		"publicUrl": pkgstorage.PublicURL(key),
	})
}

// InfectLabFile GET /api/v1/infect-files/lab/:id
func (h *Handlers) InfectLabFile(c *gin.Context) {
	var req types.IDPathRequest
	if err := c.ShouldBindUri(&req); err != nil {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	if h.db == nil {
		respgin.WriteError(c, response.Err(response.CodeNotFound))
		return
	}
	var meta models.LabTestFile
	if err := h.db.First(&meta, req.ID).Error; err != nil {
		respgin.WriteError(c, response.Err(response.CodeNotFound))
		return
	}
	h.serveStorageKey(c, meta.FilePath, meta.FileName)
}
