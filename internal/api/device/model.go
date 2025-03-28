package device

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID        uuid.UUID `gorm:"primarykey"`
	Name      string
	Brand     string
	State     string
	CreatedAt time.Time
}

type DTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	State     string    `json:"state"`
	CreatedAt string    `json:"created_at"`
}

type CreateDeviceRequest struct {
	Name  string `json:"name" validate:"required,max=255"`
	Brand string `json:"brand" validate:"required,max=255"`
	State string `json:"state" validate:"required,oneof=available in_use inactive"`
}

func NewDevice(name, brand, state string) *Device {
	return &Device{
		ID:        uuid.New(),
		Name:      name,
		Brand:     brand,
		State:     state,
		CreatedAt: time.Now(),
	}
}

func (d *Device) ToDTO() *DTO {
	return &DTO{
		ID:        d.ID,
		Name:      d.Name,
		Brand:     d.Brand,
		State:     d.State,
		CreatedAt: d.CreatedAt.Format(time.DateTime),
	}
}
