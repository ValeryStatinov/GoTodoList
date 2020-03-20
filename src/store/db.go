package store

import (
	"database/sql"
	"fmt"
	"os"
	"todolist/src/systemlogger"

	_ "github.com/lib/pq"
)

type Store struct {
	db       *sql.DB
	tasks    *TasksManager
	projects *ProjectsManager
}

func New() *Store {
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

	database := Store{db: db}

	systemlogger.Log("Connected to database")

	return &database
}

func (s *Store) Tasks() *TasksManager {
	if s.tasks == nil {
		s.tasks = &TasksManager{s.db}
	}

	return s.tasks
}

func (s *Store) Projects() *ProjectsManager {
	if s.projects == nil {
		s.projects = &ProjectsManager{s.db}
	}

	return s.projects
}
