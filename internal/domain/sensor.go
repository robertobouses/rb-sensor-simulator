package domain

import (
	"time"

	"github.com/google/uuid"
)

type SensorType string

const (
	TemperatureSensor SensorType = "temperature"
	HumiditySensor    SensorType = "humidity"
	PressureSensor    SensorType = "pressure"
)

type Sensor struct {
	ID               uuid.UUID
	Name             string
	Type             SensorType
	SamplingInterval time.Duration
	AlertThresholds  Threshold
	Unit             string
	LastReading      *[]SensorReading
}

type Threshold struct {
	Min float64
	Max float64
}
