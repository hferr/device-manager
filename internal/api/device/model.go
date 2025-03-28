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
