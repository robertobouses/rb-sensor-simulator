package use_cases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (a AppService) SaveSensorReading(reading *domain.SensorReading) error {
	if reading == nil {
		return fmt.Errorf("reading is nil")
	}

	reading.ID = uuid.New()

	sensor, err := a.repo.GetSensorByID(reading.SensorID)
	if err != nil {
		return fmt.Errorf("failed to get sensor for reading: %w", err)
	}

	var alert *domain.AlertHistorial

	switch {
	case reading.Value < sensor.AlertThresholds.Min:
		alert = &domain.AlertHistorial{
			SensorID:         sensor.ID,
			AlertType:        domain.AlertBelowMin,
			WarningReadingID: reading.ID,
		}

	case reading.Value > sensor.AlertThresholds.Max:
		alert = &domain.AlertHistorial{
			SensorID:         sensor.ID,
			AlertType:        domain.AlertAboveMax,
			WarningReadingID: reading.ID,
		}
	}

	if err := a.repo.SaveSensorReading(*reading); err != nil {
		return fmt.Errorf("failed to save sensor reading: %w", err)
	}

	if reading.Error != nil {
		return nil
	}

	if alert != nil && sensor.Status != domain.Warning {
		sensor.Status = domain.Warning
		if err := a.repo.SaveAlert(*alert); err != nil {
			return fmt.Errorf("failed to save alert: %w", err)
		}
		if err := a.repo.UpdateSensorConfig(*sensor); err != nil {
			return fmt.Errorf("failed to update sensor: %w", err)
		}
	}

	if alert == nil && sensor.Status == domain.Warning {
		sensor.Status = domain.Active
		if err := a.repo.UpdateAlertResolved(sensor.ID, reading.ID); err != nil {
			if err.Error() != "no open alert found" {
				return fmt.Errorf("failed to resolve alert: %w", err)
			}

		}
		if err := a.repo.UpdateSensorConfig(*sensor); err != nil {
			return fmt.Errorf("failed to update sensor: %w", err)
		}
	}

	return nil
}
