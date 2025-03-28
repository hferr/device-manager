package main

import (
	"fmt"
	"log"
	"net/http"

	"github/hferr/device-manager/config"
	"github/hferr/device-manager/internal/protocols/httpjson"
	"github/hferr/device-manager/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const fmtDBConnString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	c := config.New()

	if err := setupDB(&c.DB); err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      httpjson.NewRouter(),
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

func setupDB(cfg *config.ConfDB) error {
	dbConnString := fmt.Sprintf(
		fmtDBConnString,
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dbConnString))
	if err != nil {
		return err
	}

	dbHandle, err := db.DB()
	if err != nil {
		return err
	}

	if err := migrations.MaybeApplyMigrations(dbHandle); err != nil {
		return err
	}

	return nil
}
