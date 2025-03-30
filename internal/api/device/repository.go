package device

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceRepository interface {
	InsertDevice(device *Device) error
	UpdateDevice(device *Device) error
	ListDevices() (Devices, error)
	FindByID(ID uuid.UUID) (*Device, error)
	FindByState(state string) (Devices, error)
	FindByBrand(brand string) (Devices, error)
	DeleteDevice(ID uuid.UUID) error
}

type deviceRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}

func (r *deviceRepository) InsertDevice(device *Device) error {
	if err := r.db.Create(device).Error; err != nil {
		return err
	}

	return nil
}

func (r *deviceRepository) UpdateDevice(device *Device) error {
	res := r.db.Model(&Device{}).
		Select("name", "brand", "state").
		Where("id = ?", device.ID).
		Updates(device)

	return res.Error
}

func (r *deviceRepository) ListDevices() (Devices, error) {
	ds := make(Devices, 0)
	if err := r.db.Find(&ds).Error; err != nil {
		return nil, err
	}

	return ds, nil
}

func (r *deviceRepository) FindByID(ID uuid.UUID) (*Device, error) {
	d := &Device{}
	if err := r.db.Where("id = ?", ID).First(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}

func (r *deviceRepository) FindByState(state string) (Devices, error) {
	ds := make(Devices, 0)
	if err := r.db.Where("state = ?", state).Find(&ds).Error; err != nil {
		return nil, err
	}

	return ds, nil
}

func (r *deviceRepository) FindByBrand(brand string) (Devices, error) {
	ds := make(Devices, 0)
	if err := r.db.Where("brand = ?", brand).Find(&ds).Error; err != nil {
		return nil, err
	}

	return ds, nil
}

func (r *deviceRepository) DeleteDevice(ID uuid.UUID) error {
	return r.db.Where("id = ?", ID).Delete(&Device{}).Error
}
