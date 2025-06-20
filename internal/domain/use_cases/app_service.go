package use_cases

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

type Repository interface {
	SaveSensor(sensor *domain.Sensor) error
	SaveSensorReading(reading domain.SensorReading) error
	UpdateSensorLastReading(reading domain.SensorReading) error
	UpdateSensorConfig(sensorID uuid.UUID, config domain.SensorConfig) error
	GetSensorByID(id uuid.UUID) (*domain.Sensor, error)
}

func NewApp(repository Repository) AppService {
	return AppService{
		repo: repository,
	}
}

type AppService struct {
	repo Repository
}
