// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
)

// CDCCheckInStatsPage POST /api/v1/cdc/checkin/stats/page
func (h *Handlers) CDCCheckInStatsPage(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.CDC.PageCheckIns(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCDashboard GET /api/v1/cdc/dashboard
func (h *Handlers) CDCDashboard(c *gin.Context) {
	counts, err := h.svc.Admin.DashboardCounts()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	date := c.Query("date")
	stats, err := h.svc.CDC.StatisticsData(date)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	diseaseToday, _ := h.svc.CDC.DiseaseToday(date)
	respgin.Success(c, gin.H{
		"counts":        counts,
		"statistics":    stats,
		"diseaseToday":  diseaseToday,
		"checkInsToday": stats.CheckInToday,
	})
}

// CDCDiseaseStats GET /api/v1/cdc/disease/stats
func (h *Handlers) CDCDiseaseStats(c *gin.Context) {
	list, err := h.svc.CDC.ListDiseaseTypes()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, list)
}

// CDCWorkEnvironments GET /api/v1/cdc/work-environments
func (h *Handlers) CDCWorkEnvironments(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.CDC.PageWorkEnvironments(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCMaterialsManager GET /api/v1/cdc/materials
func (h *Handlers) CDCMaterialsManager(c *gin.Context) {
	var dto services.LearningMaterialsPageDTO
	_ = c.ShouldBindQuery(&dto)
	if dto.PageNo < 1 {
		dto.PageNo = 1
	}
	if dto.PageSize < 1 {
		dto.PageSize = 20
	}
	data, err := h.svc.CDC.PageLearning(dto)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}
