package use_cases

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

type Repository interface {
	SaveSensor(sensor *domain.Sensor) error
	SaveSensorReading(sensorID uuid.UUID, reading domain.SensorReading) error
	UpdateSensorLastReading(sensorID uuid.UUID, reading domain.SensorReading) error
}

func NewApp(repository Repository) AppService {
	return AppService{
		repo: repository,
	}
}

type AppService struct {
	repo Repository
}
