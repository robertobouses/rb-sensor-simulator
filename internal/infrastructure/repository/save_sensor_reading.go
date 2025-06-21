package repository

import (
	"log"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) SaveSensorReading(reading domain.SensorReading) error {
	_, err := r.saveSensorReading.Exec(
		reading.SensorID,
		reading.Timestamp,
		reading.Value,
	)

	if err != nil {
		log.Print("Error executing SaveSensorReading statement:", err)
		return err
	}

	return nil
}
