package store

import (
	"database/sql"
	"todolist/src/models"
	"todolist/src/systemlogger"
)

type TasksManager struct {
	db *sql.DB
}

func (tm *TasksManager) GetByProjectId(id int) (*[]models.Task, bool) {
	tasks := make([]models.Task, 0)
	query := "select id, name, description, priority, completed from tasks where projectId=$1"

	rows, err := tm.db.Query(query, id)
	if err != nil {
		systemlogger.Log(err.Error(), query, string(id))
		return &tasks, false
	}
	defer rows.Close()

	for rows.Next() {
		task := models.Task{}
		err = rows.Scan(&task.Id, &task.Name, &task.Description, &task.Priority, &task.Completed)

		if err != nil {
			systemlogger.Log(err.Error(), query, string(id))
			return &tasks, false
		}

		tasks = append(tasks, task)
	}

	return &tasks, true
}

func (tm *TasksManager) Create(t *models.Task, projectId int) (*models.Task, bool) {
	query := "insert into tasks (name, description, priority, completed, projectId) values ($1, $2, $3, $4, $5) returning id"
	err := tm.db.QueryRow(query, t.Name, t.Description, t.Priority, false, projectId).Scan(&t.Id)
	if err != nil {
		systemlogger.Log(err.Error(), query, t.Name, t.Description, string(t.Priority), string(t.Id))
		return t, false
	}

	return t, true
}
