package device

import "github.com/google/uuid"

type DeviceService interface {
	CreateDevice(input CreateDeviceRequest) (*Device, error)
	ListDevices() (Devices, error)
	FindByID(ID uuid.UUID) (*Device, error)
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
