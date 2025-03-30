package device

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	StateAvailable string = "available"
	StateInUse     string = "in_use"
	StateInactive  string = "inactive"
)

type DeviceService interface {
	CreateDevice(input CreateDeviceRequest) (*Device, error)
	UpdateDevice(ID uuid.UUID, input UpdateDeviceRequest) error
	ListDevices() (Devices, error)
	FindByID(ID uuid.UUID) (*Device, error)
	FindByState(state string) (Devices, error)
	FindByBrand(brand string) (Devices, error)
	DeleteDevice(ID uuid.UUID) error
}

type deviceService struct {
	repo DeviceRepository
}

func NewService(r DeviceRepository) DeviceService {
	return &deviceService{
		repo: r,
	}
}

func (s *deviceService) CreateDevice(input CreateDeviceRequest) (*Device, error) {
	d := NewDevice(input.Name, input.Brand, input.State)

	if err := s.repo.InsertDevice(d); err != nil {
		return d, err
	}

	return d, nil
}

func (s *deviceService) UpdateDevice(ID uuid.UUID, input UpdateDeviceRequest) error {
	d, err := s.FindByID(ID)
	if err != nil {
		return err
	}

	if !canUpdateDevice(d, input) {
		return fmt.Errorf("device can not be updated in it's current state")
	}

	input.Apply(d)

	if err := s.repo.UpdateDevice(d); err != nil {
		return err
	}

	return nil
}

func (s *deviceService) ListDevices() (Devices, error) {
	ds, err := s.repo.ListDevices()
	if err != nil {
		return nil, err
	}

	return ds, nil
}

func (s *deviceService) FindByID(ID uuid.UUID) (*Device, error) {
	d, err := s.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (s *deviceService) FindByState(state string) (Devices, error) {
	ds, err := s.repo.FindByState(state)
	if err != nil {
		return nil, err
	}

	return ds, nil
}

func (s *deviceService) FindByBrand(brand string) (Devices, error) {
	ds, err := s.repo.FindByBrand(brand)
	if err != nil {
		return nil, err
	}

	return ds, nil
}

func (s *deviceService) DeleteDevice(ID uuid.UUID) error {
	d, err := s.FindByID(ID)
	if err != nil {
		return err
	}

	if isDeviceInUse(d) {
		return fmt.Errorf("device is in-use and cannot be deleted")
	}

	return s.repo.DeleteDevice(ID)
}

func isDeviceInUse(d *Device) bool {
	return d.State == StateInUse
}

func canUpdateDevice(d *Device, input UpdateDeviceRequest) bool {
	// device name or brand cannot be updated if device state is 'in-use'
	if (d.State == StateInUse) &&
		(input.Name != nil || input.Brand != nil) {
		return false
	}

	return true
}
