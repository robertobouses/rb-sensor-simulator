package nats

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go/micro"
)

type GetSensorRequest struct {
	ID               string `json:"id"`
	NumberOfReadings int    `json:"number_of_readings"`
}

func (h Handler) GetSensorConfigAndLastReadings(req micro.Request) {
	log.Printf("get sensor config and readings req: %v\n", req)

	var input GetSensorRequest
	if err := json.Unmarshal(req.Data(), &input); err != nil {
		log.Printf("invalid get sensor request payload: %v\n", err)
		_ = req.Respond([]byte(`{"error":"invalid request format"}`))
		return
	}

	sensorID, err := uuid.Parse(input.ID)
	if err != nil {
		log.Printf("invalid sensor ID: %v\n", err)
		_ = req.Respond([]byte(`{"error":"invalid sensor ID format"}`))
		return
	}

	readings := input.NumberOfReadings
	if readings <= 0 {
		readings = 1
	}

	sensor, err := h.app.GetSensorConfigAndLastReadings(sensorID, readings)
	if err != nil {
		log.Printf("error retrieving sensor config and readings: %v\n", err)
		_ = req.Respond([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}

	responseBytes, err := json.Marshal(sensor)
	if err != nil {
		log.Printf("error marshalling sensor response: %v\n", err)
		_ = req.Respond([]byte(`{"error":"internal error"}`))
		return
	}

	if err := req.Respond(responseBytes); err != nil {
		log.Printf("error sending sensor response: %v\n", err)
	}
}
