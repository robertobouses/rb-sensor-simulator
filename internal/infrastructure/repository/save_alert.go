package repository

import (
	"log"

	"github.com/robertobouses/rb-sensor-simulator/internal/domain"
)

func (r *Repository) SaveAlert(alert domain.AlertHistorial) error {
	_, err := r.saveAlert.Exec(
		alert.SensorID,
		alert.AlertType,
		alert.WarningReadingID,
	)

	if err != nil {
		log.Printf("Error executing SaveAlert statement: %v", err)
		return err
	}

	return nil
}
