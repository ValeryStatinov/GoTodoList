package database

import (
	"database/sql"
	"log"
)

func prepare(query string) *sql.Stmt {
	db := GetDBInstance()
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	return stmt
}

var insertIntoTasks = "INSERT INTO tasks (ID, NAME, DESCRIPTION, PROJECT_ID, USER_ID) VALUES (?, ?, ?, ?, ?)"

// GetPreparedInsertIntoTasksStmt : returns prepared for execution statement
func GetPreparedInsertIntoTasksStmt() *sql.Stmt {
	return prepare(insertIntoTasks)
}

var getTasksByUserID = "SELECT * FROM tasks where USER_ID=$1"

// GetTasksByUserIDStmt : returns prepared for execution statement
func GetTasksByUserIDStmt() *sql.Stmt {
	return prepare(getTasksByUserID)
}
