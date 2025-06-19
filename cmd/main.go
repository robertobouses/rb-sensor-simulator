package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
	"github.com/robertobouses/rb-sensor-simulator/internal/domain/use_cases"
	"github.com/robertobouses/rb-sensor-simulator/internal/infrastructure/repository"
	"github.com/robertobouses/rb-sensor-simulator/internal/pkg/postgres"
	"github.com/robertobouses/rb-sensor-simulator/internal/simulator"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := postgres.NewPostgres(postgres.DBConfig{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	repo, err := repository.NewRepository(db)
	if err != nil {
		log.Fatal("failed to init match repository:", err)
	}

	app := use_cases.NewApp(repo)

	sensor := domain.Sensor{
		Name: "Sensor 1",
		Type: domain.TemperatureSensor,
		Config: domain.SensorConfig{
			SamplingInterval: 10 * time.Second,
			AlertThreshold:   31.5,
			Unit:             "°C",
			Enabled:          true,
		},
		LastReading: &domain.SensorReading{
			Timestamp: time.Now(),
			Value:     27.7,
			Error:     nil,
		},
	}

	if err := app.SaveSensor(&sensor); err != nil {
		log.Fatal("Failed to save sensor:", err)
	}
	log.Printf("Sensor saved with ID: %v", sensor.ID)

	log.Println("Sensor saved successfully.")

	sensors := []domain.Sensor{sensor}

	sim := simulator.NewSensorSimulator(sensors, func(sensor domain.Sensor, reading domain.SensorReading) {
		err := app.SaveSensorReading(sensor, reading)
		if err != nil {
			log.Printf("Error saving sensor reading: %v", err)
		} else {
			log.Printf("Saved reading for sensor %s: %.2f %s", sensor.Name, reading.Value, reading.Error)
		}
	})

	sim.Start()

	time.Sleep(1 * time.Minute)

	sim.Stop()
	log.Println("Simulation stopped.")
}
