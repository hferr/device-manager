package mock

import (
	"github/hferr/device-manager/internal/api/device"

	"github.com/google/uuid"
)

type DeviceRepository struct {
	InsertDeviceFunc func(d *device.Device) error
	UpdateDeviceFunc func(d *device.Device) error
	ListDevicesFunc  func() (device.Devices, error)
	FindByIDFunc     func(ID uuid.UUID) (*device.Device, error)
	FindByStateFunc  func(state string) (device.Devices, error)
	FindByBrandFunc  func(brand string) (device.Devices, error)
	DeleteDeviceFunc func(ID uuid.UUID) error
}

func (r *DeviceRepository) InsertDevice(d *device.Device) error {
	return r.InsertDeviceFunc(d)
}

func (r *DeviceRepository) UpdateDevice(d *device.Device) error {
	return r.UpdateDeviceFunc(d)
}

func (r *DeviceRepository) ListDevices() (device.Devices, error) {
	return r.ListDevicesFunc()
}

func (r *DeviceRepository) FindByID(ID uuid.UUID) (*device.Device, error) {
	return r.FindByIDFunc(ID)
}

func (r *DeviceRepository) FindByState(state string) (device.Devices, error) {
	return r.FindByStateFunc(state)
}

func (r *DeviceRepository) FindByBrand(brand string) (device.Devices, error) {
	return r.FindByBrandFunc(brand)
}

func (r *DeviceRepository) DeleteDevice(ID uuid.UUID) error {
	return r.DeleteDeviceFunc(ID)
}
