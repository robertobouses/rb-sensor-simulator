package nats

import "time"

type EventSensorReading struct {
	SensorID  string    `json:"sensor_id"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
	Error     *string   `json:"error,omitempty"`
}
