// Copyright (c) 2026 heathcetide. All rights reserved.

// Seed creates a demo railway employee and baseline disease / weight rows for local testing.
//
//	go run ./cmd/seed
package main

import (
	"fmt"
	"os"

	"github.com/heathcetide/infectious-disease/internal/configs"
	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/password"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.Load("configs/config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}
	if cfg.Database.DSN == "" {
		fmt.Fprintln(os.Stderr, "database.dsn is empty")
		os.Exit(1)
	}
	db, err := configs.InitDB(cfg.Database)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init db: %v\n", err)
		os.Exit(1)
	}
	if err := db.AutoMigrate(models.AllModels()...); err != nil {
		fmt.Fprintf(os.Stderr, "migrate: %v\n", err)
		os.Exit(1)
	}
	if err := seed(db); err != nil {
		fmt.Fprintf(os.Stderr, "seed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("seed ok (password for all demo users: Passw0rd!)")
	fmt.Println("  railway  phone 13800000000")
	fmt.Println("  admin    phone 13800000001")
	fmt.Println("  cdc      phone 13800000002")
	fmt.Println("  medical  phone 13800000003")
}

func ensureUser(db *gorm.DB, phone, email, name, role string, hash string) error {
	var existing models.User
	err := db.Where("phone_number = ?", phone).First(&existing).Error
	if err == nil {
		return nil
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	u := models.User{
		Username:    phone,
		Email:       email,
		Name:        name,
		PhoneNumber: phone,
		Phone:       phone,
		Password:    hash,
		Role:        role,
		Status:      models.UserStatusActive,
	}
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	if role == models.UserRoleRailway {
		return db.Create(&models.EmployeeProfile{UserID: u.ID}).Error
	}
	return nil
}

func seed(db *gorm.DB) error {
	hash, err := password.Hash("Passw0rd!", nil)
	if err != nil {
		return err
	}

	demos := []struct {
		phone, email, name, role string
	}{
		{"13800000000", "railway-demo@local.dev", "演示职工", models.UserRoleRailway},
		{"13800000001", "admin-demo@local.dev", "演示管理员", models.UserRoleAdmin},
		{"13800000002", "cdc-demo@local.dev", "演示疾控", models.UserRoleCDC},
		{"13800000003", "medical-demo@local.dev", "演示医务", models.UserRoleMedical},
	}
	for _, d := range demos {
		if err := ensureUser(db, d.phone, d.email, d.name, d.role, hash); err != nil {
			return err
		}
	}

	var settingsCount int64
	if err := db.Model(&models.SystemSettings{}).Count(&settingsCount).Error; err != nil {
		return err
	}
	if settingsCount == 0 {
		off := false
		if err := db.Create(&models.SystemSettings{IsAI: &off}).Error; err != nil {
			return err
		}
	}

	for _, name := range models.JavaScoringDiseaseNames {
		var dt models.DiseaseType
		if err := db.Where("disease_type_name = ?", name).First(&dt).Error; err == gorm.ErrRecordNotFound {
			dt = models.DiseaseType{DiseaseTypeName: name}
			if err := db.Create(&dt).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		var wCount int64
		if err := db.Model(&models.SymptomWeighting{}).
			Where("disease_type_name = ?", name).
			Count(&wCount).Error; err != nil {
			return err
		}
		if wCount > 0 {
			continue
		}
		// Minimal demo weights so scoring returns non-zero for common fever/cough symptoms.
		demo := []models.SymptomWeighting{
			{DiseaseTypeID: int(dt.DiseaseTypeID), DiseaseTypeName: name, SourceTable: "generalsymptoms", SymptomFieldName: "HasFever", WeightScore: 5},
			{DiseaseTypeID: int(dt.DiseaseTypeID), DiseaseTypeName: name, SourceTable: "generalsymptoms", SymptomFieldName: "HasChills", WeightScore: 3},
			{DiseaseTypeID: int(dt.DiseaseTypeID), DiseaseTypeName: name, SourceTable: "generalsymptoms", SymptomFieldName: "HasFatigue", WeightScore: 2},
			{DiseaseTypeID: int(dt.DiseaseTypeID), DiseaseTypeName: name, SourceTable: "respiratorysymptoms", SymptomFieldName: "HasCough", WeightScore: 4},
			{DiseaseTypeID: int(dt.DiseaseTypeID), DiseaseTypeName: name, SourceTable: "respiratorysymptoms", SymptomFieldName: "HasSoreThroat", WeightScore: 2},
		}
		if err := db.Create(&demo).Error; err != nil {
			return err
		}
	}

	var hospCount int64
	if err := db.Model(&models.Hospital{}).Count(&hospCount).Error; err != nil {
		return err
	}
	if hospCount == 0 {
		if err := db.Create(&models.Hospital{
			HospitalName:        "示例铁路医院",
			HospitalPhoneNumber: "010-88886666",
			Address:             "北京市海淀区示例路100号",
		}).Error; err != nil {
			return err
		}
	}

	var matCount int64
	if err := db.Model(&models.DiseaseLearningMaterial{}).Count(&matCount).Error; err != nil {
		return err
	}
	if matCount == 0 {
		deleted := false
		mats := []models.DiseaseLearningMaterial{
			{DiseaseTypeName: "流感", Title: "流感防护须知", MaterialType: "text", Link: "", SequenceNumber: 1, IsDeleted: &deleted},
			{DiseaseTypeName: "新型冠状病毒感染", Title: "新冠防护要点", MaterialType: "text", SequenceNumber: 1, IsDeleted: &deleted},
		}
		if err := db.Create(&mats).Error; err != nil {
			return err
		}
	}
	return nil
}
