package device

type DeviceService interface{}

type deviceService struct {
	repo DeviceRepository
}

func NewService(r DeviceRepository) DeviceService {
	return &deviceService{
		repo: r,
	}
}
