package use_cases

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (a AppService) GetSensorConfigAndLastReadings(sensorID uuid.UUID, numberOfReadings int) (*domain.Sensor, error) {
	sensor, err := a.repo.GetSensorConfigAndLastReadings(sensorID, numberOfReadings)
	if err != nil {
		return nil, err
	}

	return &domain.Sensor{
		ID:               sensor.ID,
		Name:             sensor.Name,
		Type:             sensor.Type,
		SamplingInterval: sensor.SamplingInterval,
		AlertThresholds:  sensor.AlertThresholds,
		Unit:             sensor.Unit,
		Status:           sensor.Status,
		LastReading:      sensor.LastReading,
	}, nil
}
