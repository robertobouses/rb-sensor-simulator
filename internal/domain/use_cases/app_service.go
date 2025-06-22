package use_cases

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

type Repository interface {
	SaveSensor(sensor *domain.Sensor) error
	SaveSensorReading(reading domain.SensorReading) error
	UpdateSensorConfig(sensor domain.Sensor) error
	GetSensorConfigAndLastReadings(id uuid.UUID, numberOfReadings int) (*domain.Sensor, error)
	GetSensorByID(id uuid.UUID) (*domain.Sensor, error)
	SaveAlert(alert domain.AlertHistorial) error
	UpdateAlertResolved(sensorID, resolvedReadingID uuid.UUID) error
}

func NewApp(repository Repository) AppService {
	return AppService{
		repo: repository,
	}
}

type AppService struct {
	repo Repository
}
