// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Dense season: 2026-11-08 … 2027-07-31 — every day high check-in volume for the screen.
func seedDenseSeason(db *gorm.DB) error {
	rng := rand.New(rand.NewSource(20261108))
	start := time.Date(2026, 11, 8, 0, 0, 0, 0, time.Local)
	end := time.Date(2027, 7, 31, 0, 0, 0, 0, time.Local)

	silent := db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})

	var workerIDs []uint
	if err := silent.Model(&models.User{}).
		Where("role = ?", models.UserRoleRailway).
		Pluck("id", &workerIDs).Error; err != nil {
		return err
	}
	if len(workerIDs) == 0 {
		return fmt.Errorf("dense season: no railway workers")
	}

	// Precompute home station by worker index.
	homes := make([]railStation, len(workerIDs))
	for i := range workerIDs {
		homes[i] = pickHomeStation(rng, i+1)
	}

	batch := make([]models.DailyHealthStatus, 0, 4000)
	flush := func() error {
		if len(batch) == 0 {
			return nil
		}
		const chunk = 500
		for i := 0; i < len(batch); i += chunk {
			j := i + chunk
			if j > len(batch) {
				j = len(batch)
			}
			if err := silent.Create(batch[i:j]).Error; err != nil {
				return err
			}
		}
		batch = batch[:0]
		return nil
	}

	totalAdded := 0
	for day := start; !day.After(end); day = day.AddDate(0, 0, 1) {
		dayStart := day
		dayEnd := day.AddDate(0, 0, 1)
		var existingIDs []uint
		_ = silent.Model(&models.DailyHealthStatus{}).
			Where("check_in_date >= ? AND check_in_date < ?", dayStart, dayEnd).
			Pluck("user_id", &existingIDs)
		have := make(map[uint]struct{}, len(existingIDs))
		for _, id := range existingIDs {
			have[id] = struct{}{}
		}

		// 拉满：工作日 ~96%，周末 ~88%，下限 85%
		rate := 0.96
		if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
			rate = 0.88
		}
		target := int(float64(len(workerIDs)) * rate)
		minFloor := len(workerIDs) * 85 / 100
		if target < minFloor {
			target = minFloor
		}
		need := target - len(have)
		if need <= 0 {
			fmt.Printf("dense %s: already %d (ok)\n", day.Format("2006-01-02"), len(have))
			continue
		}

		order := rng.Perm(len(workerIDs))
		addedToday := 0
		for _, wi := range order {
			if addedToday >= need {
				break
			}
			uid := workerIDs[wi]
			if _, ok := have[uid]; ok {
				continue
			}
			st := pickDenseCheckInStation(rng, homes[wi])

			h := true
			disease := ""
			diag := ""
			unhealthyP := 0.32
			if st.Altitude > 4500 {
				unhealthyP = 0.42
			}
			// 冬春季发病率更高，暑期略降
			switch day.Month() {
			case time.November, time.December, time.January, time.February:
				unhealthyP += 0.12
			case time.March, time.April:
				unhealthyP += 0.06
			case time.May, time.June, time.July:
				unhealthyP -= 0.04
			}
			if unhealthyP < 0.18 {
				unhealthyP = 0.18
			}
			if unhealthyP > 0.62 {
				unhealthyP = 0.62
			}
			if rng.Float64() < unhealthyP {
				h = false
				disease = symptomDiseases[rng.Intn(len(symptomDiseases))]
				diag = disease
			}
			hh := h
			checkTime := time.Date(day.Year(), day.Month(), day.Day(), 6+rng.Intn(6), rng.Intn(60), rng.Intn(60), 0, time.Local)
			batch = append(batch, models.DailyHealthStatus{
				UserID:                   uid,
				CheckInDate:              dayStart,
				CheckInTime:              checkTime,
				IsHealth:                 &hh,
				DiseaseTypeName:          disease,
				DiagnosisDiseaseTypeName: diag,
				Longitude:                round6(st.Longitude + (rng.Float64()-0.5)*0.06),
				Latitude:                 round6(st.Latitude + (rng.Float64()-0.5)*0.05),
				Province:                 st.Province,
				City:                     st.City,
				County:                   st.County,
				LocationName:             st.Name,
			})
			have[uid] = struct{}{}
			addedToday++
			totalAdded++
			if len(batch) >= 3500 {
				if err := flush(); err != nil {
					return err
				}
			}
		}
		if err := flush(); err != nil {
			return err
		}
		fmt.Printf("dense %s: +%d (~%d)\n", day.Format("2006-01-02"), addedToday, len(have))
	}

	var checkN int64
	_ = db.Model(&models.DailyHealthStatus{}).Count(&checkN)
	fmt.Printf("dense season done: added=%d total_checkins=%d workers=%d\n", totalAdded, checkN, len(workerIDs))
	return nil
}
