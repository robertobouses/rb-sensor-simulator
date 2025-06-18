package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
	"github.com/robertobouses/rb-sensor-simulator/internal/infrastructure/repository"
	"github.com/robertobouses/rb-sensor-simulator/internal/pkg/postgres"
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
	sensor := domain.Sensor{
		Name: "Sensor 1",
		Type: domain.TemperatureSensor,
		Config: domain.SensorConfig{
			SamplingInterval: 10 * time.Second,
			AlertThreshold:   30.5,
			Unit:             "°C",
			Enabled:          true,
		},
		LastReading: &domain.SensorReading{
			Timestamp: time.Now(),
			Value:     25.7,
			Error:     nil,
		},
	}

	if err := repo.SaveSensor(sensor); err != nil {
		log.Fatal("Failed to save sensor:", err)
	}

	log.Println("Sensor saved successfully.")
}
