package repository

import (
	"log"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) UpdateSensorConfig(sensor domain.Sensor) error {
	_, err := r.updateSensorConfig.Exec(
		int(sensor.SamplingInterval.Milliseconds()),
		sensor.AlertThresholds.Min,
		sensor.AlertThresholds.Max,
		sensor.Unit,
		sensor.Error,
		sensor.Status,
		sensor.ID,
	)

	if err != nil {
		log.Printf("Error updating sensor config: %v", err)
		return err
	}

	return nil
}
