package repository

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) GetSensorByID(id uuid.UUID) (*domain.Sensor, error) {
	row := r.getSensorByID.QueryRow(id)

	var sensor domain.Sensor
	err := row.Scan(
		&sensor.ID,
		&sensor.Name,
		&sensor.Type,
		&sensor.SamplingInterval,
		&sensor.AlertThresholds.Min,
		&sensor.AlertThresholds.Max,
		&sensor.Unit,
		&sensor.Error,
		&sensor.Status,
	)
	if err != nil {
		return nil, err
	}
	return &sensor, nil
}
