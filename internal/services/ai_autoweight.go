// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"log"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

// GetAutoAIWeightUpdaterEnabled reads the auto weight updater flag.
func GetAutoAIWeightUpdaterEnabled(db *gorm.DB) (bool, error) {
	s, err := GetSystemSettings(db)
	if err != nil {
		return false, err
	}
	if s.AutoAIWeightUpdater == nil {
		return false, nil
	}
	return *s.AutoAIWeightUpdater, nil
}

// SetAutoAIWeightUpdaterEnabled toggles the auto weight updater.
func SetAutoAIWeightUpdaterEnabled(db *gorm.DB, enabled bool) error {
	s, err := GetSystemSettings(db)
	if err != nil {
		return err
	}
	s.AutoAIWeightUpdater = &enabled
	return db.Model(s).Update("auto_ai_weight_updater", enabled).Error
}

// RunAutoAIWeightUpdater is a stub that logs; recompute can be added later.
func RunAutoAIWeightUpdater(db *gorm.DB) error {
	log.Printf("ai autoweight: manual run triggered (stub)")
	_ = db
	return nil
}

// UpdateSymptomWeightsBatchWithDiagnosis updates symptom weights and mirrors to AI table.
func UpdateSymptomWeightsBatchWithDiagnosis(db *gorm.DB, items []SymptomWeightUpdateDTO) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := UpdateSymptomWeightsBatch(tx, items); err != nil {
			return err
		}
		for _, item := range items {
			_ = tx.Model(&models.AISymptomWeighting{}).
				Where("symptom_weighting_id = ?", item.SymptomWeightingID).
				Update("weight_score", item.WeightScore).Error
		}
		return nil
	})
}
