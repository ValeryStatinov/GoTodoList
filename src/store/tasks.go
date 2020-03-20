package store

import (
	"database/sql"
	"todolist/src/models"
	"todolist/src/systemlogger"
)

type TasksManager struct {
	db *sql.DB
}

func (tm *TasksManager) GetAll() []models.Task {
	tasks := make([]models.Task, 0)

	rows, err := tm.db.Query("SELECT * from tasks")
	if err != nil {
		systemlogger.Log(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		task := models.Task{}
		err = rows.Scan(&task.Id, &task.Name, &task.Description, &task.Priority, &task.Completed,
			&task.ProjectId)

		if err != nil {
			systemlogger.Log(err.Error())
		}

		tasks = append(tasks, task)
	}

	return tasks
}
