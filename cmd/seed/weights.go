// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"fmt"
	"math"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/models"
)

// weightCatalog lists every slider field expected by admin weight UI.
// tableName must be PascalCase to match weightset.vue filters.
var weightCatalog = []struct {
	Table string
	Field string
}{
	// GeneralSymptoms
	{"GeneralSymptoms", "HasFever"},
	{"GeneralSymptoms", "HasChills"},
	{"GeneralSymptoms", "HasSweating"},
	{"GeneralSymptoms", "HasFatigue"},
	{"GeneralSymptoms", "HasHeadache"},
	{"GeneralSymptoms", "HasMusclePain"},
	{"GeneralSymptoms", "HasJointPain"},
	{"GeneralSymptoms", "HasLymphNodeSwelling"},
	{"GeneralSymptoms", "HasCyanosis"},
	{"GeneralSymptoms", "HasSubcutaneousAndMucosalBleedingSpots"},
	{"GeneralSymptoms", "HasPainfulRedRash"},
	{"GeneralSymptoms", "HasBloodBlisters"},
	{"GeneralSymptoms", "HasSkinUlcer"},
	{"GeneralSymptoms", "HasCongestiveOrPetechialRash"},
	{"GeneralSymptoms", "HasPressureInsensitiveRash"},
	{"GeneralSymptoms", "HasDehydration"},
	{"GeneralSymptoms", "HasItchyRash"},
	{"GeneralSymptoms", "HasEdema"},
	{"GeneralSymptoms", "HasNightSweats"},
	{"GeneralSymptoms", "HasWeightLoss"},
	{"GeneralSymptoms", "HasExhaustion"},
	// RespiratorySymptoms
	{"RespiratorySymptoms", "HasCough"},
	{"RespiratorySymptoms", "HasSputum"},
	{"RespiratorySymptoms", "HasBloodySputum"},
	{"RespiratorySymptoms", "HasHemoptysis"},
	{"RespiratorySymptoms", "HasSoreThroat"},
	{"RespiratorySymptoms", "HasDryThroat"},
	{"RespiratorySymptoms", "HasNasalCongestion"},
	{"RespiratorySymptoms", "HasRunnyNose"},
	{"RespiratorySymptoms", "HasChestTightness"},
	{"RespiratorySymptoms", "HasShortnessOfBreath"},
	{"RespiratorySymptoms", "HasBreathingDifficulty"},
	{"RespiratorySymptoms", "HasHoarseness"},
	{"RespiratorySymptoms", "HasBronchitis"},
	// DigestiveSymptoms
	{"DigestiveSymptoms", "IsDiarrheaFrequencyGEThreeTimesPerDay"},
	{"DigestiveSymptoms", "HasStoolType1"},
	{"DigestiveSymptoms", "HasStoolType2"},
	{"DigestiveSymptoms", "HasVomiting"},
	{"DigestiveSymptoms", "HasNausea"},
	{"DigestiveSymptoms", "HasAppetiteLoss"},
	{"DigestiveSymptoms", "HasAbdominalDistension"},
	{"DigestiveSymptoms", "HasAbdominalPain"},
	{"DigestiveSymptoms", "HasBorborygmus"},
	{"DigestiveSymptoms", "HasUpperAbdominalDiscomfort"},
	{"DigestiveSymptoms", "HasConstipation"},
	{"DigestiveSymptoms", "HasOliguriaOrAnuria"},
	// CirculatorySymptoms
	{"CirculatorySymptoms", "HasArrhythmia"},
	{"CirculatorySymptoms", "HasChestPain"},
	{"CirculatorySymptoms", "HasRapidPulse"},
	{"CirculatorySymptoms", "HasPalpitation"},
	{"CirculatorySymptoms", "HasLowBloodPressure"},
	// NeurologicalSymptoms
	{"NeurologicalSymptoms", "HasDelirium"},
	{"NeurologicalSymptoms", "HasConvulsion"},
	{"NeurologicalSymptoms", "HasDrowsiness"},
	{"NeurologicalSymptoms", "HasComa"},
	{"NeurologicalSymptoms", "HasStiffNeck"},
	{"NeurologicalSymptoms", "HasRestlessness"},
	{"NeurologicalSymptoms", "HasMuscleParalysisNeckShoulder"},
	{"NeurologicalSymptoms", "HasMuscleParalysisLimbs"},
	{"NeurologicalSymptoms", "HasSwallowingDifficulty"},
	{"NeurologicalSymptoms", "HasSpeechDisorder"},
	{"NeurologicalSymptoms", "HasConsciousnessDisorder"},
	{"NeurologicalSymptoms", "HasDizziness"},
	{"NeurologicalSymptoms", "HasTinnitus"},
	{"NeurologicalSymptoms", "HasHearingLoss"},
	// LocalSymptoms
	{"LocalSymptoms", "HasConjunctivitis"},
	{"LocalSymptoms", "HasConjunctivalCongestion"},
	{"LocalSymptoms", "HasEyelidEdema"},
	{"LocalSymptoms", "HasSmellTasteLoss"},
	{"LocalSymptoms", "HasFacialRednessOrPallor"},
	{"LocalSymptoms", "HasRednessNeckChest"},
	{"LocalSymptoms", "HasLipCyanosis"},
	{"LocalSymptoms", "HasSkinPallorOrBruising"},
	{"LocalSymptoms", "HasGumBleeding"},
	{"LocalSymptoms", "HasKidneyPain"},
	{"LocalSymptoms", "HasLowBackPain"},
	{"LocalSymptoms", "HasOrbitalPain"},
	{"LocalSymptoms", "HasMucosalBleedingInOralOrNasal"},
	{"LocalSymptoms", "HasSkinBleedingPoints"},
	// OtherSymptoms
	{"OtherSymptoms", "HasSuddenOnset"},
	{"OtherSymptoms", "HasRapidProgress"},
	{"OtherSymptoms", "HasPeriodicAttack"},
	{"OtherSymptoms", "HasForcedPosture"},
	{"OtherSymptoms", "HasCalfMusclePain"},
	{"OtherSymptoms", "SleepQuality"},
	{"OtherSymptoms", "NutritionStatus"},
	{"OtherSymptoms", "WorkLifeStress"},
	{"OtherSymptoms", "SymptomSeverity"},
	{"OtherSymptoms", "HasOtherSymptoms"},
	// RiskFactorsAndExposure
	{"RiskFactorsAndExposure", "ContactWithFeverPatient"},
	{"RiskFactorsAndExposure", "ContactWithDiarrheaPatient"},
	{"RiskFactorsAndExposure", "ContactWithRashPatient"},
	{"RiskFactorsAndExposure", "ContactWithTuberculosisPatient"},
	{"RiskFactorsAndExposure", "HasDrinkingRawWater"},
	{"RiskFactorsAndExposure", "HasEatingRawFood"},
	{"RiskFactorsAndExposure", "HasEatingColdCookedFood"},
	{"RiskFactorsAndExposure", "HasEatingSeafood"},
	{"RiskFactorsAndExposure", "GroupOutbreak"},
	{"RiskFactorsAndExposure", "OutdoorStayOrWorkWithinMonth"},
	{"RiskFactorsAndExposure", "PlagueArea"},
	{"RiskFactorsAndExposure", "AnthraxArea"},
	{"RiskFactorsAndExposure", "MalariaArea"},
	{"RiskFactorsAndExposure", "ContactWithAnimalProducts"},
	{"RiskFactorsAndExposure", "ContactWithAnimals"},
	{"RiskFactorsAndExposure", "IsHillyOrMountainous"},
	{"RiskFactorsAndExposure", "IsPlain"},
	{"RiskFactorsAndExposure", "HasMosquitoBite"},
	{"RiskFactorsAndExposure", "TickBite"},
	{"RiskFactorsAndExposure", "FleaBite"},
	{"RiskFactorsAndExposure", "ContactWithRatIn2Months"},
	{"RiskFactorsAndExposure", "EatenFoodContaminatedByRatFeces"},
	{"RiskFactorsAndExposure", "DrunkWaterFromDitchesOrPonds"},
	{"RiskFactorsAndExposure", "RestedNearRatHoles"},
	{"RiskFactorsAndExposure", "RatOrRatDroppingsAtWorkplace"},
	{"RiskFactorsAndExposure", "ContactWithPatientExcreta"},
}

