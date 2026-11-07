// Copyright (c) 2026 heathcetide. All rights reserved.

package utils

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// FlexInt unmarshals JSON numbers or numeric strings (e.g. uni-app text inputs).
type FlexInt int

func (v *FlexInt) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) || bytes.Equal(b, []byte(`""`)) {
		*v = 0
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		s = strings.TrimSpace(s)
		if s == "" {
			*v = 0
			return nil
		}
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*v = FlexInt(n)
		return nil
	}
	var n json.Number
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	if err := dec.Decode(&n); err != nil {
		return err
	}
	f, err := n.Float64()
	if err != nil {
		return err
	}
	*v = FlexInt(f)
	return nil
}

func (v FlexInt) Int() int { return int(v) }

// FlexFloat unmarshals JSON numbers or numeric strings.
type FlexFloat float64

func (v *FlexFloat) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) || bytes.Equal(b, []byte(`""`)) {
		*v = 0
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		s = strings.TrimSpace(s)
		if s == "" {
			*v = 0
			return nil
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*v = FlexFloat(f)
		return nil
	}
	var n json.Number
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	if err := dec.Decode(&n); err != nil {
		return err
	}
	f, err := n.Float64()
	if err != nil {
		return err
	}
	*v = FlexFloat(f)
	return nil
}

func (v FlexFloat) Float64() float64 { return float64(v) }
