package nats

import "github.com/robertobouses/rb-sensor-simulator/internal/domain"

type App interface {
	SaveSensor(sensor *domain.Sensor) error
	SaveSensorReading(reading domain.SensorReading) error
	UpdateSensorConfig(sensor domain.Sensor) error
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
