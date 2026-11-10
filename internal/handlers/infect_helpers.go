// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/models"
	jwtingin "github.com/heathcetide/infectious-disease/pkg/common/jwtutil/gin"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
)

// currentUserID resolves the authenticated user ID from JWT subject (numeric id or username).
func (h *Handlers) currentUserID(c *gin.Context) (uint, bool) {
	subject := jwtingin.UserIDFromGin(c)
	if subject == "" {
		if claims := ClaimsFromContext(c); claims != nil {
			subject = claims.Subject
		}
	}
	if subject == "" {
		respgin.WriteError(c, response.Err(response.CodeUnauthorized))
		return 0, false
	}
	if id, err := strconv.ParseUint(subject, 10, 64); err == nil {
		return uint(id), true
	}
	if h.db == nil {
		respgin.WriteError(c, response.Err(response.CodeUnauthorized))
		return 0, false
	}
	user, err := models.User{}.FindByUsername(h.db.WithContext(c.Request.Context()), subject)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respgin.WriteError(c, response.Err(response.CodeUnauthorized))
			return 0, false
		}
		respgin.WriteError(c, response.Err(response.CodeInternal))
		return 0, false
	}
	return user.ID, true
}

func parseUintParam(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil || id == 0 {
		return 0, strconv.ErrSyntax
	}
	return uint(id), nil
}

func userPhone(u *models.User) string {
	if u.PhoneNumber != "" {
		return u.PhoneNumber
	}
	return u.Phone
}
