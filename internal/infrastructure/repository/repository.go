package repository

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_sensor.sql
var saveSensorQuery string

//go:embed sql/save_sensor_reading.sql
var saveSensorReadingQuery string

//go:embed sql/update_sensor_last_reading.sql
var updateSensorLastReadingQuery string

//go:embed sql/update_sensor_config.sql
var updateSensorConfigQuery string

//go:embed sql/get_sensor_by_id.sql
var getSensorByIDQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveSensorStmt, err := db.Prepare(saveSensorQuery)
	if err != nil {
		return nil, err
	}
	saveSensorReadingStmt, err := db.Prepare(saveSensorReadingQuery)
	if err != nil {
		return nil, err
	}
	updateSensorLastReadingStmt, err := db.Prepare(updateSensorLastReadingQuery)
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

	return &Repository{
		db:                      db,
		saveSensor:              saveSensorStmt,
		saveSensorReading:       saveSensorReadingStmt,
		updateSensorLastReading: updateSensorLastReadingStmt,
		updateSensorConfig:      updateSensorConfigStmt,
		getSensorByID:           getSensorByIDStmt,
	}, nil
}

type Repository struct {
	db                      *sql.DB
	saveSensor              *sql.Stmt
	saveSensorReading       *sql.Stmt
	updateSensorLastReading *sql.Stmt
	updateSensorConfig      *sql.Stmt
	getSensorByID           *sql.Stmt
}
