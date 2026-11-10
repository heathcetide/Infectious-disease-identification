// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
)

// AdminLLMChat POST /api/v1/admin/llm/chat
func (h *Handlers) AdminLLMChat(c *gin.Context) {
	var req services.AdminChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	out, err := services.AdminLLMChat(c.Request.Context(), h.svc.LLM, h.cfg.LLM.Model, req)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, out)
}

// AdminLLMInfo GET /api/v1/admin/llm/info
func (h *Handlers) AdminLLMInfo(c *gin.Context) {
	enabled := h.svc.LLM != nil && h.svc.LLM.Enabled()
	respgin.Success(c, gin.H{
		"enabled":  enabled,
		"model":    h.cfg.LLM.Model,
		"protocol": h.cfg.LLM.Protocol,
		"baseURL":  h.cfg.LLM.BaseURL,
	})
}
