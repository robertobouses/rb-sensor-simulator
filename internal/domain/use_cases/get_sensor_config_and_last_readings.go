package use_cases

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (a *AppService) GetSensorConfigAndLastReadings(sensorID uuid.UUID) (*domain.SensorStatus, error) {
	sensor, err := a.repo.GetSensorByID(sensorID)
	if err != nil {
		return nil, err
	}

	return &domain.SensorStatus{
		SensorID:         sensor.ID,
		SamplingInterval: sensor.Config.SamplingInterval,
		AlertThreshold:   sensor.Config.AlertThreshold,
		Unit:             sensor.Config.Unit,
		Enabled:          sensor.Config.Enabled,
		LastReading:      sensor.LastReading,
	}, nil
}
