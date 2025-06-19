package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) SaveSensorReading(sensorID uuid.UUID, reading domain.SensorReading) error {
	_, err := r.saveSensorReading.Exec(
		sensorID,
		reading.Timestamp,
		reading.Value,
		reading.Error,
	)

	if err != nil {
		log.Print("Error executing SaveSensorReading statement:", err)
		return err
	}

	return nil
}
