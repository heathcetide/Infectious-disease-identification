// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"log"
	"strconv"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
)

// AIGetAutoWeight GET /api/v1/ai/auto-weight
func (h *Handlers) AIGetAutoWeight(c *gin.Context) {
	enabled, err := services.GetAutoAIWeightUpdaterEnabled(h.db)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"autoAIWeightUpdater": enabled})
}

// AIGetAutoWeightLegacy GET /ai/getAutoAIWeightUpdater
func (h *Handlers) AIGetAutoWeightLegacy(c *gin.Context) {
	enabled, err := services.GetAutoAIWeightUpdaterEnabled(h.db)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccessBool(c, enabled)
}

// AISetAutoWeight PUT /api/v1/ai/auto-weight
func (h *Handlers) AISetAutoWeight(c *gin.Context) {
	var body struct {
		AutoAIWeightUpdater bool `json:"autoAIWeightUpdater"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		if v := c.Param("enabled"); v != "" {
			body.AutoAIWeightUpdater = v == "true" || v == "1"
		} else {
			respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
			return
		}
	}
	if err := services.SetAutoAIWeightUpdaterEnabled(h.db, body.AutoAIWeightUpdater); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"autoAIWeightUpdater": body.AutoAIWeightUpdater})
}

// AISetAutoWeightLegacy GET /ai/updateAutoAIWeightUpdater/:enabled
func (h *Handlers) AISetAutoWeightLegacy(c *gin.Context) {
	v := c.Param("autoAIWeightUpdater")
	enabled, err := strconv.ParseBool(v)
	if err != nil {
		enabled = v == "1"
	}
	if err := services.SetAutoAIWeightUpdaterEnabled(h.db, enabled); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, true)
}

// AIRunAutoWeight POST /api/v1/ai/auto-weight/run
func (h *Handlers) AIRunAutoWeight(c *gin.Context) {
	if err := services.RunAutoAIWeightUpdater(h.db); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	log.Printf("ai autoweight: run endpoint invoked")
	respgin.Success(c, gin.H{"started": true})
}

// AIRunAutoWeightLegacy GET /ai/startAI
func (h *Handlers) AIRunAutoWeightLegacy(c *gin.Context) {
	_ = services.RunAutoAIWeightUpdater(h.db)
	legacyAdminSuccess(c, true)
}

// AdminWeightsBatchWithDiagnosis POST /api/v1/admin/weights/batch-with-diagnosis
func (h *Handlers) AdminWeightsBatchWithDiagnosis(c *gin.Context) {
	var items []services.SymptomWeightUpdateDTO
	if err := c.ShouldBindJSON(&items); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := services.UpdateSymptomWeightsBatchWithDiagnosis(h.db, items); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"updated": len(items)})
}
