package device_test

import (
	"github/hferr/device-manager/internal/api/device"
	"github/hferr/device-manager/test"
	"testing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestInsertDevice(t *testing.T) {
	cleanup, db := test.SetupTestDBContainer(t)
	defer cleanup()

	repo := device.NewRepository(db)

	// assert valid device creation

	d := device.NewDevice("test", "test", "available")

	if err := repo.InsertDevice(d); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// assert invalid device creation (invalid state)

	invalidDevice := device.NewDevice("test", "test", "invalid")
	if err := repo.InsertDevice(invalidDevice); err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestListDevices(t *testing.T) {
	cleanup, db := test.SetupTestDBContainer(t)
	defer cleanup()

	repo := device.NewRepository(db)

	wantedDeviceListLen := 4
	for range wantedDeviceListLen {
		err := repo.InsertDevice(
			device.NewDevice("test", "test", "in_use"),
		)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	}

	ds, err := repo.ListDevices()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if wantedDeviceListLen != len(ds) {
		t.Fatalf("wanted device list len to be %d, got %d", wantedDeviceListLen, len(ds))
	}
}

func TestFindByID(t *testing.T) {
	cleanup, db := test.SetupTestDBContainer(t)
	defer cleanup()

	repo := device.NewRepository(db)

	// create and assert device when finding by ID

	device := device.NewDevice("test", "test", "in_use")
	if err := repo.InsertDevice(device); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	deviceFromDB, err := repo.FindByID(device.ID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if device.ID != deviceFromDB.ID {
		t.Fatalf("expected device ID: %s, got: %s", device.ID, deviceFromDB.ID)
	}

	// assert device not found case

	_, err = repo.FindByID(uuid.New())
	if err == nil {
		t.Fatal("expected error, got none")
	}

	if err != gorm.ErrRecordNotFound {
		t.Fatalf("expected error: %v, got: %v", gorm.ErrRecordNotFound, err)
	}
}

func TestDeleteDevice(t *testing.T) {
	cleanup, db := test.SetupTestDBContainer(t)
	defer cleanup()

	repo := device.NewRepository(db)

	d := device.NewDevice("test", "test", "available")
	if err := repo.InsertDevice(d); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	err := repo.DeleteDevice(d.ID)
	if err != nil {
		t.Fatal("expected error, got none")
	}
}
