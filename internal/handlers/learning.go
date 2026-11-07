// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/internal/types"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
	"github.com/gin-gonic/gin"
)

// LearningMaterialsPage POST /api/v1/learning/materials/page
func (h *Handlers) LearningMaterialsPage(c *gin.Context) {
	var req services.LearningMaterialsPageDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		var page types.PageBodyRequest
		_ = c.ShouldBindQuery(&page.PageRequest)
		p, s := page.Normalize()
		req.PageNo = p
		req.PageSize = s
	}
	if req.PageNo < 1 {
		req.PageNo = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 20
	}
	pageData, err := h.svc.Learning.PageSelect(req)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, pageData)
}

// InfectLearningFile GET /api/v1/infect-files/learning/:filename
func (h *Handlers) InfectLearningFile(c *gin.Context) {
	var req types.FilenamePathRequest
	if err := c.ShouldBindUri(&req); err != nil {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	name := sanitizePath(filepath.Base(req.Filename))
	candidates := []string{
		sanitizePath(filepath.Join("infect/materials", name)),
		sanitizePath(filepath.Join("learning", name)),
		name,
	}
	var (
		reader io.ReadCloser
		size   int64
		err    error
	)
	for _, key := range candidates {
		reader, size, err = pkgstorage.Read(key)
		if err == nil {
			defer reader.Close()
			c.Header("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, name))
			c.DataFromReader(200, size, "application/octet-stream", reader, nil)
			return
		}
	}
	respgin.WriteError(c, response.Err(response.CodeNotFound))
}
