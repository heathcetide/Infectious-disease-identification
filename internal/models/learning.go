// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// DiseaseLearningMaterial is educational content for a disease type.
// Files live on object storage; only the public CDN Link is persisted.
type DiseaseLearningMaterial struct {
	MaterialID      uint       `json:"materialId" gorm:"primaryKey;autoIncrement"`
	DiseaseTypeID   int        `json:"diseaseTypeId"`
	DiseaseTypeName string     `json:"diseaseTypeName,omitempty" gorm:"size:100"`
	Title           string     `json:"title,omitempty" gorm:"size:200"`
	MaterialType    string     `json:"materialType,omitempty" gorm:"size:50"`
	Link            string     `json:"link,omitempty" gorm:"size:500"`
	SequenceNumber  int        `json:"sequenceNumber"`
	PublishDate     *time.Time `json:"publishDate,omitempty"`
	IsDeleted       *bool      `json:"isDeleted"`
	DeletedDate     *time.Time `json:"deletedDate,omitempty"`
}

func (DiseaseLearningMaterial) TableName() string { return constants.TableDiseaseLearningMaterials }

// MaterialLearningRecord tracks when a user opened a learning material.
type MaterialLearningRecord struct {
	RecordID    uint      `json:"recordId" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"userId" gorm:"index;not null"`
	MaterialID  uint      `json:"materialId" gorm:"index;not null"`
	ClickedDate time.Time `json:"clickedDate"`
}

func (MaterialLearningRecord) TableName() string { return constants.TableMaterialLearningRecords }
