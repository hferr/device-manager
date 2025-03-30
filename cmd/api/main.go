package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hferr/device-manager/config"
	"github.com/hferr/device-manager/internal/api/device"
	"github.com/hferr/device-manager/internal/protocols/httpjson"
	"github.com/hferr/device-manager/migrations"
	"github.com/hferr/device-manager/utils/validator"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const fmtDBConnString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	c := config.New()
	v := validator.New()

	db, err := setupDB(&c.DB)
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}

	// setup repos
	deviceRepo := device.NewRepository(db)

	// setup services
	deviceSvs := device.NewService(deviceRepo)

	// setup handlers
	handler := httpjson.NewHandler(deviceSvs, v)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      handler.NewRouter(),
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

func setupDB(cfg *config.ConfDB) (*gorm.DB, error) {
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
		return nil, err
	}

	dbHandle, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := migrations.MaybeApplyMigrations(dbHandle); err != nil {
		return nil, err
	}

	return db, nil
}
