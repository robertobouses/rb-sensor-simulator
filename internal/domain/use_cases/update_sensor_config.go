package use_cases

import (
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (a *AppService) UpdateSensorConfig(sensor domain.Sensor) error {
	return a.repo.UpdateSensorConfig(sensor)
}
