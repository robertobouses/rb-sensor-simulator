package use_cases

import (
	"errors"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (a AppService) SaveSensorReading(reading domain.SensorReading) error {

	comparableSensor, err := a.repo.GetSensorByID(reading.SensorID)

	if err != nil {
		return errors.New("failed to get sensor for reading: " + err.Error())
	}

	if reading.Value < comparableSensor.AlertThresholds.Min {
		alert := "Alert Value Min"
		reading.Error = &alert
	} else if reading.Value > comparableSensor.AlertThresholds.Max {
		alert := "Alert Value Max"
		reading.Error = &alert
	}

	return a.repo.SaveSensorReading(reading)
}
