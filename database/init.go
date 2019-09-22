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

func ShowTasksTable() {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		fmt.Println("error query")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, project_id, user_id int
		var name, description string

		err := rows.Scan(&id, &name, &description, &project_id, &user_id)
		if err != nil {
			fmt.Println("error scan")
			return
		}
		fmt.Println(id, name, description, project_id, user_id)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("error rows")
		return
	}
}
