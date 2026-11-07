// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"reflect"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// symptomFieldAliases maps Go struct field names to SymptomWeighting.SymptomFieldName when they differ.
var symptomFieldAliases = map[string]string{
	"ArrheaFrequencyGEThreeTimesPerDay": "IsDiarrheaFrequencyGEThreeTimesPerDay",
	"IsGroupDining":                       "IsGroupDining",
	"IsDiningOut":                         "IsDiningOut",
	"IsDiningUnclear":                     "IsDiningUnclear",
	"IsHillyOrMountainous":                "IsHillyOrMountainous",
	"IsPlain":                             "IsPlain",
	"IsOtherResidenceType":                "IsOtherResidenceType",
}

var itchyRashLocationFields = []string{
	"ItchyRashOnFingers",
	"ItchyRashOnBackOfHands",
	"ItchyRashOnUpperLimbs",
	"ItchyRashOnLowerLimbs",
	"ItchyRashOnFeet",
	"ItchyRashOnFace",
	"ItchyRashOnOther",
}

// activeSymptomFields collects PascalCase SymptomFieldName values for all true boolean fields in symptoms.
func activeSymptomFields(symptoms AllSymptoms) map[string]struct{} {
	active := make(map[string]struct{})
	general := symptoms.Generalsymptoms
	collectBoolFields(reflect.ValueOf(general), active, map[string]struct{}{
		"GeneralSymptomsID": {},
		"StatusID":          {},
	})
	collectBoolFields(reflect.ValueOf(symptoms.Respiratorysymptoms), active, map[string]struct{}{
		"RespiratorySymptomsID": {}, "StatusID": {},
	})
	digestive := symptoms.Digestivesymptoms.ToModel()
	collectBoolFields(reflect.ValueOf(digestive), active, map[string]struct{}{
		"DigestiveSymptomsID": {}, "StatusID": {},
	})
	collectBoolFields(reflect.ValueOf(symptoms.Circulatorysymptoms), active, map[string]struct{}{
		"CirculatorySymptomsID": {}, "StatusID": {},
	})
	collectBoolFields(reflect.ValueOf(symptoms.NeurologicalSymptoms), active, map[string]struct{}{
		"NeurologicalSymptomsID": {}, "StatusID": {},
	})
	collectBoolFields(reflect.ValueOf(symptoms.LocalSymptoms), active, map[string]struct{}{
		"LocalSymptomsID": {}, "StatusID": {},
	})
	collectBoolFields(reflect.ValueOf(symptoms.OtherSymptoms), active, map[string]struct{}{
		"OtherSymptomsID": {}, "StatusID": {},
	})
	risk := symptoms.RiskFactorsAndExposure.ToModel()
	collectBoolFields(reflect.ValueOf(risk), active, map[string]struct{}{
		"RiskFactorsAndExposureID": {}, "StatusID": {},
	})

	if itchyRashLocationActive(general) {
		active["HasPressureInsensitiveRash"] = struct{}{}
	}
	return active
}

func itchyRashLocationActive(g models.GeneralSymptoms) bool {
	v := reflect.ValueOf(g)
	for _, name := range itchyRashLocationFields {
		f := v.FieldByName(name)
		if f.IsValid() && f.Kind() == reflect.Bool && f.Bool() {
			return true
		}
	}
	return false
}

func collectBoolFields(v reflect.Value, active map[string]struct{}, skip map[string]struct{}) {
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		if _, ok := skip[field.Name]; ok {
			continue
		}
		fv := v.Field(i)
		if fv.Kind() == reflect.Bool && fv.Bool() {
			name := field.Name
			if alias, ok := symptomFieldAliases[name]; ok {
				name = alias
			}
			active[name] = struct{}{}
		}
	}
}

func scoreForDisease(weights []models.SymptomWeighting, diseaseName string, active map[string]struct{}) float64 {
	var total float64
	for _, w := range weights {
		if !strings.EqualFold(w.DiseaseTypeName, diseaseName) {
			continue
		}
		if _, ok := active[w.SymptomFieldName]; ok {
			total += w.WeightScore
		}
	}
	return total
}

// ComputeScores calculates weighted disease scores, persists rows, and returns them.
func ComputeScores(db *gorm.DB, statusID uint, symptoms AllSymptoms) ([]models.DiseaseScoring, error) {
	var weights []models.SymptomWeighting
	if err := db.Find(&weights).Error; err != nil {
		return nil, err
	}

	active := activeSymptomFields(symptoms)
	today := time.Now()
	dayStart := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	if err := db.Where("status_id = ? AND symptom_time = ?", statusID, dayStart).
		Delete(&models.DiseaseScoring{}).Error; err != nil {
		return nil, err
	}

	diseaseNames := models.JavaScoringDiseaseNames
	out := make([]models.DiseaseScoring, 0, len(diseaseNames))

	for _, diseaseName := range diseaseNames {
		score := scoreForDisease(weights, diseaseName, active)
		typeID := 0
		for _, w := range weights {
			if strings.EqualFold(w.DiseaseTypeName, diseaseName) {
				typeID = w.DiseaseTypeID
				break
			}
		}
		row := models.DiseaseScoring{
			StatusID:        statusID,
			SymptomTime:     &dayStart,
			DiseaseTypeID:   typeID,
			DiseaseTypeName: diseaseName,
			Score:           score,
		}
		if err := db.Create(&row).Error; err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}
