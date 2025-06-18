package repository

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_sensor.sql
var saveSensorQuery string

//go:embed sql/save_sensor_reading.sql
var saveSensorReadingQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveSensorStmt, err := db.Prepare(saveSensorQuery)
	if err != nil {
		return nil, err
	}
	saveSensorReadingStmt, err := db.Prepare(saveSensorReadingQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                db,
		saveSensor:        saveSensorStmt,
		saveSensorReading: saveSensorReadingStmt,
	}, nil
}

type Repository struct {
	db                *sql.DB
	saveSensor        *sql.Stmt
	saveSensorReading *sql.Stmt
}
