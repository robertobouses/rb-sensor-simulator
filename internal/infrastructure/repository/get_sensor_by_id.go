package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) GetSensorByID(id uuid.UUID) (*domain.Sensor, error) {
	row := r.getSensorByID.QueryRow(id)

	var sensor domain.Sensor
	var lastReadingTimestamp *time.Time
	var lastReadingValue *float64
	var lastReadingError *string

	err := row.Scan(
		&sensor.ID,
		&sensor.Name,
		&sensor.Type,
		&sensor.Config.SamplingInterval,
		&sensor.Config.AlertThreshold,
		&sensor.Config.Unit,
		&sensor.Config.Enabled,
		&lastReadingTimestamp,
		&lastReadingValue,
		lastReadingError,
	)
	if err != nil {
		return nil, err
	}

	if lastReadingTimestamp != nil && lastReadingValue != nil {
		sensor.LastReading = &domain.SensorReading{
			Timestamp: *lastReadingTimestamp,
			Value:     *lastReadingValue,
			Error:     lastReadingError,
		}
	}

	return &sensor, nil
}
