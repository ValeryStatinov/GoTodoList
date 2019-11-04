package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "user=admin password=admin dbname=todolist sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

func GetDBInstance() *sql.DB {
	if db != nil {
		return db
	}
	panic("Can't return pointer to db, no established connection")
}
