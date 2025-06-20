package use_cases

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (a *AppService) UpdateSensorConfig(sensorID uuid.UUID, config domain.SensorConfig) error {
	return a.repo.UpdateSensorConfig(sensorID, config)
}
