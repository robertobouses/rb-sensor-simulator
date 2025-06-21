package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
	natsx "github.com/robertobouses/rb-sensor-simulator/internal/infrastructure/nats"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 0; i < 10; i++ {
		sensorType := getRandomSensorType()
		sensorID := getRandomSensorID()
		reading := generateReadingPayload(sensorID, sensorType)
		if err := nc.Publish("sensor.reading", reading); err != nil {
			fmt.Println(err)
		}
	}
}

func getRandomSensorType() domain.SensorType {
	types := []domain.SensorType{
		domain.TemperatureSensor,
		domain.HumiditySensor,
		domain.PressureSensor,
	}
	return types[rand.Intn(len(types))]
}

func getRandomSensorID() string {
	sensorIDs := []string{
		"43f0e336-c080-4c9a-ab90-54542597a77b",
		"d05d1a40-67e7-4f69-b107-c3f1fdfae974",
		"b7515b94-2edb-445e-94a8-d16747144966",
	}
	return sensorIDs[rand.Intn(len(sensorIDs))]
}

func generateReadingPayload(sensorID string, sensorType domain.SensorType) []byte {
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

	reading := natsx.EventSensorReading{
		SensorID:  sensorID,
		Timestamp: timestamp,
		Value:     value,
		Error:     errStr,
	}
	data, err := json.Marshal(reading)
	if err != nil {
		log.Printf("Failed to marshal reading: %v", err)
		return []byte{}
	}

	return data
}
