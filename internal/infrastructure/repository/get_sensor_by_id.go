package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) GetSensorByID(id uuid.UUID) (*domain.Sensor, error) {
	row := r.getSensorByID.QueryRow(id)

	var sensor domain.Sensor
	var samplingInterval int
	err := row.Scan(
		&sensor.ID,
		&sensor.Name,
		&sensor.Type,
		&samplingInterval,
		&sensor.AlertThresholds.Min,
		&sensor.AlertThresholds.Max,
		&sensor.Unit,
		&sensor.Status,
	)
	if err != nil {
		return nil, err
	}

	sensor.SamplingInterval = time.Duration(samplingInterval) * time.Millisecond
	return &sensor, nil
}
