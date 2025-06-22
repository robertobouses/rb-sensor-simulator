package use_cases_test

import (
	"github.com/google/uuid"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

type MockRepo struct {
	SensorToReturn             *domain.Sensor
	SensorErrorToReturn        error
	SensorReadingErrorToReturn error
	SaveSensorError            error
	UpdateSensorConfigError    error
	SaveReadingErrorToReturn   error
	SaveSensorReadingError     error
	SaveAlertError             error
	UpdateAlertResolvedError   error
	GetSensorByIDError         error
}

func (m *MockRepo) SaveSensor(sensor *domain.Sensor) error {
	return m.SaveSensorError
}

func (m *MockRepo) SaveSensorReading(reading domain.SensorReading) error {
	return m.SaveSensorReadingError
}

func (m *MockRepo) UpdateSensorConfig(sensor domain.Sensor) error {
	return m.UpdateSensorConfigError
}

func (m *MockRepo) GetSensorConfigAndLastReadings(id uuid.UUID, numberOfReadings int) (*domain.Sensor, error) {
	return m.SensorToReturn, m.SensorErrorToReturn
}

func (m *MockRepo) GetSensorByID(id uuid.UUID) (*domain.Sensor, error) {
	return m.SensorToReturn, m.SensorErrorToReturn
}

func (m *MockRepo) SaveAlert(alert domain.AlertHistorial) error {
	return m.SaveAlertError
}

func (m *MockRepo) UpdateAlertResolved(sensorID, resolvedReadingID uuid.UUID) error {
	return m.UpdateAlertResolvedError
}
