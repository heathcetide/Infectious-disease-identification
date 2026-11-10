// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// SymptomPhoto is a check-in skin/rash photo; YOLO/VLM labels are auxiliary features.
type SymptomPhoto struct {
	PhotoID  uint `json:"photoId" gorm:"primaryKey;autoIncrement"`
	UserID   uint `json:"userId" gorm:"index;not null"`
	StatusID uint `json:"statusId" gorm:"index"`

	FileName   string    `json:"fileName,omitempty" gorm:"size:255"`
	FilePath   string    `json:"filePath,omitempty" gorm:"size:500"`
	UploadDate time.Time `json:"uploadDate"`

	// LabelsJSON: [{"label":"painful_red_rash","symptomField":"HasPainfulRedRash","confidence":0.82}]
	LabelsJSON  string  `json:"labelsJson,omitempty" gorm:"type:text"`
	TopLabel    string  `json:"topLabel,omitempty" gorm:"size:80"`
	Confidence  float64 `json:"confidence,omitempty"`
	Classifier  string  `json:"classifier,omitempty" gorm:"size:40"` // yolo | llm_vision
	RawResponse string  `json:"rawResponse,omitempty" gorm:"type:text"`
}

func (SymptomPhoto) TableName() string { return constants.TableSymptomPhotos }

// VisionSample is a labeled training image for symptom YOLO / classifier retraining.
type VisionSample struct {
	SampleID uint `json:"sampleId" gorm:"primaryKey;autoIncrement"`

	FileName string `json:"fileName,omitempty" gorm:"size:255"`
	FilePath string `json:"filePath,omitempty" gorm:"size:500"`
	MimeType string `json:"mimeType,omitempty" gorm:"size:80"`

	// ClassLabel e.g. painful_red_rash | skin_ulcer | itchy_rash | petechial_rash | blood_blister | normal_skin | other
	ClassLabel   string `json:"classLabel" gorm:"size:80;index"`
	SymptomField string `json:"symptomField,omitempty" gorm:"size:80"` // maps to AllSymptoms bool field
	Split        string `json:"split,omitempty" gorm:"size:20"`        // train|val|test
	Source       string `json:"source,omitempty" gorm:"size:40"`       // manual|checkin|import
	Notes        string `json:"notes,omitempty" gorm:"size:500"`
	BBoxJSON     string `json:"bboxJson,omitempty" gorm:"type:text"` // optional YOLO boxes [{x,y,w,h,cls}]
	IsActive     bool   `json:"isActive" gorm:"default:true"`

	CreatedBy uint      `json:"createdBy,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (VisionSample) TableName() string { return constants.TableVisionSamples }
