package store

import (
	"database/sql"
	"fmt"
	"todolist/src/models"
	"todolist/src/systemlogger"
)

type TasksManager struct {
	db *sql.DB
}

func (tm *TasksManager) GetAll() {
	tasks := make([]models.Task, 0)

	rows, err := tm.db.Query("SELECT * from tasks")
	if err != nil {
		systemlogger.Log(err.Error())
	}

	for rows.Next() {
		task := models.Task{}
		err = rows.Scan(&task.Id, &task.Name, &task.Description, &task.Priority, &task.Completed)
		tasks = append(tasks, task)
	}

	for _, v := range tasks {
		fmt.Printf("%d %s %s %d %t\n", v.Id, v.Name, v.Description, v.Priority, v.Completed)
	}
}
