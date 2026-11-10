// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// ScoringVersion snapshots prompt/weight/mode for reproducibility (P5).
type ScoringVersion struct {
	VersionID   uint   `json:"versionId" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"size:120"`
	Mode        string `json:"mode" gorm:"size:20"` // expert | ai | hybrid
	PromptHash  string `json:"promptHash,omitempty" gorm:"size:64"`
	PromptText  string `json:"promptText,omitempty" gorm:"type:text"`
	WeightsHash string `json:"weightsHash,omitempty" gorm:"size:64"`
	// SnapshotJSON holds full restore payload (mode + prompt + expert/AI weights).
	SnapshotJSON string    `json:"snapshotJson,omitempty" gorm:"type:longtext"`
	Notes        string    `json:"notes,omitempty" gorm:"size:500"`
	IsActive     bool      `json:"isActive" gorm:"default:false"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (ScoringVersion) TableName() string { return constants.TableScoringVersions }

// EvalRun is one Top-K accuracy evaluation over confirmed diagnoses.
type EvalRun struct {
	RunID        uint       `json:"runId" gorm:"primaryKey;autoIncrement"`
	VersionID    uint       `json:"versionId" gorm:"index"`
	Name         string     `json:"name" gorm:"size:120"`
	K            int        `json:"k" gorm:"default:3"`
	CaseCount    int        `json:"caseCount"`
	ExpertHitAt1 float64    `json:"expertHitAt1"`
	ExpertHitAtK float64    `json:"expertHitAtK"`
	AIHitAt1     float64    `json:"aiHitAt1"`
	AIHitAtK     float64    `json:"aiHitAtK"`
	DateFrom     *time.Time `json:"dateFrom,omitempty" gorm:"type:date"`
	DateTo       *time.Time `json:"dateTo,omitempty" gorm:"type:date"`
	CreatedAt    time.Time  `json:"createdAt"`
}

func (EvalRun) TableName() string { return constants.TableEvalRuns }

// EvalCase stores per-diagnosis Top-K outcome.
type EvalCase struct {
	CaseID        uint   `json:"caseId" gorm:"primaryKey;autoIncrement"`
	RunID         uint   `json:"runId" gorm:"index;not null"`
	DiagnosisID   uint   `json:"diagnosisId" gorm:"index"`
	UserID        uint   `json:"userId" gorm:"index"`
	GoldDisease   string `json:"goldDisease" gorm:"size:100"`
	ExpertTopJSON string `json:"expertTopJson,omitempty" gorm:"type:text"`
	AITopJSON     string `json:"aiTopJson,omitempty" gorm:"type:text"`
	ExpertHitAt1  bool   `json:"expertHitAt1"`
	ExpertHitAtK  bool   `json:"expertHitAtK"`
	AIHitAt1      bool   `json:"aiHitAt1"`
	AIHitAtK      bool   `json:"aiHitAtK"`
}

func (EvalCase) TableName() string { return constants.TableEvalCases }
