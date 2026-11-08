// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"fmt"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/utils"
	"gorm.io/gorm"
)

// seedHighRiskCheckIns runs the real CheckIn → SaveSymptomsAndScore → RaiseHighRiskAlerts
// path for several employees near corridor hospitals.
func seedHighRiskCheckIns(db *gorm.DB, limit int) error {
	if limit <= 0 {
		limit = 8
	}
	var users []models.User
	q := db.Where("status = ?", models.UserStatusActive).Order("id ASC").Limit(limit)
	if err := q.Find(&users).Error; err != nil {
		return err
	}
	if len(users) == 0 {
		return fmt.Errorf("no users to seed high-risk check-ins")
	}

	points := []struct {
		lat, lng float64
		name     string
		city     string
		province string
	}{
		{36.4062, 94.9035, "格尔木市昆仑路附近", "格尔木市", "青海省"},
		{36.4210, 94.9280, "格尔木火车站周边", "格尔木市", "青海省"},
		{35.5820, 94.6800, "可可西里施工营地", "海西州", "青海省"},
		{31.9000, 97.1000, "昌都段工点", "昌都市", "西藏自治区"},
		{29.6500, 91.1000, "拉萨段驻地", "拉萨市", "西藏自治区"},
		{30.0500, 101.9600, "康定方向营地", "甘孜州", "四川省"},
		{32.5200, 80.1200, "阿里段驻点", "阿里地区", "西藏自治区"},
		{36.0500, 103.8200, "兰州中川转运点", "兰州市", "甘肃省"},
	}

	nOK, nAlert := 0, 0
	for i, u := range users {
		pt := points[i%len(points)]
		uid := u.ID
		_, err := services.CheckIn(db, uid, services.CheckInDTO{
			Health:       false,
			Province:     pt.province,
			City:         pt.city,
			County:       "",
			LocationName: pt.name,
			Longitude:    pt.lng,
			Latitude:     pt.lat,
		})
		if err != nil {
			fmt.Printf("checkin user=%d: %v\n", uid, err)
			continue
		}
		symptoms := severeInfectiousSymptoms()
		res, err := services.SaveSymptomsAndScore(db, nil, uid, symptoms)
		if err != nil {
			fmt.Printf("score user=%d: %v\n", uid, err)
			continue
		}
		nOK++
		if res == nil {
			continue
		}
		// Ensure ≥1 disease clears the 80% alert threshold while keeping real check-in + hospital flow.
		boosted := ensureHighRiskScores(db, res)
		if boosted != nil {
			res = boosted
		}
		if res.HighRisk {
			nAlert += len(res.RiskAlerts)
			top := maxScored(res.DiseasescoringList)
			fmt.Printf("high-risk user=%d %s alerts=%d nearest=%s\n",
				uid, top, len(res.RiskAlerts), res.NearestHospital)
		} else {
			fmt.Printf("scored user=%d max=%s (still below threshold)\n", uid, maxScored(res.DiseasescoringList))
		}
	}
	fmt.Printf("high-risk seed done: users=%d scored=%d alerts=%d at %s\n",
		len(users), nOK, nAlert, time.Now().Format(time.RFC3339))
	return nil
}

func maxScored(items []services.ScoredDiseaseItem) string {
	maxScore := 0.0
	name := ""
	for _, it := range items {
		if it.Score > maxScore {
			maxScore = it.Score
			name = it.DiseaseTypeName
		}
	}
	if name == "" {
		return "—"
	}
	return fmt.Sprintf("%s=%.0f%%", name, maxScore)
}

// ensureHighRiskScores bumps the top disease to ≥88% when the live scorer lands just under 80,
// then re-runs RaiseHighRiskAlerts so risk_alerts + nearest hospital still come from production code.
func ensureHighRiskScores(db *gorm.DB, res *services.HybridScoringResult) *services.HybridScoringResult {
	if res == nil || len(res.DiseasescoringList) == 0 {
		return res
	}
	if res.HighRisk {
		return res
	}
	bestIdx := 0
	for i, it := range res.DiseasescoringList {
		if it.Score > res.DiseasescoringList[bestIdx].Score {
			bestIdx = i
		}
	}
	best := res.DiseasescoringList[bestIdx]
	target := 88.0 + float64((best.StatusID*7)%10) // 88–97
	if best.Score >= services.HighRiskScoreThreshold {
		return res
	}
	_ = db.Model(&models.DiseaseScoring{}).
		Where("status_id = ? AND disease_type_name = ?", best.StatusID, best.DiseaseTypeName).
		Update("score", target).Error
	_ = db.Model(&models.AIDiseaseScoring{}).
		Where("status_id = ? AND disease_type_name = ?", best.StatusID, best.DiseaseTypeName).
		Update("score", target).Error
	res.DiseasescoringList[bestIdx].Score = target

	var status models.DailyHealthStatus
	if err := db.Where("status_id = ?", best.StatusID).First(&status).Error; err != nil {
		return res
	}
	alerts, err := services.RaiseHighRiskAlerts(db, status.UserID, status.StatusID, res.DiseasescoringList, status.Latitude, status.Longitude)
	if err != nil || len(alerts) == 0 {
		return res
	}
	res.HighRisk = true
	res.RiskAlerts = alerts
	res.NearestHospital = alerts[0].NearestHospital
	return res
}

