package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) UpdateSensorLastReading(sensorID uuid.UUID, reading domain.SensorReading) error {
	var errMsg *string
	if reading.Error != nil {
		errMsg = reading.Error
	}

	_, err := r.updateSensorLastReading.Exec(
		sensorID,
		reading.Timestamp,
		reading.Value,
		errMsg,
	)

	if err != nil {
		log.Printf("Error updating sensor last reading: %v", err)
		return err
	}

	return nil
}
