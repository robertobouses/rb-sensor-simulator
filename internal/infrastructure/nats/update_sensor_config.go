package nats

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go/micro"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (h Handler) UpdateSensorConfig(req micro.Request) {
	log.Printf("update sensor req: %v\n", req)

	var input CreateSensorRequest
	if err := json.Unmarshal(req.Data(), &input); err != nil {
		log.Printf("invalid sensor config payload: %v\n", err)
		_ = req.Respond([]byte(`{"error":"invalid request format"}`))
		return
	}
	sensorID, err := uuid.Parse(input.ID)
	if err != nil {
		fmt.Printf("error parsing %s to UUID\n", input.ID)
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
		Error:  input.Error,
		Status: domain.SensorStatus(input.Status),
	}

	if err := h.app.UpdateSensorConfig(sensor); err != nil {
		log.Printf("error updating sensor config: %v\n", err)
		_ = req.Respond([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}

	if err := req.Respond([]byte(`{"status":"ok"}`)); err != nil {
		log.Printf("error sending response: %v\n", err)
	}
}
