package test

import (
	"context"
	"fmt"
	"github/hferr/device-manager/migrations"
	"log"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupTestDBContainer(t *testing.T) (func(), *gorm.DB) {
	// create a postgres container for testing
	r := testcontainers.ContainerRequest{
		Image:        "postgres:alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "topsecretpassword",
			"POSTGRES_DB":       "device_manager_test",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(60 * time.Second),
	}

	ctx := context.Background()

	pc, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: r,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("error setting up postgres test container: %v", err)
	}

	host, _ := pc.Host(ctx)
	port, _ := pc.MappedPort(ctx, "5432")

	connString := fmt.Sprintf(
		"host=%s port=%s user=postgres password=topsecretpassword dbname=device_manager_test sslmode=disable",
		host,
		port.Port(),
	)

	db, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	dbHandle, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect to test database handle: %v", err)
	}

	// run migrations
	if err := migrations.MaybeApplyMigrations(dbHandle); err != nil {
		log.Fatalf("failed to run migrations on test database: %v", err)
	}

	// cleanup before each test
	db.Exec("DELETE FROM devices")

	// terminate container after tests
	cleanup := func() {
		pc.Terminate(ctx)
	}

	return cleanup, db
}
