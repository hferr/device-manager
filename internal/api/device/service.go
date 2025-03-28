package device

type DeviceService interface {
	CreateDevice(input CreateDeviceRequest) (*Device, error)
	ListDevices() (Devices, error)
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
