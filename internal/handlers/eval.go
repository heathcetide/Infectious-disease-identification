// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
)

// AdminCreateScoringVersion POST /api/v1/admin/eval/versions
func (h *Handlers) AdminCreateScoringVersion(c *gin.Context) {
	var body struct {
		Name       string `json:"name"`
		Notes      string `json:"notes"`
		PromptText string `json:"promptText"`
	}
	_ = c.ShouldBindJSON(&body)
	if body.PromptText == "" {
		body.PromptText = "default hybrid scoring prompt"
	}
	v, err := services.CreateScoringVersion(h.db, body.Name, body.Notes, body.PromptText)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Created(c, v)
}

// AdminListScoringVersions GET /api/v1/admin/eval/versions
func (h *Handlers) AdminListScoringVersions(c *gin.Context) {
	rows, err := services.ListScoringVersions(h.db, 30)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, rows)
}

// AdminGetScoringVersion GET /api/v1/admin/eval/versions/:id
func (h *Handlers) AdminGetScoringVersion(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	v, err := services.GetScoringVersion(h.db, uint(id))
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	payload, err := services.ParseScoringSnapshot(v)
	if err != nil || payload == nil {
		payload = &services.ScoringSnapshotPayload{Mode: v.Mode, PromptText: v.PromptText, Notes: v.Notes}
	}
	respgin.Success(c, gin.H{
		"version":  v,
		"snapshot": payload,
		"summary": gin.H{
			"expertWeightCount": len(payload.ExpertWeights),
			"aiWeightCount":     len(payload.AIWeights),
			"restorable":        len(payload.ExpertWeights)+len(payload.AIWeights) > 0,
		},
	})
}

// AdminDownloadScoringVersion GET /api/v1/admin/eval/versions/:id/download
func (h *Handlers) AdminDownloadScoringVersion(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	v, err := services.GetScoringVersion(h.db, uint(id))
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	payload, err := services.ParseScoringSnapshot(v)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	body, err := json.MarshalIndent(gin.H{
		"versionId":   v.VersionID,
		"name":        v.Name,
		"mode":        v.Mode,
		"promptHash":  v.PromptHash,
		"weightsHash": v.WeightsHash,
		"notes":       v.Notes,
		"createdAt":   v.CreatedAt,
		"snapshot":    payload,
	}, "", "  ")
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	filename := fmt.Sprintf("scoring_version_%d.json", v.VersionID)
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/json; charset=utf-8", body)
}

// AdminRestoreScoringVersion POST /api/v1/admin/eval/versions/:id/restore
func (h *Handlers) AdminRestoreScoringVersion(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	v, err := services.RestoreScoringVersion(h.db, uint(id))
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, v)
}

// AdminRunEval POST /api/v1/admin/eval/runs
func (h *Handlers) AdminRunEval(c *gin.Context) {
	var body struct {
		Name      string `json:"name"`
		K         int    `json:"k"`
		VersionID uint   `json:"versionId"`
		DateFrom  string `json:"dateFrom"`
		DateTo    string `json:"dateTo"`
	}
	_ = c.ShouldBindJSON(&body)
	var from, to *time.Time
	if body.DateFrom != "" {
		if t, err := time.ParseInLocation("2006-01-02", body.DateFrom, time.Local); err == nil {
			from = &t
		}
	}
	if body.DateTo != "" {
		if t, err := time.ParseInLocation("2006-01-02", body.DateTo, time.Local); err == nil {
			to = &t
		}
	}
	run, err := services.RunTopKEval(h.db, body.Name, body.K, from, to, body.VersionID)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Created(c, run)
}

// AdminListEvalRuns GET /api/v1/admin/eval/runs
func (h *Handlers) AdminListEvalRuns(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	rows, err := services.ListEvalRuns(h.db, limit)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, rows)
}

// CDCLatestEvalCompare GET /api/v1/cdc/eval/latest
func (h *Handlers) CDCLatestEvalCompare(c *gin.Context) {
	run, err := services.LatestEvalCompare(h.db)
	if err != nil {
		respgin.Success(c, gin.H{"run": nil})
		return
	}
	respgin.Success(c, gin.H{"run": run})
}

// AdminListEvalCases GET /api/v1/admin/eval/runs/:id/cases
func (h *Handlers) AdminListEvalCases(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	rows, err := services.ListEvalCases(h.db, uint(id), limit)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, rows)
}
