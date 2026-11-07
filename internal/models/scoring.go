// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// DiseaseScoring stores computed disease scores for a daily health status.
type DiseaseScoring struct {
	ScoringID uint `json:"scoringId" gorm:"primaryKey;autoIncrement"`
	StatusID  uint `json:"statusId" gorm:"index;not null"`

	SymptomTime     *time.Time `json:"symptomTime,omitempty" gorm:"type:date"`
	DiseaseTypeID   int        `json:"diseaseTypeId,omitempty"`
	DiseaseTypeName string     `json:"diseaseTypeName,omitempty" gorm:"size:100"`
	Score           float64    `json:"score"`
}

func (DiseaseScoring) TableName() string { return constants.TableDiseaseScoring }

// SymptomWeighting defines symptom field weights per disease type.
type SymptomWeighting struct {
	SymptomWeightingID uint `json:"symptomWeightingId" gorm:"primaryKey;autoIncrement"`

	DiseaseTypeID   int     `json:"diseaseTypeId"`
	DiseaseTypeName string  `json:"diseaseTypeName,omitempty" gorm:"size:100"`
	SourceTable     string  `json:"tableName,omitempty" gorm:"column:table_name;size:100"`
	SymptomFieldName string `json:"symptomFieldName,omitempty" gorm:"size:100"`
	WeightScore     float64 `json:"weightScore"`
}

func (SymptomWeighting) TableName() string { return constants.TableSymptomWeighting }

// AIDiseaseScoring stores AI-assisted disease scores for a daily health status.
type AIDiseaseScoring struct {
	ScoringID uint `json:"scoringId" gorm:"primaryKey;autoIncrement"`
	StatusID  uint `json:"statusId" gorm:"index;not null"`

	SymptomTime     *time.Time `json:"symptomTime,omitempty" gorm:"type:date"`
	DiseaseTypeID   int        `json:"diseaseTypeId,omitempty"`
	DiseaseTypeName string     `json:"diseaseTypeName,omitempty" gorm:"size:100"`
	Score           float64    `json:"score"`
}

func (AIDiseaseScoring) TableName() string { return constants.TableAIDiseaseScoring }

// AISymptomWeighting defines AI symptom field weights per disease type.
type AISymptomWeighting struct {
	SymptomWeightingID uint `json:"symptomWeightingId" gorm:"primaryKey;autoIncrement"`

	DiseaseTypeID    int     `json:"diseaseTypeId"`
	DiseaseTypeName  string  `json:"diseaseTypeName,omitempty" gorm:"size:100"`
	SourceTable      string  `json:"tableName,omitempty" gorm:"column:table_name;size:100"`
	SymptomFieldName string  `json:"symptomFieldName,omitempty" gorm:"size:100"`
	WeightScore      float64 `json:"weightScore"`
}

func (AISymptomWeighting) TableName() string { return constants.TableAISymptomWeighting }

// DiseaseType is a top-level infectious disease category.
type DiseaseType struct {
	DiseaseTypeID   uint   `json:"diseaseTypeId" gorm:"primaryKey;autoIncrement"`
	DiseaseTypeName string `json:"diseaseTypeName" gorm:"size:100;not null"`
	HasSubtype      *bool  `json:"hasSubtype"`
}

func (DiseaseType) TableName() string { return constants.TableDiseaseTypes }

// DiseaseSubtype is a sub-classification under a disease type.
type DiseaseSubtype struct {
	SubtypeID     uint   `json:"subtypeId" gorm:"primaryKey;autoIncrement"`
	DiseaseTypeID uint   `json:"diseaseTypeId" gorm:"index;not null"`
	SubtypeName   string `json:"subtypeName" gorm:"size:100;not null"`
}

func (DiseaseSubtype) TableName() string { return constants.TableDiseaseSubtypes }

// SystemSettings holds global application toggles.
type SystemSettings struct {
	ID                  uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	IsAI                *bool    `json:"isAI"`
	AutoAIWeightUpdater *bool    `json:"autoAIWeightUpdater"`
	PredictorRange      *float64 `json:"predictorRange"`
}

func (SystemSettings) TableName() string { return constants.TableSystemSettings }
