package repository

import (
	"fmt"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) UpdateSensorConfig(sensor domain.Sensor) error {

	current, err := r.GetSensorByID(sensor.ID)
	if err != nil {
		return fmt.Errorf("error fetching current sensor: %w", err)
	}

	if sensor.SamplingInterval == 0 {
		sensor.SamplingInterval = current.SamplingInterval
	}
	if sensor.AlertThresholds.Min == 0 {
		sensor.AlertThresholds.Min = current.AlertThresholds.Min
	}
	if sensor.AlertThresholds.Max == 0 {
		sensor.AlertThresholds.Max = current.AlertThresholds.Max
	}
	if sensor.Unit == "" {
		sensor.Unit = current.Unit
	}

	if sensor.Status == "" {
		sensor.Status = current.Status
	}

	_, err = r.updateSensorConfig.Exec(
		int(sensor.SamplingInterval.Milliseconds()),
		sensor.AlertThresholds.Min,
		sensor.AlertThresholds.Max,
		sensor.Unit,
		sensor.Status,
		sensor.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating sensor config on UpdateSensorConfig: %w", err)
	}

	return nil
}
