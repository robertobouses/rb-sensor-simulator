package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
	natsx "github.com/robertobouses/rb-sensor-simulator/internal/infrastructure/nats"
)

const (
	typeIndex     = 0
	unitIndex     = 1
	numberIndex   = 2
	minIndex      = 3
	maxIndex      = 4
	samplingIndex = 5

	sensorReadingSubject = "sensor.reading"
)

func main() {
	sensorList, err := PrepareSensorsFromCSV()
	if err != nil {
		log.Fatal(err)
	}
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	for _, sensor := range sensorList {
		go SimulateReadings(nc, sensor)
	}

	select {}
}

func SimulateReadings(nc *nats.Conn, s domain.Sensor) {
	fmt.Printf("starting reading for sensor %s\n", s.Name)
	createRequest := natsx.CreateSensorRequest{
		ID:               s.ID.String(),
		Name:             s.Name,
		Unit:             s.Unit,
		Type:             string(s.Type),
		SamplingInterval: s.SamplingInterval,
		AlertThresholds: natsx.Threshold{
			Min: s.AlertThresholds.Min,
			Max: s.AlertThresholds.Max,
		},
	}
	req, err := json.Marshal(createRequest)
	if err != nil {
		log.Fatal(err)
	}
	_, err = nc.Request("sensor_config.create", req, time.Second*15)
	if err != nil {
		log.Fatalf("error calling sensor_config.create: %v", err)
	}
	ticker := time.NewTicker(s.SamplingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			payload := generateReadingPayload(s)
			if err := nc.Publish(sensorReadingSubject, payload); err != nil {
				log.Printf("error publishing mesage for sensor %s: %v", s.Name, err)
			}
		}
	}
}

func generateReadingPayload(s domain.Sensor) []byte {
	timestamp := time.Now()
	var value float64
	var errStr *string

	switch s.Type {
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

	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < 0.6 {
		if rand.Float64() < 0.5 {
			value += s.AlertThresholds.Max
		} else {
			value = -value - s.AlertThresholds.Min
		}
	}

	if rand.Float64() < 0.3 {
		msg := "sensor read error"
		errStr = &msg
		value = 0
	}

	reading := natsx.EventSensorReading{
		SensorID:  s.ID.String(),
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

func PrepareSensorsFromCSV() ([]domain.Sensor, error) {
	file, err := os.Open("sensors.csv")
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return nil, err
	}
	defer file.Close()

	var sensorList []domain.Sensor
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		log.Fatal(err)
	}
	for index, record := range records {
		if index == 0 {
			continue
		}

		sensorNumber, err := strconv.Atoi(strings.TrimSpace(record[numberIndex]))
		if err != nil {
			log.Fatalf("error parsing config sensor number: %v", err)
		}
		sensorMin, err := strconv.ParseFloat(strings.TrimSpace(record[minIndex]), 64)
		if err != nil {
			log.Fatalf("error parsing config sensor sensorMIn: %v", err)
		}
		sensorMax, err := strconv.ParseFloat(strings.TrimSpace(record[maxIndex]), 64)
		if err != nil {
			log.Fatalf("error parsing config sensor max: %v", err)
		}
		sensorSampling, err := time.ParseDuration(strings.TrimSpace(record[samplingIndex]))
		if err != nil {
			log.Fatalf("error parsing config sensor sampling: %v", err)
		}
		for i := 0; i < sensorNumber; i++ {
			sensorList = append(sensorList, domain.Sensor{
				ID:               uuid.New(),
				Name:             fmt.Sprintf("%s sensor %d%d", record[typeIndex], index, i),
				Type:             domain.SensorType(record[typeIndex]),
				Unit:             record[unitIndex],
				SamplingInterval: sensorSampling,
				AlertThresholds: domain.Threshold{
					Min: sensorMin,
					Max: sensorMax,
				},
			})
		}
	}
	return sensorList, nil
}
