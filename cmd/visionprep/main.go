// Copyright (c) 2026 heathcetide. All rights reserved.
//
// Dataset curator for symptom vision (P4). Downloads public educational images into
// assets/vision and can sync metadata into vision_samples.
//
//	go run ./cmd/visionprep -download
//	go run ./cmd/visionprep -download -sync-db
//
// Note: Ultralytics YOLO *weight training* is Python-based; Go owns dataset/metadata/
// export/inference orchestration. After train, set SYMPTOM_YOLO_MODEL to best.pt.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/configs"
	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/internal/services"
)

var seedURLs = map[string][]string{
	"painful_red_rash": {
		"https://upload.wikimedia.org/wikipedia/commons/thumb/8/8a/Erythema_migrans.jpg/320px-Erythema_migrans.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/3/3a/Measles.jpg/320px-Measles.jpg",
	},
	"skin_ulcer": {
		"https://upload.wikimedia.org/wikipedia/commons/thumb/4/4c/Venous_ulcer_001.jpg/320px-Venous_ulcer_001.jpg",
	},
	"itchy_rash": {
		"https://upload.wikimedia.org/wikipedia/commons/thumb/6/6f/Atopic_dermatitis.jpg/320px-Atopic_dermatitis.jpg",
	},
	"petechial_rash": {
		"https://upload.wikimedia.org/wikipedia/commons/thumb/7/7d/Petechia.jpg/320px-Petechia.jpg",
	},
	"cyanosis": {
		"https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Cyanosis.jpg/320px-Cyanosis.jpg",
	},
	"normal_skin": {
		"https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Human_skin.jpg/320px-Human_skin.jpg",
	},
	"other": {
		"https://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Acne_vulgaris.jpg/320px-Acne_vulgaris.jpg",
	},
}

func main() {
	root := flag.String("assets", "assets/vision", "asset root")
	download := flag.Bool("download", false, "download seed images")
	synthetic := flag.Bool("synthetic", false, "generate placeholder JPEGs per class (offline)")
	syncDB := flag.Bool("sync-db", false, "register local images into vision_samples")
	manifestOut := flag.String("manifest", "assets/vision/manifest.json", "write local manifest")
	cfgPath := flag.String("config", "configs/config.yaml", "config for -sync-db")
	flag.Parse()

	_ = os.MkdirAll(*root, 0o755)
	if *download {
		if err := downloadSeeds(*root); err != nil {
			fmt.Fprintf(os.Stderr, "download: %v\n", err)
		}
	}
	if *synthetic {
		if err := writeSynthetic(*root, 6); err != nil {
			fatal(err)
		}
	}
	items, err := scanLocal(*root)
	if err != nil {
		fatal(err)
	}
	payload := map[string]any{
		"format":     "ultralytics-classify-manifest",
		"classNames": servicesClassNames(),
		"count":      len(items),
		"items":      items,
		"note":       "Pure-Go YOLO weight training is not practical; use scripts/vision/yolo_train.py after this prepare step. Inference can stay Go-orchestrated via SYMPTOM_YOLO_BIN.",
	}
	b, _ := json.MarshalIndent(payload, "", "  ")
	if err = os.WriteFile(*manifestOut, b, 0o644); err != nil {
		fatal(err)
	}
	fmt.Printf("wrote %s (%d items)\n", *manifestOut, len(items))

	if !*syncDB {
		return
	}
	cfg, err := configs.Load(*cfgPath)
	if err != nil {
		fatal(err)
	}
	db, err := configs.InitDB(cfg.Database)
	if err != nil {
		fatal(err)
	}
	_ = db.AutoMigrate(&models.VisionSample{})
	n := 0
	for _, it := range items {
		classLabel := it["classLabel"].(string)
		filePath := it["filePath"].(string)
		var cnt int64
		_ = db.Model(&models.VisionSample{}).Where("file_path = ? AND is_active = ?", filePath, true).Count(&cnt)
		if cnt > 0 {
			continue
		}
		sample := &models.VisionSample{
			FileName: it["fileName"].(string), FilePath: filePath,
			ClassLabel: classLabel, SymptomField: "",
			Split: it["split"].(string), Source: "assets", Notes: "visionprep",
		}
		if err := services.CreateVisionSample(db, sample); err != nil {
			fmt.Fprintf(os.Stderr, "db: %v\n", err)
			continue
		}
		n++
	}
	fmt.Printf("synced %d new samples at %s\n", n, time.Now().Format(time.RFC3339))
}

func servicesClassNames() []string {
	return []string{
		"painful_red_rash", "skin_ulcer", "itchy_rash", "petechial_rash",
		"blood_blister", "cyanosis", "normal_skin", "other",
	}
}

