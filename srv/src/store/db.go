package store

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"todolist/src/systemlogger"

	_ "github.com/lib/pq"
)

type Store struct {
	db       *sql.DB
	tasks    *TasksManager
	projects *ProjectsManager
	users    *UsersManager
}

func New() *Store {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	var db *sql.DB
	var err error

	fmt.Println("Waiting for DB to initialize...")
	time.Sleep(20 * time.Second)

	sqlConf := fmt.Sprintf("host=todolist_db_1 port=5432 user=%s password=%s dbname=todolist sslmode=disable", dbUser, dbPass)
	db, err = sql.Open("postgres", sqlConf)
	for err != nil {
		if err != nil {
			systemlogger.Log(err.Error())
		}

		err = db.Ping()
		if err != nil {
			systemlogger.Log(err.Error())
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}

	database := Store{db: db}

	systemlogger.Log("Connected to database")

	go database.pingDB()

	return &database
}

func (s *Store) pingDB() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	sqlConf := fmt.Sprintf("host=todolist_db_1 port=5432 user=%s password=%s dbname=todolist sslmode=disable", dbUser, dbPass)

	err := s.db.Ping()
	if err != nil {
		systemlogger.Log(err.Error())
		s.db, _ = sql.Open("postgres", sqlConf)
	} else {
		fmt.Println("Connection with db is OK")
	}
	time.Sleep(3 * time.Minute)
	go s.pingDB()
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

func (s *Store) Users() *UsersManager {
	if s.users == nil {
		s.users = &UsersManager{s.db}
	}

	return s.users
}
