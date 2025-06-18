package domain

import "time"

type SensorReading struct {
	Timestamp time.Time
	Value     float64
	Error     *string
}
