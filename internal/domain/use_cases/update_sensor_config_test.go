package use_cases_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain/use_cases"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSensorConfig_Success(t *testing.T) {
	mockRepo := &MockRepo{UpdateSensorConfigError: nil}
	app := use_cases.NewApp(mockRepo)

	err := app.UpdateSensorConfig(domain.Sensor{ID: uuid.New()})
	assert.NoError(t, err)
}

func TestUpdateSensorConfig_Error(t *testing.T) {
	mockRepo := &MockRepo{UpdateSensorConfigError: errors.New("update failed")}
	app := use_cases.NewApp(mockRepo)

	err := app.UpdateSensorConfig(domain.Sensor{ID: uuid.New()})
	assert.Error(t, err)
}
