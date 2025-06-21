package nats

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

type EventSensorReading struct {
	SensorID  string    `json:"sensor_id"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
	Error     *string   `json:"error,omitempty"`
}

func (h Handler) ProcessSensorReading(msg *nats.Msg) {
	var eventSensorReading EventSensorReading
	err := json.Unmarshal(msg.Data, &eventSensorReading)
	if err != nil {
		fmt.Println("Error in ProcessSensorReading")
		return
	}
	sensorID, err := uuid.Parse(eventSensorReading.SensorID)
	if err != nil {
		fmt.Printf("error parsing %s to UUID\n", eventSensorReading.SensorID)
	}
	sensorReading := domain.SensorReading{
		SensorID:  sensorID,
		Timestamp: eventSensorReading.Timestamp,
		Value:     eventSensorReading.Value,
	}

	sensor := domain.Sensor{
		Error: eventSensorReading.Error,
	}
	if err := h.app.SaveSensorReading(&sensorReading); err != nil {
		fmt.Printf("error saving sensor reading: %v\n", err)
	}
	if err := h.app.UpdateSensorConfig(sensor); err != nil {
		fmt.Printf("error updating sensor config: %v\n", err)
	}
}
