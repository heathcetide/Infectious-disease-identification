// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
)

// CDCAccuracyStrike GET /api/v1/cdc/accuracy/strike
func (h *Handlers) CDCAccuracyStrike(c *gin.Context) {
	start := c.Query("startDate")
	end := c.Query("endDate")
	data, err := services.DiagnosticAccuracyStrike(h.db, start, end)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// CDCAccuracyStrikeLegacy GET /cdcStaff/statisticsDiagnosticAccuracy/strike
func (h *Handlers) CDCAccuracyStrikeLegacy(c *gin.Context) {
	start := c.Query("startDate")
	end := c.Query("endDate")
	data, err := services.DiagnosticAccuracyStrike(h.db, start, end)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, data)
}

// CDCSetPredictorRange PUT /api/v1/cdc/accuracy/predictor-range/:range
func (h *Handlers) CDCSetPredictorRange(c *gin.Context) {
	v, err := strconv.ParseFloat(c.Param("range"), 64)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "invalid range"))
		return
	}
	if err := services.SetPredictorRange(h.db, v); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"predictorRange": v})
}

// CDCSetPredictorRangeLegacy GET /cdcStaff/statisticsDiagnosticAccuracy/JudgmentsPredictorNumber/:range
func (h *Handlers) CDCSetPredictorRangeLegacy(c *gin.Context) {
	v, err := strconv.ParseFloat(c.Param("range"), 64)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "invalid range"))
		return
	}
	if err := services.SetPredictorRange(h.db, v); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, true)
}

// CDCGetPredictorRange GET /api/v1/cdc/accuracy/predictor/:range — store + ack (path sets range)
func (h *Handlers) CDCGetPredictorRange(c *gin.Context) {
	v, err := strconv.ParseFloat(c.Param("range"), 64)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "invalid range"))
		return
	}
	if err := services.SetPredictorRange(h.db, v); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"predictorRange": v})
}
