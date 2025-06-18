package simulator

import (
	"math/rand"
	"time"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

type SensorSimulator struct {
	Sensors  []domain.Sensor
	OnRead   func(sensor domain.Sensor, reading domain.SensorReading)
	stopChan chan struct{}
}

func NewSensorSimulator(sensors []domain.Sensor, onRead func(sensor domain.Sensor, reading domain.SensorReading)) *SensorSimulator {
	return &SensorSimulator{
		Sensors:  sensors,
		OnRead:   onRead,
		stopChan: make(chan struct{}),
	}
}

func (s *SensorSimulator) Start() {
	for _, sensor := range s.Sensors {
		go s.simulateSensor(sensor)
	}
}

func (s *SensorSimulator) Stop() {
	close(s.stopChan)
}

func (s *SensorSimulator) simulateSensor(sensor domain.Sensor) {
	ticker := time.NewTicker(sensor.Config.SamplingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopChan:
			return
		case <-ticker.C:
			reading := generateReading(sensor.Type)
			s.OnRead(sensor, reading)
		}
	}
}

func generateReading(sensorType domain.SensorType) domain.SensorReading {
	timestamp := time.Now()
	var value float64
	var errStr *string

	switch sensorType {
	case domain.TemperatureSensor:
		value = 20 + rand.Float64()*10
	case domain.HumiditySensor:
		value = 40 + rand.Float64()*30
	case domain.PressureSensor:
		value = 950 + rand.Float64()*50
	default:
		msg := "unknown sensor type"
		errStr = &msg
	}

	if rand.Float64() < 0.1 {
		msg := "sensor read error"
		errStr = &msg
		value = 0
	}

	return domain.SensorReading{
		Timestamp: timestamp,
		Value:     value,
		Error:     errStr,
	}
}
