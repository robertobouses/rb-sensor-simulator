package use_cases

import (
	"errors"
	"fmt"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

const (
	BelowMin = "Alert below min value"
	AboveMax = "Alert above max value"
)

func (a AppService) SaveSensorReading(reading *domain.SensorReading) error {
	if reading == nil {
		return fmt.Errorf("reading is nil")
	}

	comparableSensor, err := a.repo.GetSensorByID(reading.SensorID)
	if err != nil {
		return errors.New("failed to get sensor for reading: " + err.Error())
	}

	switch {
	case reading.Value < comparableSensor.AlertThresholds.Min:
		alert := BelowMin
		comparableSensor.Error = &alert
		comparableSensor.Status = domain.Warning
	case reading.Value > comparableSensor.AlertThresholds.Max:
		alert := AboveMax
		comparableSensor.Error = &alert
		comparableSensor.Status = domain.Warning
	default:
		comparableSensor.Error = nil
		comparableSensor.Status = domain.Active
	}

	if err := a.repo.SaveSensorReading(*reading); err != nil {
		return fmt.Errorf("failed to save sensor reading: %w", err)
	}

	if err := a.repo.UpdateSensorConfig(*comparableSensor); err != nil {
		return fmt.Errorf("failed to update sensor: %w", err)
	}

	return nil
}