func downloadSeeds(root string) error {
	client := &http.Client{Timeout: 45 * time.Second}
	var errs []string
	for class, urls := range seedURLs {
		dir := filepath.Join(root, "train", class)
		_ = os.MkdirAll(dir, 0o755)
		for i, u := range urls {
			name := fmt.Sprintf("%s_%02d%s", class, i+1, extFromURL(u))
			dest := filepath.Join(dir, name)
			if st, err := os.Stat(dest); err == nil && st.Size() > 0 {
				continue
			}
			req, _ := http.NewRequest(http.MethodGet, u, nil)
			req.Header.Set("User-Agent", "infectious-disease-visionprep/1.0")
			resp, err := client.Do(req)
			if err != nil {
				errs = append(errs, err.Error())
				continue
			}
			if resp.StatusCode >= 400 {
				errs = append(errs, fmt.Sprintf("%s -> %d", u, resp.StatusCode))
				_ = resp.Body.Close()
				continue
			}
			f, err := os.Create(dest)
			if err != nil {
				_ = resp.Body.Close()
				errs = append(errs, err.Error())
				continue
			}
			_, copyErr := io.Copy(f, io.LimitReader(resp.Body, 8<<20))
			_ = f.Close()
			_ = resp.Body.Close()
			if copyErr != nil {
				_ = os.Remove(dest)
				errs = append(errs, copyErr.Error())
				continue
			}
			fmt.Println("downloaded", dest)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return fmt.Errorf("%s", strings.Join(errs, "; "))
}

func scanLocal(root string) ([]map[string]any, error) {
	var items []map[string]any
	for _, split := range []string{"train", "val", "test"} {
		splitDir := filepath.Join(root, split)
		entries, err := os.ReadDir(splitDir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !e.IsDir() {
				continue
			}
			class := e.Name()
			files, _ := os.ReadDir(filepath.Join(splitDir, class))
			for _, f := range files {
				if f.IsDir() {
					continue
				}
				name := f.Name()
				low := strings.ToLower(name)
				if !strings.HasSuffix(low, ".jpg") && !strings.HasSuffix(low, ".jpeg") && !strings.HasSuffix(low, ".png") && !strings.HasSuffix(low, ".webp") {
					continue
				}
				rel := filepath.ToSlash(filepath.Join(root, split, class, name))
				items = append(items, map[string]any{
					"classLabel": class,
					"split":      split,
					"fileName":   name,
					"filePath":   rel,
					"source":     "assets",
				})
			}
		}
	}
	return items, nil
}

func writeSynthetic(root string, perClass int) error {
	palette := map[string]color.RGBA{
		"painful_red_rash": {220, 60, 60, 255},
		"skin_ulcer":       {90, 50, 40, 255},
		"itchy_rash":       {240, 160, 120, 255},
		"petechial_rash":   {160, 30, 60, 255},
		"blood_blister":    {80, 10, 30, 255},
		"cyanosis":         {70, 90, 160, 255},
		"normal_skin":      {230, 190, 160, 255},
		"other":            {180, 180, 180, 255},
	}
	for class, base := range palette {
		for _, split := range []string{"train", "val"} {
			n := perClass
			if split == "val" {
				n = 2
			}
			dir := filepath.Join(root, split, class)
			_ = os.MkdirAll(dir, 0o755)
			for i := 0; i < n; i++ {
				dest := filepath.Join(dir, fmt.Sprintf("%s_%s_%02d.jpg", class, split, i+1))
				if st, err := os.Stat(dest); err == nil && st.Size() > 0 {
					continue
				}
				img := image.NewRGBA(image.Rect(0, 0, 224, 224))
				for y := 0; y < 224; y++ {
					for x := 0; x < 224; x++ {
						jitter := uint8((x*3 + y*5 + i*17) % 40)
						c := color.RGBA{
							R: clamp8(int(base.R) + int(jitter) - 20),
							G: clamp8(int(base.G) + int(jitter)/2 - 10),
							B: clamp8(int(base.B) + int(jitter)/3 - 5),
							A: 255,
						}
						// faux lesion blob
						dx, dy := x-110-i*3, y-110+i*2
						if dx*dx+dy*dy < 35*35 {
							c = color.RGBA{R: clamp8(int(base.R) - 40), G: clamp8(int(base.G) - 30), B: clamp8(int(base.B) - 20), A: 255}
						}
						img.Set(x, y, c)
					}
				}
				f, err := os.Create(dest)
				if err != nil {
					return err
				}
				if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 85}); err != nil {
					_ = f.Close()
					return err
				}
				_ = f.Close()
				fmt.Println("synthetic", dest)
			}
		}
	}
	return nil
}

func clamp8(v int) uint8 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v)
}

func extFromURL(u string) string {
	path := u
	if i := strings.Index(path, "?"); i >= 0 {
		path = path[:i]
	}
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp":
		return ext
	default:
		return ".jpg"
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
