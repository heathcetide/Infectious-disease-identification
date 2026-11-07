// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import "time"

// monthDateRange returns [start, end) for a yyyy-MM month in local time.
func monthDateRange(yearMonth string) (start, end time.Time, err error) {
	if len(yearMonth) != 7 {
		return time.Time{}, time.Time{}, ErrInvalidDate
	}
	t, err := time.ParseInLocation("2006-01", yearMonth, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, ErrInvalidDate
	}
	start = dateOnly(t)
	end = start.AddDate(0, 1, 0)
	return start, end, nil
}

// todayDateRange returns [start, end) for the current calendar day.
func todayDateRange() (start, end time.Time) {
	start = dateOnly(time.Now())
	end = start.AddDate(0, 0, 1)
	return start, end
}
