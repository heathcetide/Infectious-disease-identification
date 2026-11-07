// Copyright (c) 2026 heathcetide. All rights reserved.

// Import promotional learning materials from the Java project's local folder
// into Qiniu Kodo, then insert DiseaseLearningMaterial rows (Link only) into SQLite.
//
//	go run ./cmd/import_materials
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/configs"
	"github.com/heathcetide/infectious-disease/internal/models"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
	"gorm.io/gorm"
)

func main() {
	srcDir := os.Getenv("MATERIALS_DIR")
	if srcDir == "" {
		srcDir = filepath.Join(os.Getenv("HOME"), "Desktop/auto-infect-detect-main/diseaseLearningMaterialPath")
	}
	if _, err := os.Stat(srcDir); err != nil {
		fmt.Fprintf(os.Stderr, "materials dir not found: %s (%v)\n", srcDir, err)
		os.Exit(1)
	}

	cfg, err := configs.Load("configs/config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}
	if cfg.Database.DSN == "" {
		fmt.Fprintln(os.Stderr, "database.dsn is empty")
		os.Exit(1)
	}

	if err := pkgstorage.InitWithOptions(
		cfg.Storage.Driver, cfg.Storage.Root, cfg.Storage.Region,
		cfg.Storage.Endpoint, cfg.Storage.AccessKey, cfg.Storage.SecretKey,
		cfg.Storage.PublicURLBase, cfg.Storage.MaxFileSize,
		pkgstorage.InitOptions{Private: cfg.Storage.Private},
	); err != nil {
		fmt.Fprintf(os.Stderr, "init storage: %v\n", err)
		os.Exit(1)
	}

	db, err := configs.InitDB(cfg.Database)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init db: %v\n", err)
		os.Exit(1)
	}
	if err := db.AutoMigrate(&models.DiseaseLearningMaterial{}, &models.DiseaseType{}); err != nil {
		fmt.Fprintf(os.Stderr, "migrate: %v\n", err)
		os.Exit(1)
	}
	// Drop legacy file_path column if still present (SQLite 3.35+).
	_ = db.Exec(`ALTER TABLE disease_learning_materials DROP COLUMN file_path`).Error

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read dir: %v\n", err)
		os.Exit(1)
	}

	diseases := models.JavaScoringDiseaseNames
	if len(diseases) == 0 {
		diseases = models.ScoringDiseaseNames
	}

	seq := nextSequence(db)
	uploaded := 0
	for i, ent := range entries {
		if ent.IsDir() {
			continue
		}
		name := ent.Name()
		ext := strings.ToLower(filepath.Ext(name))
		switch ext {
		case ".png", ".jpg", ".jpeg", ".gif", ".webp", ".mp4", ".pdf", ".mov":
		default:
			fmt.Printf("skip %s\n", name)
			continue
		}

		localPath := filepath.Join(srcDir, name)
		key := "infect/materials/" + name
		title, materialType := titleFor(ext, i)
		link := pkgstorage.StableURL(key)

		var existing models.DiseaseLearningMaterial
		existsDB := db.Where("link = ?", link).Limit(1).Find(&existing).Error == nil && existing.MaterialID > 0

		f, err := os.Open(localPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: %v\n", name, err)
			continue
		}
		info, _ := f.Stat()
		fmt.Printf("upload %s (%d bytes) → %s ...\n", name, info.Size(), key)
		if err := pkgstorage.Write(key, f); err != nil {
			_ = f.Close()
			fmt.Fprintf(os.Stderr, "  upload failed: %v\n", err)
			continue
		}
		_ = f.Close()

		link = pkgstorage.StableURL(key)
		disease := diseases[i%len(diseases)]
		diseaseID := diseaseTypeID(db, disease)
		now := time.Now()
		deleted := false

		if existsDB {
			_ = db.Model(&existing).Updates(map[string]any{
				"link":              link,
				"title":             title,
				"material_type":     materialType,
				"disease_type_name": disease,
				"disease_type_id":   diseaseID,
				"is_deleted":        false,
			}).Error
			fmt.Printf("  updated db materialId=%d link=%s\n", existing.MaterialID, link)
		} else {
			row := models.DiseaseLearningMaterial{
				DiseaseTypeID:   diseaseID,
				DiseaseTypeName: disease,
				Title:           title,
				MaterialType:    materialType,
				Link:            link,
				SequenceNumber:  seq,
				PublishDate:     &now,
				IsDeleted:       &deleted,
			}
			if err := db.Create(&row).Error; err != nil {
				fmt.Fprintf(os.Stderr, "  db insert failed: %v\n", err)
				continue
			}
			fmt.Printf("  inserted materialId=%d link=%s disease=%s\n", row.MaterialID, link, disease)
			seq++
		}
		uploaded++
	}

	fmt.Printf("done: uploaded/ensured %d materials\n", uploaded)
}

func nextSequence(db *gorm.DB) int {
	var maxSeq int
	_ = db.Model(&models.DiseaseLearningMaterial{}).Select("COALESCE(MAX(sequence_number),0)").Scan(&maxSeq)
	return maxSeq + 1
}

func diseaseTypeID(db *gorm.DB, name string) int {
	var dt models.DiseaseType
	if err := db.Where("disease_type_name = ?", name).First(&dt).Error; err != nil {
		return 0
	}
	return int(dt.DiseaseTypeID)
}

func titleFor(ext string, idx int) (title, materialType string) {
	switch ext {
	case ".mp4", ".mov":
		return fmt.Sprintf("宣传视频-%d", idx+1), "视频"
	case ".pdf":
		return fmt.Sprintf("宣传资料-%d", idx+1), "PDF"
	default:
		return fmt.Sprintf("宣传图片-%d", idx+1), "图片"
	}
}
