package repository

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_sensor.sql
var saveSensorQuery string

//go:embed sql/save_sensor_reading.sql
var saveSensorReadingQuery string

//go:embed sql/update_sensor_config.sql
var updateSensorConfigQuery string

//go:embed sql/get_sensor_by_id.sql
var getSensorByIDQuery string

//go:embed sql/get_sensor_last_readings_by_id.sql
var getSensorLastReadingsByIDQuery string

//go:embed sql/save_alert.sql
var saveAlertQuery string

//go:embed sql/update_alert_resolved.sql
var updateAlertResolvedQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveSensorStmt, err := db.Prepare(saveSensorQuery)
	if err != nil {
		return nil, err
	}
	saveSensorReadingStmt, err := db.Prepare(saveSensorReadingQuery)
	if err != nil {
		return nil, err
	}
	updateSensorConfigStmt, err := db.Prepare(updateSensorConfigQuery)
	if err != nil {
		return nil, err
	}
	getSensorByIDStmt, err := db.Prepare(getSensorByIDQuery)
	if err != nil {
		return nil, err
	}
	getSensorLastReadingsByIDStmt, err := db.Prepare(getSensorLastReadingsByIDQuery)
	if err != nil {
		return nil, err
	}
	saveAlertStmt, err := db.Prepare(saveAlertQuery)
	if err != nil {
		return nil, err
	}
	updateAlertResolvedStmt, err := db.Prepare(updateAlertResolvedQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                        db,
		saveSensor:                saveSensorStmt,
		saveSensorReading:         saveSensorReadingStmt,
		updateSensorConfig:        updateSensorConfigStmt,
		getSensorByID:             getSensorByIDStmt,
		getSensorLastReadingsByID: getSensorLastReadingsByIDStmt,
		saveAlert:                 saveAlertStmt,
		updateAlertResolved:       updateAlertResolvedStmt,
	}, nil
}

type Repository struct {
	db                        *sql.DB
	saveSensor                *sql.Stmt
	saveSensorReading         *sql.Stmt
	updateSensorConfig        *sql.Stmt
	getSensorByID             *sql.Stmt
	getSensorLastReadingsByID *sql.Stmt
	saveAlert                 *sql.Stmt
	updateAlertResolved       *sql.Stmt
}
