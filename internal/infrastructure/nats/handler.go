package nats

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

type App interface {
	CreateSensor(sensor *domain.Sensor) error
	SaveSensorReading(reading *domain.SensorReading) error
	UpdateSensorConfig(sensor domain.Sensor) error
	GetSensorConfigAndLastReadings(sensorID uuid.UUID, numberOfReadings int) (*domain.Sensor, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
