package device_test

import (
	"github/hferr/device-manager/internal/api/device"
	"github/hferr/device-manager/test"
	"testing"
)

func TestInsertDevice(t *testing.T) {
	var testCases = map[string]struct {
		wantErr bool
		d       *device.Device
	}{
		"succesfully insert new device": {
			wantErr: false,
			d:       device.NewDevice("test", "test", "available"),
		},
		"attempt to insert device with invalid state": {
			wantErr: true,
			d:       device.NewDevice("test", "test", "invalid"),
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cleanup, db := test.SetupTestDBContainer(t)
			defer cleanup()

			repo := device.NewRepository(db)

			err := repo.InsertDevice(tc.d)
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got: %v", err)
			}

			if err == nil && tc.wantErr {
				t.Fatalf("expected error, got nil")
			}
		})
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
