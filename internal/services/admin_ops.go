// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/password"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	"gorm.io/gorm"
)

const defaultResetPassword = "Passw0rd!"

// PageDiseaseTypes paginates disease types with subtypes for admin basic-data UI.
func PageDiseaseTypes(db *gorm.DB, q PageQuery) (*response.Page, error) {
	page, size := q.Normalize()
	query := db.Model(&models.DiseaseType{})
	kw := q.Keyword
	if kw == "" {
		kw = q.Text
	}
	if kw != "" {
		query = query.Where("disease_type_name LIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	var list []models.DiseaseType
	if err := query.Order("disease_type_id ASC").Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]DiseaseTypePageItem, 0, len(list))
	for _, d := range list {
		item := DiseaseTypePageItem{
			DiseaseTypeID:   d.DiseaseTypeID,
			DiseaseTypeName: d.DiseaseTypeName,
			HasSubtype:      d.HasSubtype != nil && *d.HasSubtype,
			SubDiseaseList:  []DiseaseSubtypeItem{},
		}
		var subs []models.DiseaseSubtype
		if err := db.Where("disease_type_id = ?", d.DiseaseTypeID).
			Order("subtype_id ASC").Find(&subs).Error; err != nil {
			return nil, err
		}
		for _, s := range subs {
			item.SubDiseaseList = append(item.SubDiseaseList, DiseaseSubtypeItem{
				SubtypeID:   s.SubtypeID,
				SubtypeName: s.SubtypeName,
			})
		}
		if len(item.SubDiseaseList) > 0 {
			item.HasSubtype = true
		}
		out = append(out, item)
	}
	return response.NewPage(out, total, page, size), nil
}

// GetHospitalByID loads one hospital.
func GetHospitalByID(db *gorm.DB, id uint) (*models.Hospital, error) {
	var h models.Hospital
	if err := db.First(&h, id).Error; err != nil {
		return nil, err
	}
	return &h, nil
}

// CreateHospital inserts a hospital.
func CreateHospital(db *gorm.DB, h models.Hospital) (*models.Hospital, error) {
	if err := db.Create(&h).Error; err != nil {
		return nil, err
	}
	return &h, nil
}

// UpdateHospital updates a hospital by ID.
func UpdateHospital(db *gorm.DB, id uint, h models.Hospital) (*models.Hospital, error) {
	h.HospitalID = id
	var existing models.Hospital
	if err := db.First(&existing, id).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&existing).Updates(map[string]any{
		"hospital_name":         h.HospitalName,
		"hospital_phone_number": h.HospitalPhoneNumber,
		"address":               h.Address,
	}).Error; err != nil {
		return nil, err
	}
	return GetHospitalByID(db, id)
}

// DeleteHospital removes a hospital.
func DeleteHospital(db *gorm.DB, id uint) error {
	res := db.Delete(&models.Hospital{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// SetUserActive toggles user status.
func SetUserActive(db *gorm.DB, id uint, active bool) error {
	status := models.UserStatusDisabled
	if active {
		status = models.UserStatusActive
	}
	res := db.Model(&models.User{}).Where("id = ?", id).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}

// ResetUserPassword sets password to defaultResetPassword.
func ResetUserPassword(db *gorm.DB, id uint) error {
	hash, err := password.Hash(defaultResetPassword, nil)
	if err != nil {
		return err
	}
	return models.User{}.UpdatePassword(db, id, hash)
}

// UpdateSymptomWeightsBatch updates many weight rows.
func UpdateSymptomWeightsBatch(db *gorm.DB, items []SymptomWeightUpdateDTO) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := UpdateSymptomWeight(tx, item); err != nil {
				return err
			}
		}
		return nil
	})
}

// GetFeedbackByID loads one feedback row.
func GetFeedbackByID(db *gorm.DB, id uint) (*models.UserFeedback, error) {
	var fb models.UserFeedback
	if err := db.First(&fb, id).Error; err != nil {
		return nil, err
	}
	return &fb, nil
}

