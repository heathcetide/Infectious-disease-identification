// Copyright (c) 2026 heathcetide. All rights reserved.

// One-shot: AutoMigrate GORM models onto the configured database (MySQL target).
package main

import (
	"fmt"
	"os"

	"github.com/heathcetide/infectious-disease/internal/configs"
	"github.com/heathcetide/infectious-disease/internal/models"
)

func main() {
	cfg, err := configs.Load("configs/config.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "load config:", err)
		os.Exit(1)
	}
	if cfg.Database.Driver != "mysql" {
		fmt.Fprintf(os.Stderr, "expected mysql driver, got %q\n", cfg.Database.Driver)
		os.Exit(1)
	}
	db, err := configs.InitDB(cfg.Database)
	if err != nil {
		fmt.Fprintln(os.Stderr, "connect:", err)
		os.Exit(1)
	}
	if err := db.AutoMigrate(models.AllModels()...); err != nil {
		fmt.Fprintln(os.Stderr, "automigrate:", err)
		os.Exit(1)
	}
	fmt.Println("schema ok on", cfg.Database.Driver)
}
