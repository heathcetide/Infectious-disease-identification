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

func (h *Handlers) bindAdminPage(c *gin.Context) (page, size int, surveyID uint) {
	var p adminPageParams
	_ = c.ShouldBindQuery(&p)
	_ = c.ShouldBindJSON(&p)
	page, size = p.normalized()
	surveyID = p.SurveyID
	return page, size, surveyID
}

// AdminCreateSatisfactionSurvey POST /api/v1/admin/satisfaction/surveys
func (h *Handlers) AdminCreateSatisfactionSurvey(c *gin.Context) {
	var body services.SatisfactionSurveyCreateDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	row, err := services.CreateSatisfactionSurvey(h.db, body)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, row)
}

// AdminCreateSatisfactionSurveyLegacy POST /admin/satisfaction/addSatisfactionSurvey
func (h *Handlers) AdminCreateSatisfactionSurveyLegacy(c *gin.Context) {
	var body services.SatisfactionSurveyCreateDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	_, err := services.CreateSatisfactionSurvey(h.db, body)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, true)
}

// AdminUpdateSatisfactionSurvey PUT /api/v1/admin/satisfaction/surveys/:id
func (h *Handlers) AdminUpdateSatisfactionSurvey(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	var body services.SatisfactionSurveyCreateDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	row, err := services.UpdateSatisfactionSurvey(h.db, id, body)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, row)
}

// AdminUpdateSatisfactionLegacy POST /admin/satisfaction/updateSatisfaction
func (h *Handlers) AdminUpdateSatisfactionLegacy(c *gin.Context) {
	var body services.SatisfactionSurveyCreateDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if body.SurveyID == 0 {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "surveyID required"))
		return
	}
	_, err := services.UpdateSatisfactionSurvey(h.db, body.SurveyID, body)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, true)
}

// AdminOpenSatisfactionSurvey PUT /api/v1/admin/satisfaction/surveys/:id/open
func (h *Handlers) AdminOpenSatisfactionSurvey(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	if err := services.OpenSatisfactionSurvey(h.db, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"surveyID": id, "isOpen": true})
}

// AdminCloseSatisfactionSurvey PUT /api/v1/admin/satisfaction/surveys/:id/close
func (h *Handlers) AdminCloseSatisfactionSurvey(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	if err := services.CloseSatisfactionSurvey(h.db, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"surveyID": id, "isOpen": false})
}

// AdminOpenSatisfactionLegacy GET /admin/satisfaction/openSatisfaction/:surveyID
func (h *Handlers) AdminOpenSatisfactionLegacy(c *gin.Context) {
	id, ok := h.pathUint(c, "surveyID")
	if !ok {
		return
	}
	if err := services.OpenSatisfactionSurvey(h.db, id); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, true)
}

// AdminCloseSatisfactionLegacy GET /admin/satisfaction/closeSatisfaction/:surveyID
func (h *Handlers) AdminCloseSatisfactionLegacy(c *gin.Context) {
	id, ok := h.pathUint(c, "surveyID")
	if !ok {
		return
	}
	if err := services.CloseSatisfactionSurvey(h.db, id); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, true)
}

// AdminSatisfactionStats GET /api/v1/admin/satisfaction/surveys/:id/stats
func (h *Handlers) AdminSatisfactionStats(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	stats, err := services.GetSatisfactionStats(h.db, id)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, stats)
}

// AdminSatisfactionStatsLegacy GET /admin/satisfaction/getNumberInfo/:surveyID
func (h *Handlers) AdminSatisfactionStatsLegacy(c *gin.Context) {
	id, ok := h.pathUint(c, "surveyID")
	if !ok {
		return
	}
	stats, err := services.GetSatisfactionStats(h.db, id)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, stats)
}

// AdminPageSatisfactionSurveys GET /api/v1/admin/satisfaction/surveys
func (h *Handlers) AdminPageSatisfactionSurveys(c *gin.Context) {
	page, size, _ := h.bindAdminPage(c)
	data, err := services.PageSatisfactionSurveys(h.db, page, size)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminPageSatisfactionSurveysLegacy GET /admin/satisfaction/pageSelectSatisfactionSurvey
func (h *Handlers) AdminPageSatisfactionSurveysLegacy(c *gin.Context) {
	page, size, _ := h.bindAdminPage(c)
	data, err := services.PageSatisfactionSurveys(h.db, page, size)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, data)
}

// AdminPageSatisfactionRatings GET /api/v1/admin/satisfaction/ratings
func (h *Handlers) AdminPageSatisfactionRatings(c *gin.Context) {
	page, size, surveyID := h.bindAdminPage(c)
	data, err := services.PageSatisfactionRatings(h.db, surveyID, page, size)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminPageSatisfactionRatingsLegacy GET /admin/satisfaction/pageSelectSatisfactionRating
func (h *Handlers) AdminPageSatisfactionRatingsLegacy(c *gin.Context) {
	page, size, surveyID := h.bindAdminPage(c)
	data, err := services.PageSatisfactionRatings(h.db, surveyID, page, size)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	legacyAdminSuccess(c, data)
}
