package tasks

import (
	"database/sql"

	"github.com/valerystatinov/TodoList/database"
)

var insertIntoTasks = "INSERT INTO tasks (name, description, project_id) VALUES ($1, $2, $3)"

var getTasksByProjectId = "SELECT * FROM tasks WHERE project_id=$1"

func GetPreparedInsertIntoTasksStmt() *sql.Stmt {
	return database.PrepareSQL(insertIntoTasks)
}

func getPreparedGetTasksByProjectId() *sql.Stmt {
	return database.PrepareSQL(getTasksByProjectId)
}
