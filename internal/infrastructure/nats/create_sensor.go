package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go/micro"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

type CreateSensorRequest struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Type             string        `json:"type"`
	SamplingInterval time.Duration `json:"sampling_interval"`
	AlertThresholds  Threshold     `json:"alert_thresholds"`
	Unit             string        `json:"unit"`
	Status           string        `json:"status"`
}
type Threshold struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func (h Handler) CreateSensor(req micro.Request) {
	log.Printf("create sensor req: %v\n", req)

	var input CreateSensorRequest
	if err := json.Unmarshal(req.Data(), &input); err != nil {
		log.Printf("invalid sensor config payload: %v\n", err)
		_ = req.Respond([]byte(`{"error":"invalid request format"}`))
		return
	}
	sensorID, err := uuid.Parse(input.ID)
	if err != nil {
		log.Printf("invalid sensor id: %v\n", err)
		_ = req.Respond([]byte(`{"error":"invalid request format"}`))
		return
	}
	sensor := domain.Sensor{
		ID:               sensorID,
		Name:             input.Name,
		Type:             domain.SensorType(input.Type),
		SamplingInterval: input.SamplingInterval,
		AlertThresholds: domain.Threshold{
			Min: input.AlertThresholds.Min,
			Max: input.AlertThresholds.Max,
		},
		Unit:   input.Unit,
		Status: domain.SensorStatus(input.Status),
	}

	if err := h.app.CreateSensor(&sensor); err != nil {
		log.Printf("error saving sensor: %v\n", err)
		_ = req.Respond([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}

	if err := req.Respond([]byte(`{"status":"created"}`)); err != nil {
		log.Printf("error responding create sensor request: %v\n", err)
	}
}
