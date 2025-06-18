package use_cases

import "github.com/robertobouses/rb-sensor-simulator/internal/domain"

type Repository interface {
	Save(sensor *domain.Sensor) error
}
