package database

import (
	"database/sql"
	"todolist/src/systemlogger"
)

type Database struct {
	db *sql.DB
}

func New() *Database {
	db, err := sql.Open("postgres", "user=admin password=admin dbname=todolist sslmode=disable")
	if err != nil {
		systemlogger.Log(err.Error())
	}

	database := Database{db}

	return &database
}
