package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) UpdateSensorConfig(sensorID uuid.UUID, config domain.SensorConfig) error {

	_, err := r.updateSensorConfig.Exec(
		int(config.SamplingInterval.Milliseconds()),
		config.AlertThreshold,
		config.Unit,
		config.Enabled,
		sensorID,
	)

	if err != nil {
		log.Printf("Error updating sensor config: %v", err)
		return err
	}

	return nil
}
