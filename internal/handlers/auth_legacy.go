// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/constants"
	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/internal/types"
	"github.com/heathcetide/infectious-disease/pkg/common"
	"github.com/heathcetide/infectious-disease/pkg/common/jwtutil"
	"github.com/heathcetide/infectious-disease/pkg/common/password"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/heathcetide/infectious-disease/pkg/common/validate"
)

// UserLogin POST /user/login — legacy admin login envelope (code=1).
func (h *Handlers) UserLogin(c *gin.Context) {
	if h.jwtAuth == nil || h.db == nil {
		respgin.WriteError(c, response.Err(response.CodeServiceUnavail))
		return
	}
	var req types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := validate.Validate(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if req.Username == "" && req.Email == "" && req.PhoneNumber == "" {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "username, email, or phoneNumber is required"))
		return
	}

	ctxDB := h.db.WithContext(c.Request.Context())
	var user *models.User
	var err error
	switch {
	case req.PhoneNumber != "":
		user, err = models.User{}.FindByPhone(ctxDB, req.PhoneNumber)
	default:
		user, err = models.User{}.FindByUsernameOrEmail(ctxDB, req.Username, req.Email)
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.NewI18n(response.CodeCredentialInvalid, response.KeyAuthInvalidCredentials))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	if !user.IsActive() {
		respgin.WriteError(c, response.New(response.CodeForbidden, "account is disabled"))
		return
	}
	if !password.Verify(req.Password, user.Password) {
		respgin.WriteError(c, response.NewI18n(response.CodeCredentialInvalid, response.KeyAuthInvalidCredentials))
		return
	}
	pair, err := h.jwtAuth.Login(strconv.FormatUint(uint64(user.ID), 10), jwtutil.WithRoles(user.Role))
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}
	_ = models.User{}.UpdateLastLogin(ctxDB, user.ID, time.Now().Unix())
	common.Sig().Emit(constants.EventUserLogin, h, user.ID, user.Username)

	infoComplete := false
	if h.svc != nil && h.svc.Profile != nil {
		infoComplete = h.svc.Profile.InformationComplete(user.ID)
	}
	mustChange := isInitialPassword(ctxDB, user, req.Password)
	legacyAdminSuccess(c, buildLoginResponse(user, pair, infoComplete, mustChange))
}
