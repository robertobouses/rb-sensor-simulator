package repository

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) GetSensorConfigAndLastReadings(id uuid.UUID, numberOfReadings int) (*domain.Sensor, error) {
	row := r.getSensorByID.QueryRow(id)

	var sensor domain.Sensor
	err := row.Scan(
		&sensor.ID,
		&sensor.Name,
		&sensor.Type,
		&sensor.Config.SamplingInterval,
		&sensor.Config.AlertThreshold,
		&sensor.Config.Unit,
		&sensor.Config.Enabled,
	)
	if err != nil {
		return nil, err
	}

	rows, err := r.getSensorLastReadingsByID.Query(id, numberOfReadings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readings []domain.SensorReading
	for rows.Next() {
		var reading domain.SensorReading
		reading.SensorID = id
		err := rows.Scan(
			&reading.Timestamp,
			&reading.Value,
			&reading.Error)
		if err != nil {
			return nil, err
		}
		readings = append(readings, reading)
	}

	sensor.LastReading = &readings

	return &sensor, nil
}
