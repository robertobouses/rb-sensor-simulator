package use_cases

import "github.com/robertobouses/rb-sensor-simulator/internal/domain"

func (a *AppService) SaveSensorReading(sensor domain.Sensor, reading domain.SensorReading) error {
	err := a.repo.SaveSensorReading(sensor.ID, reading)
	if err != nil {
		return err
	}

	return nil
}
