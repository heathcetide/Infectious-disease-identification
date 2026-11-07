// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/internal/types"

	"github.com/heathcetide/infectious-disease/pkg/common"
	"github.com/heathcetide/infectious-disease/pkg/common/jwtutil"
	"github.com/heathcetide/infectious-disease/pkg/common/password"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	"github.com/heathcetide/infectious-disease/pkg/common/validate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login authenticates a user and issues JWT access/refresh tokens.
// Supports username, email, or phoneNumber + password (infect mobile login).
//
// Password is verified using ling-base/common/password (Argon2id/bcrypt).
func (h *Handlers) Login(c *gin.Context) {
	if h.jwtAuth == nil {
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

	if h.db == nil {
		respgin.WriteError(c, response.Err(response.CodeServiceUnavail))
		return
	}

	// Validate that at least one identifier is provided
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
			respgin.WriteError(c, response.New(response.CodeUnauthorized, "invalid credentials"))
			return
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}

	// Check account status
	if !user.IsActive() {
		respgin.WriteError(c, response.New(response.CodeForbidden, "account is disabled"))
		return
	}

	// Verify password using ling-base/common/password
	if !password.Verify(req.Password, user.Password) {
		respgin.WriteError(c, response.New(response.CodeUnauthorized, "invalid credentials"))
		return
	}

	// Issue tokens — subject is numeric user ID so infect handlers can resolve reliably
	pair, err := h.jwtAuth.Login(strconv.FormatUint(uint64(user.ID), 10), jwtutil.WithRoles(user.Role))
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return
	}

	// Update last login time (unix timestamp)
	_ = models.User{}.UpdateLastLogin(ctxDB, user.ID, time.Now().Unix())

	// Emit login event for post-processing (audit log, notifications, etc.)
	common.Sig().Emit(constants.EventUserLogin, h, user.ID, user.Username)

	infoComplete := false
	if h.svc != nil && h.svc.Profile != nil {
		infoComplete = h.svc.Profile.InformationComplete(user.ID)
	}

	respgin.Success(c, buildLoginResponse(user, pair, infoComplete))
}

func buildLoginResponse(user *models.User, pair *jwtutil.TokenPair, infoComplete bool) types.LoginResponse {
	return types.LoginResponse{
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
		ExpiresIn:    pair.ExpiresIn,
		TokenType:    pair.TokenType,
		User: types.InfectUserBrief{
			ID:    user.ID,
			Name:  user.Name,
			Phone: userPhone(user),
			Role:  user.Role,
		},
		InformationComplete: infoComplete,
		Token:               pair.AccessToken,
		UserType:            types.RoleDisplayName(user.Role),
		UserName:            user.Name,
	}
}

// Refresh exchanges a refresh token for a new access/refresh token pair.
func (h *Handlers) Refresh(c *gin.Context) {
	if h.jwtAuth == nil {
		respgin.WriteError(c, response.Err(response.CodeServiceUnavail))
		return
	}

	var req types.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	if err := validate.Validate(&req); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}

	pair, err := h.jwtAuth.Refresh(req.RefreshToken)
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeUnauthorized))
		return
	}
	respgin.Success(c, types.TokenResponse{
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
		ExpiresIn:    pair.ExpiresIn,
		TokenType:    pair.TokenType,
	})
}

// ClaimsFromContext extracts JWT claims from the gin.Context
// (must be used after jwtingin.Middleware).
func ClaimsFromContext(c *gin.Context) *jwtutil.Claims {
	if v, ok := c.Get("jwt_claims"); ok {
		if claims, ok := v.(*jwtutil.Claims); ok {
			return claims
		}
	}
	return nil
}
