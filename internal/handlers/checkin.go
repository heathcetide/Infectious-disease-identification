// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"strconv"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/internal/types"
	"github.com/heathcetide/infectious-disease/pkg/common/geocode"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
)

// CheckInReverseGeocode GET /api/v1/checkin/reverse-geocode
// Free reverse geocoding (no map vendor key): BigDataCloud, then Nominatim fallback.
func (h *Handlers) CheckInReverseGeocode(c *gin.Context) {
	lat, err1 := strconv.ParseFloat(strings.TrimSpace(c.Query("latitude")), 64)
	lon, err2 := strconv.ParseFloat(strings.TrimSpace(c.Query("longitude")), 64)
	if err1 != nil || err2 != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "latitude/longitude required"))
		return
	}

	const perTry = 5 * time.Second
	client := geocode.NewClient(
		geocode.WithProvider(geocode.ProviderBigDataCloud),
		geocode.WithLanguage("zh"),
		geocode.WithUserAgent("infectious-disease/1.0"),
		geocode.WithTimeout(perTry),
	)
	result, err := client.Reverse(lat, lon)
	if err != nil {
		fallback := geocode.NewClient(
			geocode.WithProvider(geocode.ProviderNominatim),
			geocode.WithLanguage("zh"),
			geocode.WithUserAgent("infectious-disease/1.0"),
			geocode.WithTimeout(perTry),
		)
		result, err = fallback.Reverse(lat, lon)
		if err != nil {
			respgin.WriteError(c, response.New(response.CodeInternal, "reverse geocode failed"))
			return
		}
	}

	province := strings.TrimSpace(result.PrincipalSub)
	city := strings.TrimSpace(result.Locality)
	county := strings.TrimSpace(result.SubLocality)
	locationName := strings.TrimSpace(result.DisplayName)
	if locationName == "" {
		parts := make([]string, 0, 3)
		for _, p := range []string{province, city, county} {
			if p == "" {
				continue
			}
			if len(parts) > 0 && parts[len(parts)-1] == p {
				continue
			}
			parts = append(parts, p)
		}
		locationName = strings.Join(parts, "")
	}

	respgin.Success(c, gin.H{
		"latitude":     lat,
		"longitude":    lon,
		"province":     province,
		"city":         city,
		"county":       county,
		"locationName": locationName,
		"provider":     string(result.Provider),
	})
}

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
	result, err := h.svc.CheckIn.SaveSymptomsAndScore(userID, body)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, result)
}

// CheckInConfirmScoring POST /api/v1/checkin/scoring/confirm
func (h *Handlers) CheckInConfirmScoring(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	meta, err := h.svc.CheckIn.ConfirmScoring(userID)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, meta)
}

// CheckInEnvironment POST /api/v1/checkin/environment
func (h *Handlers) CheckInEnvironment(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	var dto services.WorkEnvironmentDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		ae := response.New(response.CodeBadRequest, err.Error())
		ae.Details = map[string]any{"bindError": err.Error()}
		respgin.WriteError(c, ae)
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
