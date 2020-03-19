package database

import (
	"database/sql"
	"fmt"
	"os"
	"todolist/src/systemlogger"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func New() *Database {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	sqlConf := fmt.Sprintf("user=%s password=%s dbname=todolist sslmode=disable", dbUser, dbPass)
	db, err := sql.Open("postgres", sqlConf)
	if err != nil {
		systemlogger.Log(err.Error())
	}

	err = db.Ping()
	if err != nil {
		systemlogger.Log(err.Error())
	}

	database := Database{db}

	systemlogger.Log("Connected to database")

	return &database
}
