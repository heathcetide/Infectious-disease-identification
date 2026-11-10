// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
)

// EmployeeNearbyHospitals GET /api/v1/employee/hospitals/nearby
func (h *Handlers) EmployeeNearbyHospitals(c *gin.Context) {
	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	lng, _ := strconv.ParseFloat(c.Query("lng"), 64)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	rows, err := services.NearbyHospitals(h.db, lat, lng, limit)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"list": rows, "threshold": services.HighRiskScoreThreshold})
}

// AdminListRiskAlerts GET /api/v1/admin/risk-alerts
func (h *Handlers) AdminListRiskAlerts(c *gin.Context) {
	onlyOpen := c.Query("open") == "1" || c.Query("open") == "true"
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	rows, err := services.ListRiskAlerts(h.db, onlyOpen, limit)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, rows)
}

// AdminAckRiskAlert PUT /api/v1/admin/risk-alerts/:id/ack
func (h *Handlers) AdminAckRiskAlert(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	if err := services.AcknowledgeRiskAlert(h.db, uint(id)); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"ok": true})
}
