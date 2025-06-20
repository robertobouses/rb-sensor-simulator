package use_cases

import "github.com/robertobouses/rb-sensor-simulator/internal/domain"

func (a *AppService) SaveSensorReading(reading domain.SensorReading) error {
	err := a.repo.SaveSensorReading(reading)
	if err != nil {
		return err
	}

	err = a.repo.UpdateSensorLastReading(reading)
	if err != nil {
		return err
	}
	return nil
}
