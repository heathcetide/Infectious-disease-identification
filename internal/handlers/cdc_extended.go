// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CDCCheckInStatsByID GET /api/v1/cdc/checkin/stats/:id
func (h *Handlers) CDCCheckInStatsByID(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.CDC.GetCheckInDetail(id)
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

// CDCCheckInStatsLocation GET /api/v1/cdc/checkin/stats/:id/location
func (h *Handlers) CDCCheckInStatsLocation(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.CDC.GetCheckInLocation(id)
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

// CDCCheckInStatsExport GET /api/v1/cdc/checkin/stats/export
func (h *Handlers) CDCCheckInStatsExport(c *gin.Context) {
	data, err := h.svc.CDC.ExportCheckInsCSV()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "checkins.csv", data)
}

// CDCDashboardStatistics GET /api/v1/cdc/dashboard/statistics
func (h *Handlers) CDCDashboardStatistics(c *gin.Context) {
	data, err := h.svc.CDC.StatisticsData()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCDashboardDiseaseToday GET /api/v1/cdc/dashboard/disease-today
func (h *Handlers) CDCDashboardDiseaseToday(c *gin.Context) {
	data, err := h.svc.CDC.DiseaseToday()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCDashboardUserStations GET /api/v1/cdc/dashboard/user-stations
func (h *Handlers) CDCDashboardUserStations(c *gin.Context) {
	data, err := h.svc.CDC.UserStations()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCDashboardCheckinSearch GET /api/v1/cdc/dashboard/checkin-search
func (h *Handlers) CDCDashboardCheckinSearch(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		q = c.Query("phoneNumber")
	}
	limit := 20
	if n, err := strconv.Atoi(c.DefaultQuery("limit", "20")); err == nil {
		limit = n
	}
	data, err := h.svc.CDC.CheckInSearch(q, limit)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCDashboardCheckinDailySeries POST /api/v1/cdc/dashboard/checkin-daily-series
func (h *Handlers) CDCDashboardCheckinDailySeries(c *gin.Context) {
	var req services.CheckInDailySeriesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	data, err := h.svc.CDC.CheckInDailySeries(req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidDate) {
			respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCDashboardImportantUsers GET /api/v1/cdc/dashboard/important-users
func (h *Handlers) CDCDashboardImportantUsers(c *gin.Context) {
	data, err := h.svc.CDC.ImportantUsers()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCDiseaseStatistics GET /api/v1/cdc/disease/statistics
func (h *Handlers) CDCDiseaseStatistics(c *gin.Context) {
	q := services.DiseaseStatisticsQuery{
		DiseaseTypeName: c.Query("diseaseTypeName"),
		StartDate:       c.Query("startDate"),
		EndDate:         c.Query("endDate"),
	}
	data, err := h.svc.CDC.DiseaseStatistics(q)
	if err != nil {
		if errors.Is(err, services.ErrInvalidDate) {
			respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCGetWorkEnvironment GET /api/v1/cdc/work-environments/:id
func (h *Handlers) CDCGetWorkEnvironment(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.CDC.GetWorkEnvironment(id)
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

// CDCExportWorkEnvironments GET /api/v1/cdc/work-environments/export
func (h *Handlers) CDCExportWorkEnvironments(c *gin.Context) {
	data, err := h.svc.CDC.ExportWorkEnvironmentsCSV()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "work_environments.csv", data)
}

// CDCCreateMaterial POST /api/v1/cdc/materials
// Accepts JSON or multipart/form-data (Java LearningMaterialDTO shape).
func (h *Handlers) CDCCreateMaterial(c *gin.Context) {
	var dto services.LearningMaterialCreateDTO
	ct := c.ContentType()
	var err error
	if strings.Contains(ct, "application/json") {
		err = c.ShouldBindJSON(&dto)
	} else {
		err = c.ShouldBind(&dto)
	}
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	data, err := h.svc.CDC.CreateMaterial(dto)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCGetMaterial GET /api/v1/cdc/materials/:id
func (h *Handlers) CDCGetMaterial(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.CDC.GetMaterial(id)
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

// CDCUpdateMaterial PUT /api/v1/cdc/materials/:id
func (h *Handlers) CDCUpdateMaterial(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	var dto services.LearningMaterialCreateDTO
	ct := c.ContentType()
	var err error
	if strings.Contains(ct, "application/json") {
		err = c.ShouldBindJSON(&dto)
	} else {
		err = c.ShouldBind(&dto)
	}
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	data, err := h.svc.CDC.UpdateMaterial(id, dto)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCMaterialStatus PUT /api/v1/cdc/materials/:id/status
func (h *Handlers) CDCMaterialStatus(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	var body struct {
		Disabled *bool `json:"disabled"`
		Enable   *bool `json:"enable"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	disabled := false
	if body.Disabled != nil {
		disabled = *body.Disabled
	} else if body.Enable != nil {
		disabled = !*body.Enable
	}
	if err := h.svc.CDC.SetMaterialStatus(id, disabled); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"disabled": disabled})
}

// CDCMaterialSequence PUT /api/v1/cdc/materials/:id/sequence
func (h *Handlers) CDCMaterialSequence(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	var body services.MaterialSequenceDirection
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.CDC.MoveMaterialSequence(id, body.Direction); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}
