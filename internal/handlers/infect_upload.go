// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
)

func (h *Handlers) uploadMultipartFile(c *gin.Context, dir string) (key, filename string, err error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", "", err
	}
	if err = pkgstorage.CheckFileSize(file.Size); err != nil {
		return "", "", err
	}
	filename = sanitizePath(filepath.Base(file.Filename))
	key = fmt.Sprintf("%s/%s_%s", sanitizePath(dir), time.Now().Format("20060102_150405"), filename)
	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer func() { _ = src.Close() }()
	if err = pkgstorage.Write(key, src); err != nil {
		return "", "", err
	}
	return key, filename, nil
}

func (h *Handlers) serveStorageKey(c *gin.Context, key, downloadName string) {
	if key == "" {
		respgin.WriteError(c, response.Err(response.CodeNotFound))
		return
	}
	reader, size, err := pkgstorage.Read(key)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeNotFound))
		return
	}
	defer func() { _ = reader.Close() }()
	name := downloadName
	if name == "" {
		name = filepath.Base(key)
	}
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, name))
	c.DataFromReader(200, size, "application/octet-stream", reader, nil)
}
