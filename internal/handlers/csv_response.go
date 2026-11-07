// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func writeCSV(c *gin.Context, filename string, data []byte) {
	c.Header("Content-Type", "text/csv; charset=utf-8")
	c.Header("Content-Disposition", `attachment; filename="`+filename+`"`)
	c.Data(http.StatusOK, "text/csv; charset=utf-8", data)
}
