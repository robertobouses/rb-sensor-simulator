package use_cases

import "github.com/robertobouses/rb-sensor-simulator/internal/domain"

func (a AppService) CreateSensor(sensor *domain.Sensor) error {
	sensor.Status = domain.Active
	return a.repo.SaveSensor(sensor)
}
