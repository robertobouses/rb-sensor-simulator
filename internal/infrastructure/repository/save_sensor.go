package repository

import (
	"log"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) SaveSensor(sensor *domain.Sensor) error {

	_, err := r.saveSensor.Exec(
		sensor.Name,
		string(sensor.Type),
		int(sensor.SamplingInterval.Seconds()),
		sensor.AlertThresholds.Min,
		sensor.AlertThresholds.Max,
		sensor.Unit,
		sensor.Error,
		sensor.Status,
	)

	if err != nil {
		log.Print("Error executing SaveSensor statement:", err)
		return err
	}

	return nil
}
