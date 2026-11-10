// Copyright (c) 2026 heathcetide. All rights reserved.

package handlers

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/internal/services"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	respgin "github.com/heathcetide/infectious-disease/pkg/common/response/gin"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
)

// CheckInUploadSymptomPhoto POST /api/v1/checkin/symptom-photo
func (h *Handlers) CheckInUploadSymptomPhoto(c *gin.Context) {
	userID, ok := h.currentUserID(c)
	if !ok {
		return
	}
	key, name, err := h.uploadMultipartFile(c, "symptom-photos")
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	var statusID uint
	if v := c.PostForm("statusId"); v != "" {
		if id, e := strconv.ParseUint(v, 10, 64); e == nil {
			statusID = uint(id)
		}
	}
	photo, labels, classErr := services.SaveAndClassifySymptomPhoto(h.db, h.svc.LLM, userID, statusID, key, name)
	if photo == nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, "classify failed"))
		return
	}
	respgin.Created(c, gin.H{
		"photo":     photo,
		"labels":    labels,
		"publicUrl": pkgstorage.PublicURL(key),
		"classifyError": func() string {
			if classErr != nil {
				return classErr.Error()
			}
			return ""
		}(),
	})
}

// AdminVisionSamplesPage GET /api/v1/admin/vision/samples
func (h *Handlers) AdminVisionSamplesPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	rows, total, err := services.PageVisionSamples(h.db, c.Query("classLabel"), c.Query("split"), page, size)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	list := make([]gin.H, 0, len(rows))
	for _, r := range rows {
		item := gin.H{
			"sampleId": r.SampleID, "fileName": r.FileName, "filePath": r.FilePath,
			"classLabel": r.ClassLabel, "symptomField": r.SymptomField, "split": r.Split,
			"source": r.Source, "notes": r.Notes, "bboxJson": r.BBoxJSON,
			"isActive": r.IsActive, "createdAt": r.CreatedAt,
			"previewUrl": "/api/v1/admin/vision/samples/" + strconv.FormatUint(uint64(r.SampleID), 10) + "/file",
			"publicUrl":  pkgstorage.PublicURL(r.FilePath),
		}
		list = append(list, item)
	}
	respgin.Success(c, gin.H{"list": list, "total": total, "page": page, "pageSize": size})
}

// AdminVisionSampleCreate POST /api/v1/admin/vision/samples
func (h *Handlers) AdminVisionSampleCreate(c *gin.Context) {
	userID, _ := h.currentUserID(c)
	key, name, err := h.uploadMultipartFile(c, "vision-dataset")
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	sample := &models.VisionSample{
		FileName: name, FilePath: key,
		ClassLabel:   strings.TrimSpace(c.PostForm("classLabel")),
		SymptomField: strings.TrimSpace(c.PostForm("symptomField")),
		Split:        strings.TrimSpace(c.DefaultPostForm("split", "train")),
		Source:       strings.TrimSpace(c.DefaultPostForm("source", "manual")),
		Notes:        c.PostForm("notes"),
		BBoxJSON:     c.PostForm("bboxJson"),
		CreatedBy:    userID,
	}
	if err := services.CreateVisionSample(h.db, sample); err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Created(c, gin.H{"sample": sample, "publicUrl": pkgstorage.PublicURL(key)})
}

// AdminVisionDatasetSummary GET /api/v1/admin/vision/summary
func (h *Handlers) AdminVisionDatasetSummary(c *gin.Context) {
	out, err := services.ExportVisionDatasetSummary(h.db)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Success(c, out)
}

// AdminVisionSampleDelete DELETE /api/v1/admin/vision/samples/:id
func (h *Handlers) AdminVisionSampleDelete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	_ = h.db.Model(&models.VisionSample{}).Where("sample_id = ?", id).
		Updates(map[string]any{"is_active": false, "updated_at": time.Now()}).Error
	respgin.Success(c, gin.H{"ok": true})
}

// AdminVisionDatasetExport GET /api/v1/admin/vision/export
func (h *Handlers) AdminVisionDatasetExport(c *gin.Context) {
	rows, err := services.ExportVisionDatasetManifest(h.db)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	classes := make([]string, 0)
	summary, _ := services.ExportVisionDatasetSummary(h.db)
	if summary != nil {
		if cn, ok := summary["classNames"].([]string); ok {
			classes = cn
		}
	}
	respgin.Success(c, gin.H{
		"format":     "ultralytics-classify-manifest",
		"classNames": classes,
		"items":      rows,
		"count":      len(rows),
		"trainHint":  "python scripts/vision/yolo_train.py --manifest export.json --out runs/symptom-cls",
	})
}

// AdminVisionSampleFile GET /api/v1/admin/vision/samples/:id/file
func (h *Handlers) AdminVisionSampleFile(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	var sample models.VisionSample
	if err := h.db.First(&sample, id).Error; err != nil || !sample.IsActive {
		respgin.WriteError(c, response.Err(response.CodeNotFound))
		return
	}
	path := sample.FilePath
	if path == "" {
		respgin.WriteError(c, response.Err(response.CodeNotFound))
		return
	}
	// Local assets or absolute path
	if strings.HasPrefix(path, "assets/") || strings.HasPrefix(path, "/") {
		c.File(path)
		return
	}
	r, _, err := pkgstorage.Read(path)
	if err != nil {
		// fallback: treat as workspace-relative
		c.File(path)
		return
	}
	defer func() { _ = r.Close() }()
	c.Header("Content-Type", "image/jpeg")
	_, _ = io.Copy(c.Writer, r)
}

// AdminVisionClassifyTrial POST /api/v1/admin/vision/classify
// Upload an image to try the current YOLO (or LLM vision fallback) classifier.
func (h *Handlers) AdminVisionClassifyTrial(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	f, err := file.Open()
	if err != nil {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	defer func() { _ = f.Close() }()
	raw, err := io.ReadAll(io.LimitReader(f, 8<<20))
	if err != nil || len(raw) == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	mime := file.Header.Get("Content-Type")
	if mime == "" {
		mime = "image/jpeg"
	}
	labels, classifier, rawResp, classErr := services.ClassifySymptomImage(c.Request.Context(), h.svc.LLM, mime, raw)
	out := gin.H{
		"labels":     labels,
		"classifier": classifier,
		"raw":        rawResp,
		"yoloEnv":    os.Getenv("SYMPTOM_YOLO_MODEL") != "",
	}
	if classErr != nil {
		out["error"] = classErr.Error()
		respgin.Success(c, out)
		return
	}
	respgin.Success(c, out)
}

// AdminVisionPromotePhoto POST /api/v1/admin/vision/promote
func (h *Handlers) AdminVisionPromotePhoto(c *gin.Context) {
	userID, _ := h.currentUserID(c)
	var body struct {
		PhotoID    uint   `json:"photoId"`
		ClassLabel string `json:"classLabel"`
		Split      string `json:"split"`
		Notes      string `json:"notes"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.PhotoID == 0 {
		respgin.WriteError(c, response.Err(response.CodeBadRequest))
		return
	}
	sample, err := services.PromoteSymptomPhotoToSample(h.db, body.PhotoID, userID, body.ClassLabel, body.Split, body.Notes)
	if err != nil {
		respgin.WriteError(c, response.New(response.CodeBadRequest, err.Error()))
		return
	}
	respgin.Created(c, sample)
}
