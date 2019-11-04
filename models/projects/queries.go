package projects

import (
	"database/sql"

	"github.com/valerystatinov/TodoList/database"
)

var getProjects = "SELECT * from projects"

var postProject = "INSERT INTO projects (name) VALUES ($1)"

func GetPreparedGetProjectsStmt() *sql.Stmt {
	return database.PrepareSQL(getProjects)
}

func GetPreparedPostProjectStmt() *sql.Stmt {
	return database.PrepareSQL(postProject)
}
