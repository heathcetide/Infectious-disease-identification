// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	"gorm.io/gorm"
)

// LearningMaterialPageItem matches Java DiseaseLearningMaterialsPageVO (no filePath).
type LearningMaterialPageItem struct {
	MaterialID      uint       `json:"materialId"`
	Title           string     `json:"title,omitempty"`
	Link            string     `json:"link,omitempty"`
	PublishDate     *time.Time `json:"publishDate,omitempty"`
	StudyCount      int64      `json:"studyCount"`
	SequenceNumber  int        `json:"sequenceNumber"`
	IsDeleted       *bool      `json:"isDeleted"`
	MaterialType    string     `json:"materialType,omitempty"`
	DiseaseTypeName string     `json:"diseaseTypeName,omitempty"`
	DiseaseTypeID   int        `json:"diseaseTypeId,omitempty"`
}

// PageSelectLearningMaterials returns paginated materials filtered by disease type name.
// Matches Java: empty diseaseTypeName = all; order SequenceNumber DESC; include studyCount.
func PageSelectLearningMaterials(db *gorm.DB, dto LearningMaterialsPageDTO) (*response.Page, error) {
	page, size := dto.PageNo, dto.PageSize
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	q := db.Model(&models.DiseaseLearningMaterial{})
	if name := strings.TrimSpace(dto.DiseaseTypeName); name != "" {
		q = q.Where("disease_type_name = ?", name)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []models.DiseaseLearningMaterial
	offset := (page - 1) * size
	if err := q.Order("sequence_number DESC").Offset(offset).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}

	items := make([]LearningMaterialPageItem, 0, len(list))
	for _, m := range list {
		var studyCount int64
		_ = db.Model(&models.MaterialLearningRecord{}).
			Where("material_id = ?", m.MaterialID).
			Count(&studyCount).Error
		items = append(items, LearningMaterialPageItem{
			MaterialID:      m.MaterialID,
			Title:           m.Title,
			Link:            m.Link,
			PublishDate:     m.PublishDate,
			StudyCount:      studyCount,
			SequenceNumber:  m.SequenceNumber,
			IsDeleted:       m.IsDeleted,
			MaterialType:    m.MaterialType,
			DiseaseTypeName: m.DiseaseTypeName,
			DiseaseTypeID:   m.DiseaseTypeID,
		})
	}
	return response.NewPage(items, total, page, size), nil
}