// GetSystemSettings loads or creates the singleton settings row.
func GetSystemSettings(db *gorm.DB) (*models.SystemSettings, error) {
	var s models.SystemSettings
	err := db.Order("id ASC").First(&s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		off := false
		s = models.SystemSettings{IsAI: &off}
		if err := db.Create(&s).Error; err != nil {
			return nil, err
		}
		return &s, nil
	}
	return &s, err
}

// SetSystemAIMode persists AI toggle.
func SetSystemAIMode(db *gorm.DB, isAI bool) (*models.SystemSettings, error) {
	s, err := GetSystemSettings(db)
	if err != nil {
		return nil, err
	}
	s.IsAI = &isAI
	if err := db.Model(s).Update("is_ai", isAI).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// AIModeDTO is the API shape for AI settings.
type AIModeDTO struct {
	IsAI bool `json:"isAI"`
}

// GetAIModeDTO reads AI mode from settings.
func GetAIModeDTO(db *gorm.DB) (AIModeDTO, error) {
	s, err := GetSystemSettings(db)
	if err != nil {
		return AIModeDTO{}, err
	}
	if s.IsAI == nil {
		return AIModeDTO{IsAI: false}, nil
	}
	return AIModeDTO{IsAI: *s.IsAI}, nil
}

// --- CSV export (bytes) ---

func ExportUsersCSV(db *gorm.DB) ([]byte, error) {
	var users []models.User
	if err := db.Order("id ASC").Find(&users).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"id", "username", "email", "name", "phoneNumber", "role", "status"})
	for _, u := range users {
		phone := u.PhoneNumber
		if phone == "" {
			phone = u.Phone
		}
		_ = w.Write([]string{
			strconv.FormatUint(uint64(u.ID), 10),
			u.Username, u.Email, u.Name, phone, u.Role, strconv.Itoa(u.Status),
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

func ExportHospitalsCSV(db *gorm.DB) ([]byte, error) {
	var list []models.Hospital
	if err := db.Order("hospital_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"hospitalId", "hospitalName", "hospitalPhoneNumber", "address"})
	for _, h := range list {
		_ = w.Write([]string{
			strconv.FormatUint(uint64(h.HospitalID), 10),
			h.HospitalName, h.HospitalPhoneNumber, h.Address,
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

func HospitalImportTemplateCSV() []byte {
	return []byte("hospitalName,hospitalPhoneNumber,address\n示例医院,010-12345678,北京市示例区1号\n")
}

func ExportDiseasesCSV(db *gorm.DB) ([]byte, error) {
	list, err := ListDiseaseTypes(db)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"diseaseTypeId", "diseaseTypeName", "hasSubtype"})
	for _, d := range list {
		hs := "false"
		if d.HasSubtype != nil && *d.HasSubtype {
			hs = "true"
		}
		_ = w.Write([]string{
			strconv.FormatUint(uint64(d.DiseaseTypeID), 10),
			d.DiseaseTypeName, hs,
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

func ExportSymptomWeightsCSV(db *gorm.DB, disease string) ([]byte, error) {
	q := db.Model(&models.SymptomWeighting{})
	if disease != "" {
		q = q.Where("disease_type_name = ?", disease)
	}
	var list []models.SymptomWeighting
	if err := q.Order("symptom_weighting_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"symptomWeightingId", "diseaseTypeName", "tableName", "symptomFieldName", "weightScore"})
	for _, row := range list {
		_ = w.Write([]string{
			strconv.FormatUint(uint64(row.SymptomWeightingID), 10),
			row.DiseaseTypeName, row.SourceTable, row.SymptomFieldName,
			strconv.FormatFloat(row.WeightScore, 'f', -1, 64),
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

func SymptomWeightsTemplateCSV() []byte {
	return []byte("symptomWeightingId,weightScore\n1,5.0\n")
}

func ExportFeedbackCSV(db *gorm.DB) ([]byte, error) {
	var list []models.UserFeedback
	if err := db.Order("feedback_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"feedbackId", "userId", "name", "phoneNumber", "feedbackTitle", "feedbackText", "feedbackDate"})
	for _, fb := range list {
		_ = w.Write([]string{
			strconv.FormatUint(uint64(fb.FeedbackID), 10),
			strconv.FormatUint(uint64(fb.UserID), 10),
			fb.Name, fb.PhoneNumber, fb.FeedbackTitle, fb.FeedbackText,
			fb.FeedbackDate.Format(time.RFC3339),
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

// ImportUsersCSV upserts users from CSV (username,email,name,phoneNumber,role).
func ImportUsersCSV(db *gorm.DB, r io.Reader) (int, error) {
	rows, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return 0, err
	}
	if len(rows) < 2 {
		return 0, fmt.Errorf("empty csv")
	}
	header := rows[0]
	col := func(name string) int {
		for i, h := range header {
			if strings.EqualFold(strings.TrimSpace(h), name) {
				return i
			}
		}
		return -1
	}
	idxUser := col("username")
	idxEmail := col("email")
	idxName := col("name")
	idxPhone := col("phoneNumber")
	idxRole := col("role")
	if idxUser < 0 || idxEmail < 0 {
		return 0, fmt.Errorf("csv must include username and email columns")
	}
	defaultHash, err := password.Hash(defaultResetPassword, nil)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, row := range rows[1:] {
		if len(row) <= idxUser {
			continue
		}
		username := strings.TrimSpace(row[idxUser])
		email := strings.TrimSpace(row[idxEmail])
		if username == "" || email == "" {
			continue
		}
		var u models.User
		err := db.Where("username = ?", username).First(&u).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u = models.User{
				Username: username, Email: email, Password: defaultHash,
				Role: models.UserRoleRailway, Status: models.UserStatusActive,
			}
			if idxName >= 0 && idxName < len(row) {
				u.Name = strings.TrimSpace(row[idxName])
			}
			if idxPhone >= 0 && idxPhone < len(row) {
				u.PhoneNumber = strings.TrimSpace(row[idxPhone])
				u.Phone = u.PhoneNumber
			}
			if idxRole >= 0 && idxRole < len(row) {
				if r := strings.TrimSpace(row[idxRole]); r != "" {
					u.Role = r
				}
			}
			if err := db.Create(&u).Error; err != nil {
				return count, err
			}
			count++
			continue
		}
		if err != nil {
			return count, err
		}
		updates := map[string]any{"email": email}
		if idxName >= 0 && idxName < len(row) {
			updates["name"] = strings.TrimSpace(row[idxName])
		}
		if idxPhone >= 0 && idxPhone < len(row) {
			p := strings.TrimSpace(row[idxPhone])
			updates["phone_number"] = p
			updates["phone"] = p
		}
		if idxRole >= 0 && idxRole < len(row) {
			if r := strings.TrimSpace(row[idxRole]); r != "" {
				updates["role"] = r
			}
		}
		if err := db.Model(&u).Updates(updates).Error; err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

// ImportHospitalsCSV imports hospitals.
func ImportHospitalsCSV(db *gorm.DB, r io.Reader) (int, error) {
	rows, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return 0, err
	}
	if len(rows) < 2 {
		return 0, fmt.Errorf("empty csv")
	}
	count := 0
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 1 {
			continue
		}
		name := strings.TrimSpace(row[0])
		if name == "" {
			continue
		}
		phone, addr := "", ""
		if len(row) > 1 {
			phone = strings.TrimSpace(row[1])
		}
		if len(row) > 2 {
			addr = strings.TrimSpace(row[2])
		}
		h := models.Hospital{HospitalName: name, HospitalPhoneNumber: phone, Address: addr}
		if err := db.Create(&h).Error; err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

// ImportSymptomWeightsCSV updates weight scores by id.
func ImportSymptomWeightsCSV(db *gorm.DB, r io.Reader) (int, error) {
	rows, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return 0, err
	}
	count := 0
	for i, row := range rows {
		if i == 0 && len(row) > 0 && strings.Contains(strings.ToLower(row[0]), "symptom") {
			continue
		}
		if len(row) < 2 {
			continue
		}
		id, err1 := strconv.ParseUint(strings.TrimSpace(row[0]), 10, 64)
		score, err2 := strconv.ParseFloat(strings.TrimSpace(row[1]), 64)
		if err1 != nil || err2 != nil {
			continue
		}
		if err := UpdateSymptomWeight(db, SymptomWeightUpdateDTO{
			SymptomWeightingID: uint(id), WeightScore: score,
		}); err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}
