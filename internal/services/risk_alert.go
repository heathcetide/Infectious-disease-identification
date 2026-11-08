// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

const HighRiskScoreThreshold = 80.0

// HospitalDistance is a hospital with distance in km from a point.
type HospitalDistance struct {
	models.Hospital
	DistanceKm float64 `json:"distanceKm"`
}

func haversineKm(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371.0
	toRad := func(d float64) float64 { return d * math.Pi / 180 }
	dLat := toRad(lat2 - lat1)
	dLng := toRad(lng2 - lng1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(toRad(lat1))*math.Cos(toRad(lat2))*math.Sin(dLng/2)*math.Sin(dLng/2)
	return 2 * R * math.Asin(math.Min(1, math.Sqrt(a)))
}

// NearbyHospitals returns hospitals sorted by distance; missing coords sort last.
func NearbyHospitals(db *gorm.DB, lat, lng float64, limit int) ([]HospitalDistance, error) {
	if limit <= 0 || limit > 50 {
		limit = 10
	}
	var rows []models.Hospital
	if err := db.Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]HospitalDistance, 0, len(rows))
	for _, h := range rows {
		d := HospitalDistance{Hospital: h}
		if h.Latitude != 0 || h.Longitude != 0 {
			if lat != 0 || lng != 0 {
				d.DistanceKm = haversineKm(lat, lng, h.Latitude, h.Longitude)
			}
		} else {
			d.DistanceKm = 1e9
		}
		out = append(out, d)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].DistanceKm < out[j].DistanceKm })
	if len(out) > limit {
		out = out[:limit]
	}
	return out, nil
}

// RaiseHighRiskAlerts persists alerts for scores >= threshold and returns the highest one.
func RaiseHighRiskAlerts(db *gorm.DB, userID, statusID uint, items []ScoredDiseaseItem, userLat, userLng float64) ([]models.RiskAlert, error) {
	var raised []models.RiskAlert
	nearestName := ""
	var nearestID *uint
	hospitals, _ := NearbyHospitals(db, userLat, userLng, 1)
	if len(hospitals) > 0 {
		h := hospitals[0]
		nearestName = h.HospitalName
		id := h.HospitalID
		nearestID = &id
	}
	for _, it := range items {
		if it.Score < HighRiskScoreThreshold {
			continue
		}
		msg := fmt.Sprintf("打卡风险评分 %.0f%% 超过预警阈值 %.0f%%，建议尽快就医", it.Score, HighRiskScoreThreshold)
		alert := models.RiskAlert{
			UserID: userID, StatusID: statusID,
			DiseaseTypeName: it.DiseaseTypeName, Score: it.Score,
			Threshold: HighRiskScoreThreshold,
			NearestHospital: nearestName, HospitalID: nearestID,
			Latitude: userLat, Longitude: userLng,
			Message: msg, CreatedAt: time.Now(),
		}
		if err := db.Create(&alert).Error; err != nil {
			return raised, err
		}
		raised = append(raised, alert)
	}
	return raised, nil
}

// RiskAlertView enriches RiskAlert with user display fields for admin UI.
type RiskAlertView struct {
	models.RiskAlert
	Username    string `json:"username"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

// ListRiskAlerts for admin/CDC review.
func ListRiskAlerts(db *gorm.DB, onlyOpen bool, limit int) ([]RiskAlertView, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	q := db.Model(&models.RiskAlert{})
	if onlyOpen {
		q = q.Where("acknowledged = ?", false)
	}
	var rows []models.RiskAlert
	if err := q.Order("alert_id DESC").Limit(limit).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]RiskAlertView, 0, len(rows))
	if len(rows) == 0 {
		return out, nil
	}
	ids := make([]uint, 0, len(rows))
	seen := map[uint]struct{}{}
	for _, r := range rows {
		if _, ok := seen[r.UserID]; ok {
			continue
		}
		seen[r.UserID] = struct{}{}
		ids = append(ids, r.UserID)
	}
	type userLite struct {
		ID          uint
		Username    string
		Name        string
		PhoneNumber string
		Phone       string
	}
	var users []userLite
	_ = db.Model(&models.User{}).Select("id", "username", "name", "phone_number", "phone").
		Where("id IN ?", ids).Find(&users).Error
	byID := make(map[uint]userLite, len(users))
	for _, u := range users {
		byID[u.ID] = u
	}
	for _, r := range rows {
		view := RiskAlertView{RiskAlert: r}
		if u, ok := byID[r.UserID]; ok {
			view.Username = u.Username
			view.Name = u.Name
			view.PhoneNumber = u.PhoneNumber
			if view.PhoneNumber == "" {
				view.PhoneNumber = u.Phone
			}
		}
		out = append(out, view)
	}
	return out, nil
}

func AcknowledgeRiskAlert(db *gorm.DB, id uint) error {
	return db.Model(&models.RiskAlert{}).Where("alert_id = ?", id).
		Update("acknowledged", true).Error
}
