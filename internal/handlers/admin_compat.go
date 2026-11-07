// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const legacyAdminCodeSuccess = 1

// legacyAdminSuccess writes the Java-style admin envelope {code:1, msg, data}.
func legacyAdminSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": legacyAdminCodeSuccess,
		"msg":  "success",
		"data": data,
	})
}

// legacyAdminSuccessBool returns a bare boolean in data (autoweight getter).
func legacyAdminSuccessBool(c *gin.Context, v bool) {
	legacyAdminSuccess(c, v)
}

type adminPageParams struct {
	PageNo     int `form:"pageNo" json:"pageNo"`
	PageNumber int `form:"pageNumber" json:"pageNumber"`
	PageNum    int `form:"pageNum" json:"pageNum"`
	Page       int `form:"page" json:"page"`
	Size       int `form:"size" json:"size"`
	SurveyID   uint `form:"surveyID" json:"surveyID"`
}

func (p adminPageParams) normalized() (page, size int) {
	page = p.PageNo
	if page < 1 {
		page = p.PageNum
	}
	if page < 1 {
		page = p.Page
	}
	if page < 1 {
		page = 1
	}
	size = p.PageNumber
	if size < 1 {
		size = p.Size
	}
	if size < 1 {
		size = 15
	}
	if size > 100 {
		size = 100
	}
	return page, size
}

func pathUintOrQuery(c *gin.Context, key string) (uint, bool) {
	if v := c.Param(key); v != "" {
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil || id == 0 {
			return 0, false
		}
		return uint(id), true
	}
	q := c.Query(key)
	if q == "" {
		return 0, false
	}
	id, err := strconv.ParseUint(q, 10, 64)
	if err != nil || id == 0 {
		return 0, false
	}
	return uint(id), true
}
