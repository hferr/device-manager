package device

import (
	"time"

	"github.com/google/uuid"
)

type state uint8

const (
	StateAvailable state = iota
	StateInUse
	StateInactive
)

type Device struct {
	ID        uuid.UUID `gorm:"primarykey"`
	Name      string
	Brand     string
	State     state
	CreatedAt time.Time
}
