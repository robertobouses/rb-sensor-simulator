package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func RunServer(nc *nats.Conn, app App) error {
	sensorConfigSvc, err := micro.AddService(nc, micro.Config{
		Name:        "sensor_config",
		Version:     "0.0.1",
		Description: "Create a new sensor",
	})
	if err != nil {
		return err
	}

	fmt.Printf("Created service: %s (%s)\n", sensorConfigSvc.Info().Name, sensorConfigSvc.Info().ID)

	sensorConfigGroup := sensorConfigSvc.AddGroup("sensor_config")
	handler := NewHandler(app)

	if err := sensorConfigGroup.AddEndpoint("create", micro.HandlerFunc(handler.CreateSensor)); err != nil {
		return err
	}
	if err := sensorConfigGroup.AddEndpoint("update", micro.HandlerFunc(handler.UpdateSensorConfig)); err != nil {
		return err
	}
	if err := sensorConfigGroup.AddEndpoint("get", micro.HandlerFunc(handler.UpdateSensorConfig)); err != nil {
		return err
	}

	if _, err := nc.Subscribe("sensor.reading", handler.ProcessSensorReading); err != nil {
		return err
	}
	return nil
}
