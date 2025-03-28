package device

import "gorm.io/gorm"

type DeviceRepository interface {
	InsertDevice(device *Device) error
}

type deviceRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}

func (d *deviceRepository) InsertDevice(device *Device) error {
	if err := d.db.Create(device).Error; err != nil {
		return err
	}

	return nil
}