// diseaseWeightFocus maps disease -> field -> score. Remaining score is
// spread as tiny baseline so every slider has a row and totals equal 100.
var diseaseWeightFocus = map[string]map[string]float64{
	"流感": {
		"HasFever": 12, "HasChills": 6, "HasSweating": 3, "HasFatigue": 6,
		"HasHeadache": 6, "HasMusclePain": 8, "HasJointPain": 4,
		"HasCough": 10, "HasSputum": 3, "HasSoreThroat": 8, "HasDryThroat": 2,
		"HasNasalCongestion": 4, "HasRunnyNose": 4,
		"HasSuddenOnset": 6, "SymptomSeverity": 3,
		"ContactWithFeverPatient": 5, "GroupOutbreak": 5,
	},
	"新型冠状病毒感染": {
		"HasFever": 10, "HasFatigue": 6, "HasHeadache": 4, "HasMusclePain": 5,
		"HasCough": 10, "HasSoreThroat": 4, "HasDryThroat": 3,
		"HasChestTightness": 5, "HasShortnessOfBreath": 6, "HasBreathingDifficulty": 6,
		"HasSmellTasteLoss": 10, "HasSuddenOnset": 5,
		"ContactWithFeverPatient": 8, "GroupOutbreak": 8,
	},
	"感染性腹泻": {
		"HasFever": 4, "HasDehydration": 8, "HasExhaustion": 3,
		"IsDiarrheaFrequencyGEThreeTimesPerDay": 14, "HasStoolType1": 10, "HasStoolType2": 8,
		"HasVomiting": 8, "HasNausea": 5, "HasAbdominalPain": 8, "HasAbdominalDistension": 4,
		"HasAppetiteLoss": 3, "HasBorborygmus": 3,
		"HasDrinkingRawWater": 6, "HasEatingRawFood": 6, "ContactWithDiarrheaPatient": 5, "GroupOutbreak": 5,
	},
	"鼠疫": {
		"HasFever": 10, "HasChills": 6, "HasFatigue": 4, "HasHeadache": 4,
		"HasLymphNodeSwelling": 14, "HasPainfulRedRash": 4, "HasBloodBlisters": 4,
		"HasCough": 4, "HasBloodySputum": 5, "HasHemoptysis": 5,
		"HasSuddenOnset": 8, "HasRapidProgress": 8,
		"PlagueArea": 8, "FleaBite": 6, "ContactWithAnimals": 5, "ContactWithRatIn2Months": 5,
	},
	"炭疽": {
		"HasFever": 8, "HasFatigue": 4, "HasSkinUlcer": 12, "HasPainfulRedRash": 6,
		"HasBloodBlisters": 8, "HasItchyRash": 4,
		"HasCough": 4, "HasChestTightness": 4, "HasBreathingDifficulty": 5,
		"HasAbdominalPain": 4, "HasNausea": 3, "HasVomiting": 3,
		"HasSuddenOnset": 6, "HasRapidProgress": 5,
		"AnthraxArea": 10, "ContactWithAnimalProducts": 8, "ContactWithAnimals": 6,
	},
	"登革热（蚊媒传染病）": {
		"HasFever": 12, "HasHeadache": 6, "HasMusclePain": 6, "HasJointPain": 8,
		"HasCongestiveOrPetechialRash": 8, "HasPainfulRedRash": 5,
		"HasNausea": 4, "HasVomiting": 3, "HasAbdominalPain": 3,
		"HasConjunctivalCongestion": 4, "HasFacialRednessOrPallor": 4,
		"HasOrbitalPain": 6, "HasSuddenOnset": 6,
		"HasMosquitoBite": 12, "MalariaArea": 4, "OutdoorStayOrWorkWithinMonth": 3,
	},
	"疟疾（蚊媒传染病）": {
		"HasFever": 14, "HasChills": 12, "HasSweating": 8, "HasFatigue": 6,
		"HasHeadache": 5, "HasMusclePain": 4, "HasPeriodicAttack": 10,
		"HasNausea": 3, "HasVomiting": 3, "HasAbdominalPain": 3,
		"HasSuddenOnset":  4,
		"HasMosquitoBite": 14, "MalariaArea": 10, "OutdoorStayOrWorkWithinMonth": 4,
	},
	"森林脑炎（蜱媒传染病）": {
		"HasFever": 10, "HasHeadache": 8, "HasFatigue": 4, "HasMusclePain": 4,
		"HasStiffNeck": 8, "HasConsciousnessDisorder": 8, "HasDrowsiness": 5,
		"HasDelirium": 4, "HasConvulsion": 4, "HasDizziness": 4,
		"HasMuscleParalysisNeckShoulder": 6, "HasMuscleParalysisLimbs": 5,
		"HasSuddenOnset": 5,
		"TickBite":       14, "OutdoorStayOrWorkWithinMonth": 7, "IsHillyOrMountainous": 4,
	},
	"发热伴血小板减少综合征（蜱媒传染病）": {
		"HasFever": 12, "HasFatigue": 6, "HasHeadache": 5, "HasMusclePain": 5,
		"HasCongestiveOrPetechialRash": 6, "HasSubcutaneousAndMucosalBleedingSpots": 8,
		"HasGumBleeding": 5, "HasMucosalBleedingInOralOrNasal": 5, "HasSkinBleedingPoints": 5,
		"HasNausea": 3, "HasVomiting": 3, "HasAbdominalPain": 3, "HasOliguriaOrAnuria": 4,
		"HasSuddenOnset": 5,
		"TickBite":       12, "OutdoorStayOrWorkWithinMonth": 6, "IsHillyOrMountainous": 3,
	},
	"斑疹伤寒": {
		"HasFever": 12, "HasChills": 5, "HasHeadache": 8, "HasFatigue": 5,
		"HasCongestiveOrPetechialRash": 10, "HasPressureInsensitiveRash": 8,
		"HasPainfulRedRash": 5, "HasMusclePain": 4,
		"HasSuddenOnset": 6, "HasRapidProgress": 4,
		"FleaBite": 10, "ContactWithRatIn2Months": 8, "RatOrRatDroppingsAtWorkplace": 6,
		"OutdoorStayOrWorkWithinMonth": 4, "IsHillyOrMountainous": 5,
	},
	"流行性出血热": {
		"HasFever": 10, "HasHeadache": 5, "HasFatigue": 4, "HasMusclePain": 4,
		"HasFacialRednessOrPallor": 6, "HasRednessNeckChest": 5,
		"HasConjunctivalCongestion": 5, "HasEyelidEdema": 4,
		"HasKidneyPain": 8, "HasLowBackPain": 6, "HasOrbitalPain": 5,
		"HasOliguriaOrAnuria": 8, "HasSubcutaneousAndMucosalBleedingSpots": 5,
		"HasSuddenOnset":          5,
		"ContactWithRatIn2Months": 6, "EatenFoodContaminatedByRatFeces": 5,
		"RatOrRatDroppingsAtWorkplace": 5, "OutdoorStayOrWorkWithinMonth": 4,
	},
	"结核病": {
		"HasFever": 6, "HasNightSweats": 10, "HasWeightLoss": 10, "HasFatigue": 6,
		"HasCough": 12, "HasSputum": 6, "HasBloodySputum": 8, "HasHemoptysis": 6,
		"HasChestTightness": 4, "HasShortnessOfBreath": 4,
		"HasSuddenOnset":                 2,
		"ContactWithTuberculosisPatient": 14, "GroupOutbreak": 6, "ContactWithFeverPatient": 6,
	},
}

