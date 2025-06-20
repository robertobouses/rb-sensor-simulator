package repository

import (
	"log"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) UpdateSensorLastReading(reading domain.SensorReading) error {
	var errMsg *string
	if reading.Error != nil {
		errMsg = reading.Error
	}

	_, err := r.updateSensorLastReading.Exec(
		reading.SensorID,
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
