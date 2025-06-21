package domain

import (
	"time"

	"github.com/google/uuid"
)

type SensorReading struct {
	SensorID  uuid.UUID
	Timestamp time.Time
	Value     float64
}
