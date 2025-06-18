package repository

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_sensor.sql
var saveSensorQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveSensorStmt, err := db.Prepare(saveSensorQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:         db,
		saveSensor: saveSensorStmt,
	}, nil
}

type Repository struct {
	db         *sql.DB
	saveSensor *sql.Stmt
}