func seedFullWeights(db *gorm.DB) error {
	if err := db.Exec("DELETE FROM symptom_weighting").Error; err != nil {
		return err
	}
	if err := db.Exec("DELETE FROM ai_symptom_weighting").Error; err != nil {
		return err
	}

	totalRows := 0
	for _, name := range models.JavaScoringDiseaseNames {
		var dt models.DiseaseType
		if err := db.Where("disease_type_name = ?", name).First(&dt).Error; err == gorm.ErrRecordNotFound {
			dt = models.DiseaseType{DiseaseTypeName: name}
			if err = db.Create(&dt).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		focus := diseaseWeightFocus[name]
		if focus == nil {
			focus = map[string]float64{}
		}
		scores := make([]float64, len(weightCatalog))
		var focusSum float64
		zeroIdx := make([]int, 0, len(weightCatalog))
		for i, item := range weightCatalog {
			if v, ok := focus[item.Field]; ok && v > 0 {
				scores[i] = v
				focusSum += v
			} else {
				zeroIdx = append(zeroIdx, i)
			}
		}
		// Reserve baseline points for non-focus sliders so the UI is never all zeros.
		const reserveForBaseline = 18.0
		if n := len(zeroIdx); n > 0 {
			capFocus := 100 - reserveForBaseline
			if focusSum > capFocus && focusSum > 0 {
				scale := capFocus / focusSum
				focusSum = 0
				for i := range scores {
					if scores[i] > 0 {
						scores[i] = math.Round(scores[i]*scale*10) / 10
						focusSum += scores[i]
					}
				}
			}
			remain := 100 - focusSum
			if remain < 0 {
				remain = 0
			}
			each := math.Floor((remain/float64(n))*100) / 100
			if each < 0.1 {
				each = 0.1
			}
			for _, idx := range zeroIdx {
				scores[idx] = each
			}
			var total float64
			for _, s := range scores {
				total += s
			}
			diff := math.Round((100-total)*100) / 100
			if diff != 0 {
				// Adjust the largest focus weight so every slider stays positive.
				best := 0
				for i := 1; i < len(scores); i++ {
					if scores[i] > scores[best] {
						best = i
					}
				}
				scores[best] = math.Round((scores[best]+diff)*100) / 100
				if scores[best] < 0.1 {
					scores[best] = 0.1
				}
			}
		} else if focusSum != 100 {
			scores[len(scores)-1] += 100 - focusSum
		}

		rows := make([]models.SymptomWeighting, 0, len(weightCatalog))
		aiRows := make([]models.AISymptomWeighting, 0, len(weightCatalog))
		for i, item := range weightCatalog {
			score := scores[i]
			rows = append(rows, models.SymptomWeighting{
				DiseaseTypeID:    int(dt.DiseaseTypeID),
				DiseaseTypeName:  name,
				SourceTable:      item.Table,
				SymptomFieldName: item.Field,
				WeightScore:      score,
			})
			aiRows = append(aiRows, models.AISymptomWeighting{
				DiseaseTypeID:    int(dt.DiseaseTypeID),
				DiseaseTypeName:  name,
				SourceTable:      item.Table,
				SymptomFieldName: item.Field,
				WeightScore:      score,
			})
		}
		if err := db.CreateInBatches(&rows, 200).Error; err != nil {
			return err
		}
		if err := db.CreateInBatches(&aiRows, 200).Error; err != nil {
			return err
		}
		totalRows += len(rows)
	}

	fmt.Printf("full weights seeded: diseases=%d rows=%d catalog=%d\n",
		len(models.JavaScoringDiseaseNames), totalRows, len(weightCatalog))
	return nil
}
