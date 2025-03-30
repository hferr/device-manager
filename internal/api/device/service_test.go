package device_test

import (
	"fmt"
	"github/hferr/device-manager/internal/api/device"
	"github/hferr/device-manager/test"
	"github/hferr/device-manager/test/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreateDevice(t *testing.T) {
	var testCases = map[string]struct {
		wantErr bool
		repo    mock.DeviceRepository
		input   device.CreateDeviceRequest
	}{
		"successfully calls repo to insert device": {
			wantErr: false,
			repo: mock.DeviceRepository{
				InsertDeviceFunc: func(d *device.Device) error {
					return nil
				},
			},
			input: device.CreateDeviceRequest{
				Name:  "test",
				Brand: "test",
				State: device.StateAvailable,
			},
		},
		"repo returns error": {
			wantErr: true,
			repo: mock.DeviceRepository{
				InsertDeviceFunc: func(d *device.Device) error {
					return fmt.Errorf("boom")
				},
			},
			input: device.CreateDeviceRequest{
				Name:  "test",
				Brand: "test",
				State: device.StateAvailable,
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := device.NewService(&tc.repo)

			_, err := s.CreateDevice(tc.input)
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got %v", err)
			}

			if err == nil && tc.wantErr {
				t.Fatal("expected error, got none")
			}
		})
	}
}

func TestServiceDeleteDevice(t *testing.T) {
	var testCases = map[string]struct {
		wantErr bool
		repo    mock.DeviceRepository
		inputID uuid.UUID
	}{
		"successfully deletes device": {
			wantErr: false,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return &device.Device{ID: ID, State: device.StateAvailable}, nil
				},
				DeleteDeviceFunc: func(ID uuid.UUID) error {
					return nil
				},
			},
			inputID: uuid.New(),
		},
		"device is in use and cannot be deleted": {
			wantErr: true,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return &device.Device{ID: ID, State: device.StateInUse}, nil
				},
			},
			inputID: uuid.New(),
		},
		"repo returns error on delete": {
			wantErr: true,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return &device.Device{ID: ID, State: device.StateAvailable}, nil
				},
				DeleteDeviceFunc: func(ID uuid.UUID) error {
					return fmt.Errorf("boom")
				},
			},
			inputID: uuid.New(),
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := device.NewService(&tc.repo)

			err := s.DeleteDevice(tc.inputID)
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got %v", err)
			}

			if err == nil && tc.wantErr {
				t.Fatal("expected error, got none")
			}
		})
	}
}

func TestServiceListDevices(t *testing.T) {
	var deviceList = device.Devices{
		device.NewDevice("test", "test", device.StateAvailable),
		device.NewDevice("test", "test", device.StateInUse),
	}

	var testCases = map[string]struct {
		wantErr bool
		repo    mock.DeviceRepository
	}{
		"successfully lists devices": {
			wantErr: false,
			repo: mock.DeviceRepository{
				ListDevicesFunc: func() (device.Devices, error) {
					return deviceList, nil
				},
			},
		},
		"repo returns error": {
			wantErr: true,
			repo: mock.DeviceRepository{
				ListDevicesFunc: func() (device.Devices, error) {
					return nil, fmt.Errorf("boom")
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := device.NewService(&tc.repo)

			ds, err := s.ListDevices()
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got %v", err)
			}

			if err == nil {
				if tc.wantErr {
					t.Fatal("expected error, got none")
				}

				if len(deviceList) != len(ds) {
					t.Fatalf("expected %d devices, got %d", len(deviceList), len(ds))
				}
			}
		})
	}
}

func TestServiceFindByID(t *testing.T) {
	want := device.NewDevice("test", "test", device.StateAvailable)

	var testCases = map[string]struct {
		wantErr bool
		repo    mock.DeviceRepository
		inputID uuid.UUID
	}{
		"successfully finds device by ID": {
			wantErr: false,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return want, nil
				},
			},
			inputID: want.ID,
		},
		"repo returns error": {
			wantErr: true,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return nil, fmt.Errorf("boom")
				},
			},
			inputID: uuid.New(),
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := device.NewService(&tc.repo)

			got, err := s.FindByID(tc.inputID)
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got %v", err)
			}

			if err == nil {
				if tc.wantErr {
					t.Fatal("expected error, got none")
				}

				assert.Equal(t, got, want)
			}
		})
	}
}

