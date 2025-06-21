package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain/use_cases"
	natsx "github.com/robertobouses/rb-sensor-simulator/internal/infrastructure/nats"
	"github.com/robertobouses/rb-sensor-simulator/internal/infrastructure/repository"
	"github.com/robertobouses/rb-sensor-simulator/internal/pkg/postgres"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("failed to connect to NATS:", err)
	}
	defer nc.Close()

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
		log.Fatal("failed to init repository:", err)
	}

	app := use_cases.NewApp(repo)

	err = natsx.RunServer(nc, app)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("running nats api")
	select {}
}
