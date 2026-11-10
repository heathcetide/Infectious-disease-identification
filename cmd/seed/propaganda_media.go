// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/configs"
	"github.com/heathcetide/infectious-disease/internal/models"
	pkgstorage "github.com/heathcetide/infectious-disease/pkg/storage"
)

func seedPropagandaMediaToOSS(db *gorm.DB) error {
	if err := seedPropagandaMediaLocal(db); err != nil {
		fmt.Printf("local propaganda seed: %v\n", err)
	}
	cfg, err := configs.Load("configs/config.yaml")
	if err != nil {
		return nil // local rows already inserted
	}
	if strings.TrimSpace(cfg.Storage.AccessKey) == "" || strings.TrimSpace(cfg.Storage.SecretKey) == "" {
		fmt.Println("skip OSS propaganda upload: empty storage keys")
		return nil
	}
	if err = pkgstorage.InitWithOptions(
		cfg.Storage.Driver, cfg.Storage.Root, cfg.Storage.Region,
		cfg.Storage.Endpoint, cfg.Storage.AccessKey, cfg.Storage.SecretKey,
		cfg.Storage.PublicURLBase, cfg.Storage.MaxFileSize,
		pkgstorage.InitOptions{Private: cfg.Storage.Private},
	); err != nil {
		fmt.Printf("skip OSS propaganda upload: %v\n", err)
		return nil
	}

	root, err := os.Getwd()
	if err != nil {
		return nil
	}
	mediaDir := filepath.Join(root, "data", "seed_materials")
	script := filepath.Join(root, "cmd", "seed", "gen_propaganda_media.py")
	if _, err = os.Stat(filepath.Join(mediaDir, "manifest.tsv")); err != nil {
		fmt.Println("generating propaganda posters/videos …")
		cmd := exec.Command("python3", script)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Run(); err != nil {
			fmt.Printf("generate media skipped: %v\n", err)
			return nil
		}
	}

	manifestPath := filepath.Join(mediaDir, "manifest.tsv")
	f, err := os.Open(manifestPath)
	if err != nil {
		return nil
	}
	defer func() { _ = f.Close() }()

	uploaded := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, "\t")
		if len(parts) < 4 {
			continue
		}
		name := parts[3]
		localPath := filepath.Join(mediaDir, name)
		key := "infect/materials/" + name
		file, err := os.Open(localPath)
		if err != nil {
			continue
		}
		fmt.Printf("upload %s → %s\n", name, key)
		if err := pkgstorage.Write(key, file); err != nil {
			_ = file.Close()
			fmt.Printf("upload %s failed: %v (keep local link)\n", name, err)
			return nil
		}
		_ = file.Close()
		uploaded++
	}
	fmt.Printf("propaganda media uploaded to OSS: %d\n", uploaded)
	return nil
}

// seedPropagandaMediaLocal binds generated posters/videos/pdfs via infect-files API URLs
// so mini-app can preview without Qiniu.
func seedPropagandaMediaLocal(db *gorm.DB) error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	mediaDir := filepath.Join(root, "data", "seed_materials")
	manifestPath := filepath.Join(mediaDir, "manifest.tsv")
	if _, err = os.Stat(manifestPath); err != nil {
		script := filepath.Join(root, "cmd", "seed", "gen_propaganda_media.py")
		cmd := exec.Command("python3", script)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Run(); err != nil {
			return fmt.Errorf("generate media: %w", err)
		}
	}
	f, err := os.Open(manifestPath)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	typeIDByName := map[string]int{}
	var dts []models.DiseaseType
	_ = db.Find(&dts).Error
	for _, d := range dts {
		typeIDByName[d.DiseaseTypeName] = int(d.DiseaseTypeID)
	}

	off := false
	inserted := 0
	seq := 100
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, "\t")
		if len(parts) < 4 {
			continue
		}
		typ, disease, title, name := parts[0], parts[1], parts[2], parts[3]
		if _, err := os.Stat(filepath.Join(mediaDir, name)); err != nil {
			continue
		}
		var existing int64
		_ = db.Model(&models.DiseaseLearningMaterial{}).
			Where("title = ? AND disease_type_name = ?", title, disease).
			Count(&existing)
		if existing > 0 {
			continue
		}
		link := "http://127.0.0.1:8080/api/v1/infect-files/learning/" + name
		pub := time.Now().AddDate(0, 0, -(inserted % 40))
		row := models.DiseaseLearningMaterial{
			DiseaseTypeID:   typeIDByName[disease],
			DiseaseTypeName: disease,
			Title:           title,
			MaterialType:    typ,
			Link:            link,
			SequenceNumber:  seq,
			PublishDate:     &pub,
			IsDeleted:       &off,
		}
		if err := db.Create(&row).Error; err != nil {
			return err
		}
		inserted++
		seq++
	}
	if err := sc.Err(); err != nil {
		return err
	}
	fmt.Printf("propaganda media local bound: %d\n", inserted)
	return nil
}