func TestServiceFindByState(t *testing.T) {
	want := device.Devices{
		device.NewDevice("test", "test", device.StateAvailable),
		device.NewDevice("test", "test", device.StateAvailable),
	}

	var testCases = map[string]struct {
		wantErr bool
		repo    mock.DeviceRepository
		state   string
	}{
		"successfully finds devices by state": {
			wantErr: false,
			repo: mock.DeviceRepository{
				FindByStateFunc: func(state string) (device.Devices, error) {
					return want, nil
				},
			},
			state: device.StateAvailable,
		},
		"repo returns error": {
			wantErr: true,
			repo: mock.DeviceRepository{
				FindByStateFunc: func(state string) (device.Devices, error) {
					return nil, fmt.Errorf("boom")
				},
			},
			state: device.StateAvailable,
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := device.NewService(&tc.repo)

			got, err := s.FindByState(tc.state)
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got %v", err)
			}

			if err == nil {
				if tc.wantErr {
					t.Fatal("expected error, got none")
				}

				if len(want) != len(got) {
					t.Fatalf("expected %d devices, got %d", len(want), len(got))
				}
			}
		})
	}
}

func TestServiceFindByBrand(t *testing.T) {
	want := device.Devices{
		device.NewDevice("test", "test", device.StateAvailable),
		device.NewDevice("test", "test", device.StateAvailable),
	}

	var testCases = map[string]struct {
		wantErr bool
		repo    mock.DeviceRepository
		brand   string
	}{
		"successfully finds devices by brand": {
			wantErr: false,
			repo: mock.DeviceRepository{
				FindByBrandFunc: func(brand string) (device.Devices, error) {
					return want, nil
				},
			},
			brand: "test",
		},
		"repo returns error": {
			wantErr: true,
			repo: mock.DeviceRepository{
				FindByBrandFunc: func(brand string) (device.Devices, error) {
					return nil, fmt.Errorf("boom")
				},
			},
			brand: "test",
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := device.NewService(&tc.repo)

			got, err := s.FindByBrand(tc.brand)
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got %v", err)
			}

			if err == nil {
				if tc.wantErr {
					t.Fatal("expected error, got none")
				}

				if len(want) != len(got) {
					t.Fatalf("expected %d devices, got %d", len(want), len(got))
				}
			}
		})
	}
}

func TestServiceUpdateDevice(t *testing.T) {
	var testCases = map[string]struct {
		wantErr bool
		repo    mock.DeviceRepository
		inputID uuid.UUID
		input   device.UpdateDeviceRequest
	}{
		"successfully updates device": {
			wantErr: false,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return &device.Device{ID: ID, State: device.StateAvailable}, nil
				},
				UpdateDeviceFunc: func(d *device.Device) error {
					return nil
				},
			},
			inputID: uuid.New(),
			input: device.UpdateDeviceRequest{
				Name:  test.Ptr("updated-name"),
				Brand: test.Ptr("updated-brand"),
			},
		},
		"device is in use and cannot be updated": {
			wantErr: true,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return &device.Device{ID: ID, State: device.StateInUse}, nil
				},
			},
			inputID: uuid.New(),
			input: device.UpdateDeviceRequest{
				Name:  test.Ptr("updated-name"),
				Brand: test.Ptr("updated-brand"),
			},
		},
		"repo returns error on find": {
			wantErr: true,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return nil, fmt.Errorf("boom")
				},
			},
			inputID: uuid.New(),
			input: device.UpdateDeviceRequest{
				Name:  test.Ptr("updated-name"),
				Brand: test.Ptr("updated-brand"),
			},
		},
		"repo returns error on update": {
			wantErr: true,
			repo: mock.DeviceRepository{
				FindByIDFunc: func(ID uuid.UUID) (*device.Device, error) {
					return &device.Device{ID: ID, State: device.StateAvailable}, nil
				},
				UpdateDeviceFunc: func(d *device.Device) error {
					return fmt.Errorf("boom")
				},
			},
			inputID: uuid.New(),
			input: device.UpdateDeviceRequest{
				Name:  test.Ptr("updated-name"),
				Brand: test.Ptr("updated-brand"),
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := device.NewService(&tc.repo)

			err := s.UpdateDevice(tc.inputID, tc.input)
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got %v", err)
			}

			if err == nil && tc.wantErr {
				t.Fatal("expected error, got none")
			}
		})
	}
}
