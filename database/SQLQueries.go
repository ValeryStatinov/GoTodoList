package database

import (
	"database/sql"
)

func prepare(query string) *sql.Stmt {
	db := GetDBInstance()
	stmt, err := db.Prepare(query)
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}

	return stmt
}

var insertIntoTasks = "INSERT INTO tasks (NAME, DESCRIPTION, PROJECT_ID, USER_ID) VALUES (?, ?, ?, ?)"

func GetPreparedInsertIntoTasksStmt() *sql.Stmt {
	return prepare(insertIntoTasks)
}

var getTasksByUserID = "SELECT * FROM tasks where USER_ID=$1"

func GetTasksByUserIDStmt() *sql.Stmt {
	return prepare(getTasksByUserID)
}
