// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handlers) pathUint(c *gin.Context, key string) (uint, bool) {
	id, ok := pathUintParam(c, key)
	return id, ok
}

func pathUintParam(c *gin.Context, key string) (uint, bool) {
	id, err := parseUintParam(c.Param(key))
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "invalid id"))
		return 0, false
	}
	return id, true
}

// AdminHello GET /api/v1/admin/hello
func (h *Handlers) AdminHello(c *gin.Context) {
	respgin.Success(c, gin.H{"message": "hello"})
}

// AdminGetUser GET /api/v1/admin/users/:id — profile + user fields for drawers.
func (h *Handlers) AdminGetUser(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.Admin.GetUserDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminCreateStaffUser POST /api/v1/admin/users
// Accepts the admin drawer payload (phone/name/userType + profile fields).
func (h *Handlers) AdminCreateStaffUser(c *gin.Context) {
	var dto services.AdminCreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "请求参数无效: "+err.Error()))
		return
	}
	user, err := h.svc.Admin.CreateStaffUser(dto)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "不能为空") || strings.Contains(msg, "无效") || strings.Contains(msg, "已存在") {
			respgin.WriteError(c, response.New(response.CodeBadRequest, msg))
			return
		}
		respgin.WriteError(c, response.New(response.CodeInternal, msg))
		return
	}
	idNumber := strings.TrimSpace(dto.IDNumber)
	plainPwd := services.InitialPasswordForDisplay(idNumber, strings.TrimSpace(dto.PhoneNumber))
	respgin.Success(c, gin.H{
		"userId":           strconv.FormatUint(uint64(user.ID), 10),
		"name":             user.Name,
		"phoneNumber":      user.PhoneNumber,
		"role":             user.Role,
		"initialPassword":  plainPwd,
		"message":          "创建成功，初始密码为身份证号后六位：" + plainPwd,
	})
}

// AdminDiseasesPage POST /api/v1/admin/diseases/page
func (h *Handlers) AdminDiseasesPage(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.Admin.PageDiseases(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminHospitalPage POST /api/v1/admin/hospitals/page
func (h *Handlers) AdminHospitalPage(c *gin.Context) {
	q := bindPageQuery(c)
	data, err := h.svc.Admin.PageHospitals(q)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminGetHospital GET /api/v1/admin/hospitals/:id
func (h *Handlers) AdminGetHospital(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.Admin.GetHospital(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminCreateHospital POST /api/v1/admin/hospitals
func (h *Handlers) AdminCreateHospital(c *gin.Context) {
	var body models.Hospital
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	data, err := h.svc.Admin.CreateHospital(body)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminUpdateHospital PUT /api/v1/admin/hospitals/:id
func (h *Handlers) AdminUpdateHospital(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	var body models.Hospital
	if err := c.ShouldBindJSON(&body); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	data, err := h.svc.Admin.UpdateHospital(id, body)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminDeleteHospital DELETE /api/v1/admin/hospitals/:id
func (h *Handlers) AdminDeleteHospital(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	if err := h.svc.Admin.DeleteHospital(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"deleted": true})
}

// AdminSetUserActive PUT /api/v1/admin/users/:id/active
// Body optional: {active|status}. Empty body toggles current status (admin UI).
func (h *Handlers) AdminSetUserActive(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	var body struct {
		Active *bool `json:"active"`
		Status *int  `json:"status"`
	}
	_ = c.ShouldBindJSON(&body)

	active := true
	switch {
	case body.Active != nil:
		active = *body.Active
	case body.Status != nil:
		active = *body.Status == models.UserStatusActive
	default:
		u, err := models.User{}.FindByID(h.db.WithContext(c.Request.Context()), id)
		if err != nil {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		active = !u.IsActive()
	}
	if err := h.svc.Admin.SetUserActive(id, active); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"active": active})
}

// AdminResetUserPassword PUT /api/v1/admin/users/:id/password/reset
func (h *Handlers) AdminResetUserPassword(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	plain, err := h.svc.Admin.ResetUserPassword(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{
		"message":          "password reset",
		"defaultPassword":  plain,
		"initialPassword":  plain,
		"hint":             "已重置为身份证号后六位（无身份证则手机号后六位）",
	})
}

// AdminWeightsBatch PUT /api/v1/admin/weights/batch
func (h *Handlers) AdminWeightsBatch(c *gin.Context) {
	var items []services.SymptomWeightUpdateDTO
	if err := c.ShouldBindJSON(&items); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := h.svc.Admin.UpdateWeightsBatch(items); err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, gin.H{"updated": len(items)})
}

// AdminGetFeedback GET /api/v1/admin/feedback/:id
func (h *Handlers) AdminGetFeedback(c *gin.Context) {
	id, ok := h.pathUint(c, "id")
	if !ok {
		return
	}
	data, err := h.svc.Admin.GetFeedback(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeNotFound))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	respgin.Success(c, data)
}

// AdminExportUsers GET /api/v1/admin/users/export
func (h *Handlers) AdminExportUsers(c *gin.Context) {
	data, err := h.svc.Admin.ExportUsersCSV()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "users.csv", data)
}

// AdminImportUsers POST /api/v1/admin/users/import
func (h *Handlers) AdminImportUsers(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "file required"))
		return
	}
	f, err := file.Open()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	defer f.Close()
	n, err := h.svc.Admin.ImportUsersCSV(f)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"imported": n})
}

// AdminExportHospitals GET /api/v1/admin/hospitals/export
func (h *Handlers) AdminExportHospitals(c *gin.Context) {
	data, err := h.svc.Admin.ExportHospitalsCSV()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "hospitals.csv", data)
}

// AdminImportHospitals POST /api/v1/admin/hospitals/import
func (h *Handlers) AdminImportHospitals(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "file required"))
		return
	}
	f, err := file.Open()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	defer f.Close()
	n, err := h.svc.Admin.ImportHospitalsCSV(f)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"imported": n})
}

// AdminHospitalsImportTemplate GET /api/v1/admin/hospitals/import/template
func (h *Handlers) AdminHospitalsImportTemplate(c *gin.Context) {
	writeCSV(c, "hospitals_template.csv", services.HospitalImportTemplateCSV())
}

// AdminExportDiseases GET /api/v1/admin/diseases/export
func (h *Handlers) AdminExportDiseases(c *gin.Context) {
	data, err := h.svc.Admin.ExportDiseasesCSV()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "diseases.csv", data)
}

// AdminExportWeights GET /api/v1/admin/weights/export
func (h *Handlers) AdminExportWeights(c *gin.Context) {
	data, err := h.svc.Admin.ExportWeightsCSV(c.Query("diseaseTypeName"))
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "weights.csv", data)
}

// AdminImportWeights POST /api/v1/admin/weights/import
func (h *Handlers) AdminImportWeights(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "file required"))
		return
	}
	f, err := file.Open()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	defer f.Close()
	n, err := h.svc.Admin.ImportWeightsCSV(f)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, gin.H{"imported": n})
}

// AdminWeightsExportTemplate GET /api/v1/admin/weights/export-template
func (h *Handlers) AdminWeightsExportTemplate(c *gin.Context) {
	writeCSV(c, "weights_template.csv", services.SymptomWeightsTemplateCSV())
}

// AdminExportFeedback GET /api/v1/admin/feedback/export
func (h *Handlers) AdminExportFeedback(c *gin.Context) {
	data, err := h.svc.Admin.ExportFeedbackCSV()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	writeCSV(c, "feedback.csv", data)
}
