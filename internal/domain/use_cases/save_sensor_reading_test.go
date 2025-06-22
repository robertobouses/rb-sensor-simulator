package use_cases_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain/use_cases"
	"github.com/stretchr/testify/assert"
)

func TestSaveSensorReading_Success_NoAlert(t *testing.T) {
	id := uuid.New()
	mockSensor := &domain.Sensor{
		ID:              id,
		AlertThresholds: domain.Threshold{Min: 10, Max: 100},
	}
	mockRepo := &MockRepo{
		SensorToReturn:      mockSensor,
		SensorErrorToReturn: nil,
	}

	app := use_cases.NewApp(mockRepo)

	reading := &domain.SensorReading{SensorID: id, Value: 50}
	err := app.SaveSensorReading(reading)
	assert.NoError(t, err)
}

func TestSaveSensorReading_AlertMin(t *testing.T) {
	id := uuid.New()
	mockSensor := &domain.Sensor{
		ID:              id,
		AlertThresholds: domain.Threshold{Min: 10, Max: 100},
	}
	mockRepo := &MockRepo{SensorToReturn: mockSensor}
	app := use_cases.NewApp(mockRepo)

	reading := domain.SensorReading{SensorID: id, Value: 5}
	err := app.SaveSensorReading(&reading)

	assert.NoError(t, err)
}

func TestSaveSensorReading_AlertMax(t *testing.T) {
	id := uuid.New()
	mockSensor := &domain.Sensor{
		ID:              id,
		AlertThresholds: domain.Threshold{Min: 10, Max: 100},
	}
	mockRepo := &MockRepo{SensorToReturn: mockSensor}

	app := use_cases.NewApp(mockRepo)

	reading := domain.SensorReading{SensorID: id, Value: 150}
	err := app.SaveSensorReading(&reading)
	assert.NoError(t, err)
}

func TestSaveSensorReading_GetSensorError(t *testing.T) {
	mockRepo := &MockRepo{SensorErrorToReturn: errors.New("not found")}
	app := use_cases.NewApp(mockRepo)

	reading := domain.SensorReading{SensorID: uuid.New(), Value: 50}
	err := app.SaveSensorReading(&reading)
	assert.Error(t, err)
}

func TestSaveSensorReading_NilReading(t *testing.T) {
	mockRepo := &MockRepo{}
	app := use_cases.NewApp(mockRepo)

	err := app.SaveSensorReading(nil)
	assert.Error(t, err)
	assert.EqualError(t, err, "reading is nil")
}

func TestSaveSensorReading_SaveReadingError(t *testing.T) {
	id := uuid.New()
	mockSensor := &domain.Sensor{
		ID:              id,
		AlertThresholds: domain.Threshold{Min: 10, Max: 100},
	}
	mockRepo := &MockRepo{
		SensorToReturn:         mockSensor,
		SaveSensorReadingError: errors.New("db error"),
	}
	app := use_cases.NewApp(mockRepo)

	reading := &domain.SensorReading{SensorID: id, Value: 50}
	err := app.SaveSensorReading(reading)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save sensor reading")
}
func TestSaveSensorReading_UpdateSensorError(t *testing.T) {
	id := uuid.New()
	mockSensor := &domain.Sensor{
		ID:              id,
		AlertThresholds: domain.Threshold{Min: 10, Max: 100},
		Status:          domain.Active,
	}
	mockRepo := &MockRepo{
		SensorToReturn:          mockSensor,
		UpdateSensorConfigError: errors.New("update failed")}
	app := use_cases.NewApp(mockRepo)

	reading := &domain.SensorReading{SensorID: id, Value: 110}
	err := app.SaveSensorReading(reading)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to update sensor")
}
