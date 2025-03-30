package httpjson_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/hferr/device-manager/internal/api/device"
	"github/hferr/device-manager/internal/protocols/httpjson"
	"github/hferr/device-manager/test"
	"github/hferr/device-manager/test/mock"
	"github/hferr/device-manager/utils/validator"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestHandlerCreateFile(t *testing.T) {
	var testCases = map[string]struct {
		wantCode int
		input    device.CreateDeviceRequest
		s        mock.DeviceService
	}{
		"succesfully calls device service": {
			wantCode: http.StatusCreated,
			input: device.CreateDeviceRequest{
				Name:  "test",
				Brand: "test",
				State: device.StateAvailable,
			},
			s: mock.DeviceService{
				CreateDeviceFunc: func(input device.CreateDeviceRequest) (*device.Device, error) {
					return device.NewDevice(input.Name, input.Brand, input.State), nil
				},
			},
		},
		"bad request - no device name provided": {
			wantCode: http.StatusBadRequest,
			input: device.CreateDeviceRequest{
				Brand: "test",
				State: device.StateAvailable,
			},
			s: mock.DeviceService{},
		},
		"bad request - no device brand name provided": {
			wantCode: http.StatusBadRequest,
			input: device.CreateDeviceRequest{
				Name:  "test",
				State: device.StateAvailable,
			},
			s: mock.DeviceService{},
		},
		"bad request - invalid device state provided": {
			wantCode: http.StatusBadRequest,
			input: device.CreateDeviceRequest{
				Name:  "test",
				Brand: "test",
				State: "invalid",
			},
			s: mock.DeviceService{},
		},
		"service returns error": {
			wantCode: http.StatusInternalServerError,
			input: device.CreateDeviceRequest{
				Name:  "test",
				Brand: "test",
				State: device.StateAvailable,
			},
			s: mock.DeviceService{
				CreateDeviceFunc: func(input device.CreateDeviceRequest) (*device.Device, error) {
					return nil, fmt.Errorf("boom")
				},
			},
		},
	}

	v := validator.New()

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			reqJson, err := json.Marshal(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			handler := httpjson.NewHandler(&tc.s, v)
			resp := test.DoHttpRequest(
				handler,
				http.MethodPost,
				"/devices",
				bytes.NewReader(reqJson),
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Fatalf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}
		})
	}
}

func TestHandlerListDevices(t *testing.T) {
	wantDs := device.Devices{
		&device.Device{Name: "test1", Brand: "brand1", State: device.StateAvailable},
		&device.Device{Name: "test2", Brand: "brand2", State: device.StateAvailable},
	}

	var testCases = map[string]struct {
		wantCode int
		s        mock.DeviceService
	}{
		"successfully lists devices": {
			wantCode: http.StatusOK,
			s: mock.DeviceService{
				ListDevicesFunc: func() (device.Devices, error) {
					return wantDs, nil
				},
			},
		},
		"service returns error": {
			wantCode: http.StatusInternalServerError,
			s: mock.DeviceService{
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

			handler := httpjson.NewHandler(&tc.s, nil)
			resp := test.DoHttpRequest(
				handler,
				http.MethodGet,
				"/devices",
				nil,
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Fatalf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}

			if tc.wantCode == http.StatusOK {
				respBody := &device.Devices{}
				err := json.NewDecoder(resp.Body).Decode(respBody)
				if err != nil {
					t.Fatal(err)
				}

				if len(*respBody) != len(wantDs) {
					t.Fatalf("expected % devices, got: %d", len(*respBody), len(wantDs))
				}
			}
		})
	}
}

func TestHandlerUpdateDevice(t *testing.T) {
	var testCases = map[string]struct {
		wantCode int
		input    device.UpdateDeviceRequest
		s        mock.DeviceService
	}{
		"successfully updates device": {
			wantCode: http.StatusOK,
			input: device.UpdateDeviceRequest{
				Name:  test.Ptr("updated"),
				Brand: test.Ptr("updated"),
			},
			s: mock.DeviceService{
				UpdateDeviceFunc: func(id uuid.UUID, input device.UpdateDeviceRequest) error {
					return nil
				},
			},
		},
		"bad request - invalid state": {
			wantCode: http.StatusBadRequest,
			input: device.UpdateDeviceRequest{
				Name:  test.Ptr("updated"),
				Brand: test.Ptr("updated"),
				State: test.Ptr("invalid"),
			},
			s: mock.DeviceService{
				UpdateDeviceFunc: func(id uuid.UUID, input device.UpdateDeviceRequest) error {
					return nil
				},
			},
		},
		"device is in use error": {
			wantCode: http.StatusUnprocessableEntity,
			input: device.UpdateDeviceRequest{
				Name:  test.Ptr("updated"),
				Brand: test.Ptr("updated"),
			},
			s: mock.DeviceService{
				UpdateDeviceFunc: func(id uuid.UUID, input device.UpdateDeviceRequest) error {
					return device.ErrDeviceInUse
				},
			},
		},
		"service returns error": {
			wantCode: http.StatusInternalServerError,
			input: device.UpdateDeviceRequest{
				Name:  test.Ptr("updated"),
				Brand: test.Ptr("updated"),
			},
			s: mock.DeviceService{
				UpdateDeviceFunc: func(id uuid.UUID, input device.UpdateDeviceRequest) error {
					return fmt.Errorf("boom")
				},
			},
		},
	}

	v := validator.New()

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			reqJson, err := json.Marshal(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			handler := httpjson.NewHandler(&tc.s, v)
			resp := test.DoHttpRequest(
				handler,
				http.MethodPatch,
				"/devices/"+uuid.New().String(),
				bytes.NewReader(reqJson),
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Fatalf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}
		})
	}
}

