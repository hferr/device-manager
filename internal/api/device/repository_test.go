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

func TestFindByState(t *testing.T) {
	cleanup, db := test.SetupTestDBContainer(t)
	defer cleanup()

	repo := device.NewRepository(db)

	devicesInUse := []*device.Device{
		device.NewDevice("test", "test", "in_use"),
		device.NewDevice("test", "test", "in_use"),
		device.NewDevice("test", "test", "in_use"),
	}

	devicesAvailable := []*device.Device{
		device.NewDevice("test", "test", "available"),
		device.NewDevice("test", "test", "available"),
	}

	allDevices := append(devicesInUse, devicesAvailable...)

	for _, d := range allDevices {
		if err := repo.InsertDevice(d); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	}

	// fetch by state 'in_use'

	devicesInUseFromDB, err := repo.FindByState("in_use")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(devicesInUse) != len(devicesInUseFromDB) {
		t.Fatalf("expected %d devices 'in_use', got: %d", len(devicesInUse), len(devicesInUseFromDB))
	}

	// fetch by state 'available'

	devicesAvailableFromDB, err := repo.FindByState("available")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(devicesAvailable) != len(devicesAvailableFromDB) {
		t.Fatalf("expected %d devices 'in_use', got: %d", len(devicesAvailable), len(devicesAvailableFromDB))
	}
}

func TestFindByBrand(t *testing.T) {
	cleanup, db := test.SetupTestDBContainer(t)
	defer cleanup()

	repo := device.NewRepository(db)

	devicesWithBrand1 := []*device.Device{
		device.NewDevice("test", "cool_brand", "in_use"),
		device.NewDevice("test", "cool_brand", "in_use"),
		device.NewDevice("test", "cool_brand", "in_use"),
	}

	devicesWithBrand2 := []*device.Device{
		device.NewDevice("test", "nice_brand", "available"),
		device.NewDevice("test", "nice_brand", "available"),
	}

	allDevices := append(devicesWithBrand1, devicesWithBrand2...)

	for _, d := range allDevices {
		if err := repo.InsertDevice(d); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	}

	// fetch by brand 'cool_brand'

	devicesWithBrand1FromDB, err := repo.FindByBrand("cool_brand")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(devicesWithBrand1) != len(devicesWithBrand1FromDB) {
		t.Fatalf("expected %d devices, got: %d", len(devicesWithBrand1), len(devicesWithBrand1FromDB))
	}

	// fetch by brand 'nice_brand'

	devicesWithBrand2FromDB, err := repo.FindByBrand("nice_brand")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(devicesWithBrand2) != len(devicesWithBrand2FromDB) {
		t.Fatalf("expected %d devices, got: %d", len(devicesWithBrand2), len(devicesWithBrand2FromDB))
	}
}
