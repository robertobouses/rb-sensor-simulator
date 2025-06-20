package nats

import (
	"github.com/robertobouses/rb-sensor-simulator/internal/domain/use_cases"
)

type Handler struct {
	app use_cases.AppService
}

func NewHandler(app use_cases.AppService) *Handler {
	return &Handler{app: app}
}
