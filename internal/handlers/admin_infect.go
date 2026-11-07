// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
)

func bindPageQuery(c *gin.Context) services.PageQuery {
	var q services.PageQuery
	if err := c.ShouldBindJSON(&q); err != nil {
		_ = c.ShouldBindQuery(&q)
	}
	return q
}

// AdminUsersPage POST /api/v1/admin/users/page
func (h *Handlers) AdminUsersPage(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.Admin.PageUsers(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminListHospitals GET /api/v1/admin/hospitals
func (h *Handlers) AdminListHospitals(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.Admin.PageHospitals(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminListDiseases GET /api/v1/admin/diseases
func (h *Handlers) AdminListDiseases(c *gin.Context) {
	list, err := h.svc.Admin.ListDiseaseTypes()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, list)
}

// AdminListWeights GET /api/v1/admin/weights
func (h *Handlers) AdminListWeights(c *gin.Context) {
	q := bindPageQuery(c)
	disease := c.Query("diseaseTypeName")
	data, err := h.svc.Admin.PageSymptomWeights(q, disease)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminListFeedback GET /api/v1/admin/feedback
func (h *Handlers) AdminListFeedback(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.Admin.PageFeedback(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminGetAIMode GET /api/v1/admin/settings/ai-mode
func (h *Handlers) AdminGetAIMode(c *gin.Context) {
	data, err := h.svc.Admin.GetAIMode()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminSetAIMode PUT /api/v1/admin/settings/ai-mode
func (h *Handlers) AdminSetAIMode(c *gin.Context) {
	var body services.AIModeDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	data, err := h.svc.Admin.SetAIMode(body.IsAI)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminDashboard GET /api/v1/admin/dashboard — optional aggregate
func (h *Handlers) AdminDashboard(c *gin.Context) {
	counts, err := h.svc.Admin.DashboardCounts()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, counts)
}
