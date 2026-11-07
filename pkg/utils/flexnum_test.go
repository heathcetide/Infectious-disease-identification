// Copyright (c) 2026 heathcetide. All rights reserved.

package utils

import (
	"encoding/json"
	"testing"
)

func TestFlexIntAndFloatAcceptStringNumbers(t *testing.T) {
	type sample struct {
		People   FlexInt   `json:"people"`
		Pressure FlexFloat `json:"pressure"`
		MaxTemp  FlexFloat `json:"maxTemp"`
		MinTemp  FlexFloat `json:"minTemp"`
	}
	raw := `{"people":"3","pressure":"101.5","maxTemp":36.5,"minTemp":""}`
	var s sample
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if s.People.Int() != 3 {
		t.Fatalf("people=%d", s.People.Int())
	}
	if s.Pressure.Float64() != 101.5 {
		t.Fatalf("pressure=%v", s.Pressure.Float64())
	}
	if s.MaxTemp.Float64() != 36.5 {
		t.Fatalf("maxTemp=%v", s.MaxTemp.Float64())
	}
	if s.MinTemp.Float64() != 0 {
		t.Fatalf("empty minTemp should be 0, got %v", s.MinTemp.Float64())
	}
}

func TestFlexDateAcceptsEmptyAndYMD(t *testing.T) {
	type sample struct {
		ContactDate FlexDate `json:"contactDate"`
	}
	var empty sample
	if err := json.Unmarshal([]byte(`{"contactDate":""}`), &empty); err != nil {
		t.Fatalf("empty: %v", err)
	}
	if empty.ContactDate.Ptr() != nil {
		t.Fatalf("empty should be nil")
	}
	var dated sample
	if err := json.Unmarshal([]byte(`{"contactDate":"2024-06-15"}`), &dated); err != nil {
		t.Fatalf("ymd: %v", err)
	}
	if dated.ContactDate.Ptr() == nil || dated.ContactDate.Ptr().Format("2006-01-02") != "2024-06-15" {
		t.Fatalf("ymd parse failed: %#v", dated.ContactDate.Ptr())
	}
}
