package use_cases_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain/use_cases"
	"github.com/stretchr/testify/assert"
)

func TestSaveSensor_Success(t *testing.T) {
	mockRepo := &MockRepo{SaveSensorError: nil}
	app := use_cases.NewApp(mockRepo)

	err := app.CreateSensor(&domain.Sensor{ID: uuid.New()})
	assert.NoError(t, err)
}

func TestSaveSensor_Error(t *testing.T) {
	mockRepo := &MockRepo{SaveSensorError: errors.New("db error")}
	app := use_cases.NewApp(mockRepo)

	err := app.CreateSensor(&domain.Sensor{ID: uuid.New()})
	assert.Error(t, err)
}
