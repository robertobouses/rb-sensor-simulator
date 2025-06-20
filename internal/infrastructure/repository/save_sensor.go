package repository

import (
	"log"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) SaveSensor(sensor *domain.Sensor) error {

	err := r.saveSensor.QueryRow(
		sensor.Name,
		string(sensor.Type),
		int(sensor.Config.SamplingInterval.Seconds()),
		sensor.Config.AlertThreshold,
		sensor.Config.Unit,
		sensor.Config.Enabled,
	).Scan(&sensor.ID)

	if err != nil {
		log.Print("Error executing SaveSensor statement:", err)
		return err
	}

	return nil
}
