package use_cases_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain/use_cases"
	"github.com/stretchr/testify/assert"
)

func TestGetSensorConfigAndLastReadings_Success(t *testing.T) {
	id := uuid.New()
	readings := []domain.SensorReading{
		{SensorID: id, Value: 23.5, Timestamp: time.Now()},
	}
	mockSensor := &domain.Sensor{
		ID:               id,
		Name:             "Sensor1",
		Type:             domain.TemperatureSensor,
		SamplingInterval: time.Second * 5,
		AlertThresholds:  domain.Threshold{Min: 0, Max: 100},
		Unit:             "C",
		Error:            nil,
		Status:           domain.Active,
		LastReading:      &readings,
	}
	mockRepo := &MockRepo{SensorToReturn: mockSensor}
	app := use_cases.NewApp(mockRepo)

	result, err := app.GetSensorConfigAndLastReadings(id, 1)
	assert.NoError(t, err)
	assert.Equal(t, mockSensor, result)

}

func TestGetSensorConfigAndLastReadings_Error(t *testing.T) {
	mockRepo := &MockRepo{SensorErrorToReturn: errors.New("db error")}
	app := use_cases.NewApp(mockRepo)

	result, err := app.GetSensorConfigAndLastReadings(uuid.New(), 1)
	assert.Error(t, err)
	assert.Nil(t, result)
}
