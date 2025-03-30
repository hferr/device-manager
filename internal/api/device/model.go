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

type Devices []*Device

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

type UpdateDeviceRequest struct {
	Name  *string `json:"name"`
	Brand *string `json:"brand"`
	State *string `json:"state" validate:"required,oneof=available in_use inactive"`
}

func (r *UpdateDeviceRequest) Apply(d *Device) {
	if r.Name != nil {
		d.Name = *r.Name
	}

	if r.Brand != nil {
		d.Brand = *r.Brand
	}

	if r.State != nil {
		d.State = *r.State
	}
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

func (d *Device) ToDto() *DTO {
	return &DTO{
		ID:        d.ID,
		Name:      d.Name,
		Brand:     d.Brand,
		State:     d.State,
		CreatedAt: d.CreatedAt.Format(time.DateTime),
	}
}

func (ds Devices) ToDto() []*DTO {
	dtos := make([]*DTO, len(ds))
	for i, v := range ds {
		dtos[i] = v.ToDto()
	}

	return dtos
}
