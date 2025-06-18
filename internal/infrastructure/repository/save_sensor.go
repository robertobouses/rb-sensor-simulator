package repository

import (
	"log"
	"time"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) SaveSensor(sensor domain.Sensor) error {
	var lastTimestamp *time.Time
	var lastValue *float64
	var lastError *string

	if sensor.LastReading != nil {
		lastTimestamp = &sensor.LastReading.Timestamp
		lastValue = &sensor.LastReading.Value
		lastError = sensor.LastReading.Error
	}

	_, err := r.saveSensor.Exec(
		sensor.Name,
		string(sensor.Type),
		int(sensor.Config.SamplingInterval.Seconds()),
		sensor.Config.AlertThreshold,
		sensor.Config.Unit,
		sensor.Config.Enabled,
		lastTimestamp,
		lastValue,
		lastError,
	)

	if err != nil {
		log.Print("Error executing SaveSensor statement:", err)
		return err
	}

	return nil
}
