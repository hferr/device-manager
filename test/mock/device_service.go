package mock

import (
	"github/hferr/device-manager/internal/api/device"

	"github.com/google/uuid"
)

type DeviceService struct {
	CreateDeviceFunc func(input device.CreateDeviceRequest) (*device.Device, error)
	UpdateDeviceFunc func(ID uuid.UUID, input device.UpdateDeviceRequest) error
	ListDevicesFunc  func() (device.Devices, error)
	FindByIDFunc     func(ID uuid.UUID) (*device.Device, error)
	FindByStateFunc  func(state string) (device.Devices, error)
	FindByBrandFunc  func(brand string) (device.Devices, error)
	DeleteDeviceFunc func(ID uuid.UUID) error
}

func (ds *DeviceService) CreateDevice(input device.CreateDeviceRequest) (*device.Device, error) {
	return ds.CreateDeviceFunc(input)
}

func (ds *DeviceService) UpdateDevice(ID uuid.UUID, input device.UpdateDeviceRequest) error {
	return ds.UpdateDeviceFunc(ID, input)
}

func (ds *DeviceService) ListDevices() (device.Devices, error) {
	return ds.ListDevicesFunc()
}

func (ds *DeviceService) FindByID(ID uuid.UUID) (*device.Device, error) {
	return ds.FindByIDFunc(ID)
}

func (ds *DeviceService) FindByState(state string) (device.Devices, error) {
	return ds.FindByStateFunc(state)
}

func (ds *DeviceService) FindByBrand(brand string) (device.Devices, error) {
	return ds.FindByBrandFunc(brand)
}

func (ds *DeviceService) DeleteDevice(ID uuid.UUID) error {
	return ds.DeleteDeviceFunc(ID)
}
