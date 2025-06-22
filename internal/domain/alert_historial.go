package domain

import (
	"time"

	"github.com/google/uuid"
)

type AlertType string

const (
	AlertBelowMin AlertType = "below_min"
	AlertAboveMax AlertType = "above_max"
)

type AlertHistorial struct {
	ID                uuid.UUID
	SensorID          uuid.UUID
	AlertType         AlertType
	WarningReadingID  uuid.UUID
	ResolvedReadingID uuid.UUID
	CreatedAt         time.Time
}
