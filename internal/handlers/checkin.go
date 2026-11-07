// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/internal/types"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
)

// CheckInGeo POST /api/v1/checkin/
func (h *Handlers) CheckInGeo(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var req types.CheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	dto := services.CheckInDTO{
		City:         req.City,
		County:       req.County,
		Health:       req.Health,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		LocationName: req.LocationName,
		Province:     req.Province,
	}
	rec, err := h.svc.CheckIn.CheckIn(userID, dto)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeInternal, err.Error()))
		return
	}
	respgin.Success(c, rec)
}

// CheckInDaily GET /api/v1/checkin/daily
func (h *Handlers) CheckInDaily(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var q types.DateQuery
	_ = c.ShouldBindQuery(&q)
	day := time.Now()
	if q.Date != "" {
		parsed, err := services.ParseDate(q.Date)
		if err != nil {
			respgin.WriteError(c, response.New(response.CodeBadRequest, "invalid date"))
			return
		}
		day = parsed
	}
	vo, err := h.svc.CheckIn.GetDaily(userID, day)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, vo)
}

// CheckInMonth GET /api/v1/checkin/month
func (h *Handlers) CheckInMonth(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var q types.MonthQuery
	_ = c.ShouldBindQuery(&q)
	rows, err := h.svc.CheckIn.GetMonth(userID, q.Month)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, rows)
}

// CheckInSymptoms POST /api/v1/checkin/symptoms
func (h *Handlers) CheckInSymptoms(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var body services.AllSymptoms
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	scores, err := h.svc.CheckIn.SaveSymptomsAndScore(userID, body)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"diseasescoringList": scores})
}

// CheckInEnvironment POST /api/v1/checkin/environment
func (h *Handlers) CheckInEnvironment(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var dto services.WorkEnvironmentDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	env, err := h.svc.CheckIn.SaveWorkEnvironment(userID, dto)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, env)
}

// CheckInEnvironmentUpdate PUT /api/v1/checkin/environment
func (h *Handlers) CheckInEnvironmentUpdate(c *gin.Context) {
	h.CheckInEnvironment(c)
}

// CheckInEnvironmentGet GET /api/v1/checkin/environment
func (h *Handlers) CheckInEnvironmentGet(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var q types.MonthQuery
	_ = c.ShouldBindQuery(&q)
	month := q.Month
	if month == "" {
		month = time.Now().Format("2006-01")
	}
	env, err := h.svc.CheckIn.GetWorkEnvironment(userID, month)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, env)
}