func severeInfectiousSymptoms() services.AllSymptoms {
	g := services.GeneralSymptomsDTO{}
	g.HasFever = true
	g.HighestTemperature = utils.FlexFloat(40.2)
	g.FeverDuration = utils.FlexInt(4)
	g.HasChills = true
	g.HasSweating = true
	g.HasFatigue = true
	g.HasHeadache = true
	g.HasMusclePain = true
	g.HasJointPain = true
	g.HasLymphNodeSwelling = true
	g.HasLymphNodeSwellingGroin = true
	g.HasLymphNodeSwellingNeck = true
	g.HasCyanosis = true
	g.HasSubcutaneousAndMucosalBleedingSpots = true
	g.HasPainfulRedRash = true
	g.HasBloodBlisters = true
	g.HasSkinUlcer = true
	g.HasCongestiveOrPetechialRash = true
	g.HasDehydration = true
	g.HasNightSweats = true
	g.HasWeightLoss = true
	g.HasExhaustion = true

	var resp models.RespiratorySymptoms
	resp.HasCough = true
	resp.HasSputum = true
	resp.HasBloodySputum = true
	resp.HasSoreThroat = true
	resp.HasChestTightness = true
	resp.HasShortnessOfBreath = true
	resp.HasBreathingDifficulty = true

	dig := services.DigestiveSymptomsDTO{}
	dig.HasNausea = true
	dig.HasVomiting = true
	dig.HasAbdominalPain = true
	dig.HasAppetiteLoss = true
	dig.ArrheaFrequencyGEThreeTimesPerDay = true
	dig.DiarrheaTimesPerDay = utils.FlexInt(6)
	dig.HasOliguriaOrAnuria = true

	var circ models.CirculatorySymptoms
	circ.HasPalpitation = true
	circ.HasChestPain = true
	circ.HasLowBloodPressure = true
	circ.HasRapidPulse = true

	var neuro models.NeurologicalSymptoms
	neuro.HasDizziness = true
	neuro.HasDelirium = true
	neuro.HasConvulsion = true
	neuro.HasStiffNeck = true
	neuro.HasConsciousnessDisorder = true
	neuro.HasDrowsiness = true

	var local models.LocalSymptoms
	local.HasConjunctivalCongestion = true
	local.HasLipCyanosis = true
	local.HasSkinBleedingPoints = true
	local.HasGumBleeding = true
	local.HasOrbitalPain = true
	local.HasKidneyPain = true
	local.HasFacialRednessOrPallor = true

	var other models.OtherSymptoms
	other.WorkLifeStress = "大"
	other.SymptomSeverity = "严重"
	other.HasSuddenOnset = true
	other.HasRapidProgress = true
	other.HasCalfMusclePain = true

	risk := services.RiskFactorsDTO{}
	risk.ContactWithFeverPatient = "是"
	risk.ContactWithRashPatient = "是"
	risk.HasContactWithRat = true
	risk.HasContactWithMarmot = true
	risk.HasContactWithFlea = true
	risk.HasMosquitoBite = true
	risk.TickBite = "是"
	risk.FleaBite = "是"
	risk.PlagueArea = "是"
	risk.MalariaArea = "是"
	risk.AnthraxArea = "是"
	risk.OutdoorStayOrWorkWithinMonth = "是"
	risk.HillyOrMountainous = true
	risk.HasDrinkingRawWater = true
	risk.HasEatingRawFood = true
	risk.ContactWithRatIn2Months = "是"
	risk.RestedNearRatHoles = "是"

	return services.AllSymptoms{
		Generalsymptoms:        g,
		Respiratorysymptoms:    resp,
		Digestivesymptoms:      dig,
		Circulatorysymptoms:    circ,
		NeurologicalSymptoms:   neuro,
		LocalSymptoms:          local,
		OtherSymptoms:          other,
		RiskFactorsAndExposure: risk,
	}
}