func TestHandlerFindByID(t *testing.T) {
	var testCases = map[string]struct {
		wantCode int
		s        mock.DeviceService
	}{
		"successfully finds device by ID": {
			wantCode: http.StatusOK,
			s: mock.DeviceService{
				FindByIDFunc: func(id uuid.UUID) (*device.Device, error) {
					return device.NewDevice("test", "brand", device.StateAvailable), nil
				},
			},
		},
		"device not found": {
			wantCode: http.StatusNotFound,
			s: mock.DeviceService{
				FindByIDFunc: func(id uuid.UUID) (*device.Device, error) {
					return nil, gorm.ErrRecordNotFound
				},
			},
		},
		"service returns error": {
			wantCode: http.StatusInternalServerError,
			s: mock.DeviceService{
				FindByIDFunc: func(id uuid.UUID) (*device.Device, error) {
					return nil, fmt.Errorf("boom")
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			handler := httpjson.NewHandler(&tc.s, nil)
			resp := test.DoHttpRequest(
				handler,
				http.MethodGet,
				"/devices/"+uuid.New().String(),
				nil,
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Fatalf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}
		})
	}
}

func TestHandlerFindByBrand(t *testing.T) {
	wantDs := device.Devices{
		&device.Device{Name: "test1", Brand: "brand1", State: device.StateAvailable},
		&device.Device{Name: "test2", Brand: "brand1", State: device.StateAvailable},
	}

	var testCases = map[string]struct {
		wantCode int
		s        mock.DeviceService
	}{
		"successfully finds devices by brand": {
			wantCode: http.StatusOK,
			s: mock.DeviceService{
				FindByBrandFunc: func(brand string) (device.Devices, error) {
					return wantDs, nil
				},
			},
		},
		"devices not found": {
			wantCode: http.StatusNotFound,
			s: mock.DeviceService{
				FindByBrandFunc: func(brand string) (device.Devices, error) {
					return nil, gorm.ErrRecordNotFound
				},
			},
		},
		"service returns error": {
			wantCode: http.StatusInternalServerError,
			s: mock.DeviceService{
				FindByBrandFunc: func(brand string) (device.Devices, error) {
					return nil, fmt.Errorf("boom")
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			handler := httpjson.NewHandler(&tc.s, nil)
			resp := test.DoHttpRequest(
				handler,
				http.MethodGet,
				"/devices/brand/test-brand",
				nil,
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Fatalf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}

			if tc.wantCode == http.StatusOK {
				respBody := &device.Devices{}
				err := json.NewDecoder(resp.Body).Decode(respBody)
				if err != nil {
					t.Fatal(err)
				}

				if len(*respBody) != len(wantDs) {
					t.Fatalf("expected % devices, got: %d", len(*respBody), len(wantDs))
				}
			}
		})
	}
}

func TestHandlerFindByState(t *testing.T) {
	wantDs := device.Devices{
		&device.Device{Name: "test1", Brand: "brand", State: device.StateAvailable},
		&device.Device{Name: "test2", Brand: "brand", State: device.StateAvailable},
	}

	var testCases = map[string]struct {
		wantCode int
		s        mock.DeviceService
	}{
		"successfully finds devices by state": {
			wantCode: http.StatusOK,
			s: mock.DeviceService{
				FindByStateFunc: func(state string) (device.Devices, error) {
					return wantDs, nil
				},
			},
		},
		"devices not found": {
			wantCode: http.StatusNotFound,
			s: mock.DeviceService{
				FindByStateFunc: func(state string) (device.Devices, error) {
					return nil, gorm.ErrRecordNotFound
				},
			},
		},
		"service returns error": {
			wantCode: http.StatusInternalServerError,
			s: mock.DeviceService{
				FindByStateFunc: func(state string) (device.Devices, error) {
					return nil, fmt.Errorf("boom")
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			handler := httpjson.NewHandler(&tc.s, nil)
			resp := test.DoHttpRequest(
				handler,
				http.MethodGet,
				"/devices/state/"+device.StateAvailable,
				nil,
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Fatalf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}

		})
	}
}

func TestHandlerDeleteDevice(t *testing.T) {
	var testCases = map[string]struct {
		wantCode int
		s        mock.DeviceService
	}{
		"successfully deletes device": {
			wantCode: http.StatusOK,
			s: mock.DeviceService{
				DeleteDeviceFunc: func(id uuid.UUID) error {
					return nil
				},
			},
		},
		"device is in use error": {
			wantCode: http.StatusUnprocessableEntity,
			s: mock.DeviceService{
				DeleteDeviceFunc: func(id uuid.UUID) error {
					return device.ErrDeviceInUse
				},
			},
		},
		"service returns error": {
			wantCode: http.StatusInternalServerError,
			s: mock.DeviceService{
				DeleteDeviceFunc: func(id uuid.UUID) error {
					return fmt.Errorf("boom")
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			handler := httpjson.NewHandler(&tc.s, nil)
			resp := test.DoHttpRequest(
				handler,
				http.MethodDelete,
				"/devices/"+uuid.New().String(),
				nil,
			)

			gotCode := resp.StatusCode

			if tc.wantCode != gotCode {
				t.Fatalf("expected status code %d, got: %d", tc.wantCode, gotCode)
			}
		})
	}
}
