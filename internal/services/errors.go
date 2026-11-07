// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import "errors"

var (
	ErrNotImplemented    = errors.New("not implemented")
	ErrNoDailyCheckIn    = errors.New("no daily check-in for today")
	ErrAlreadySubmitted  = errors.New("already submitted for today")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrPasswordMismatch  = errors.New("new password and confirm password do not match")
	ErrUserNotFound      = errors.New("user not found")
	ErrProfileNotFound   = errors.New("employee profile not found")
	ErrInvalidDate       = errors.New("invalid date")
)
