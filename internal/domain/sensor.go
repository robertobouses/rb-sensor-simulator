package domain

import "github.com/google/uuid"

type SensorType string

const (
	TemperatureSensor SensorType = "temperature"
	HumiditySensor    SensorType = "humidity"
	PressureSensor    SensorType = "pressure"
)

type Sensor struct {
	ID          uuid.UUID
	Name        string
	Type        SensorType
	Config      SensorConfig
	LastReading *SensorReading
}
