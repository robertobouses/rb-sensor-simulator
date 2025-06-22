package domain

import (
	"time"

	"github.com/google/uuid"
)

type SensorReading struct {
	ID        uuid.UUID
	SensorID  uuid.UUID
	Timestamp time.Time
	Value     float64
	Error     *string
}
