package device

import "gorm.io/gorm"

type DeviceRepository interface{}

type deviceRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}
