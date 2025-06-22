package repository

import (
	"log"

	"github.com/google/uuid"
)

func (r *Repository) UpdateAlertResolved(sensorID, resolvedReadingID uuid.UUID) error {
	_, err := r.updateAlertResolved.Exec(sensorID, resolvedReadingID)
	if err != nil {
		log.Printf("Error updating alert_historial on UpdateAlertResolved: %v", err)
		return err
	}
	return nil
}
