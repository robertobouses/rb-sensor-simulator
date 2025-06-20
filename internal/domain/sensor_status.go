package domain

import (
	"time"

	"github.com/google/uuid"
)

type SensorStatus struct {
	SensorID         uuid.UUID        `json:"sensor_id"`
	SamplingInterval time.Duration    `json:"sampling_interval"`
	AlertThreshold   float64          `json:"alert_threshold"`
	Unit             string           `json:"unit"`
	Enabled          bool             `json:"enabled"`
	LastReading      *[]SensorReading `json:"last_reading,omitempty"`
}
