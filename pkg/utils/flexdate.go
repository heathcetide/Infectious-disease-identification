// Copyright (c) 2026 heathcetide. All rights reserved.

package utils

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"
)

// FlexDate unmarshals JSON date strings (or null/empty) into *time.Time.
// Accepts yyyy-MM-dd and common RFC3339 variants from uni-app pickers.
type FlexDate struct {
	t *time.Time
}

func (d *FlexDate) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) || bytes.Equal(b, []byte(`""`)) {
		d.t = nil
		return nil
	}
	var s string
	if b[0] == '"' {
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
	} else {
		s = string(b)
	}
	s = strings.TrimSpace(s)
	if s == "" {
		d.t = nil
		return nil
	}
	layouts := []string{
		"2006-01-02",
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006/01/02",
	}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, s, time.Local); err == nil {
			d.t = &t
			return nil
		}
	}
	return &time.ParseError{Layout: "2006-01-02", Value: s, LayoutElem: "2006", ValueElem: s}
}

// Ptr returns the parsed time, or nil when empty/unset.
func (d FlexDate) Ptr() *time.Time { return d.t }
