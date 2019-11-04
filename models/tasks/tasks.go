package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID   int    `json:"projectId"`
}

func newTask(id int, name string, description string, projectID int) Task {
	task := Task{id, name, description, projectID}
	return task
}

func HandleTasks(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		queryParams := request.URL.Query()
		var projectId string
		if len(queryParams["projectId"]) == 1 && queryParams["projectId"][0] != "" {
			projectId = queryParams["projectId"][0]
		} else {
			http.Error(writer, "wrong query param", 404)
			return
		}

		rows, err := getPreparedGetTasksByProjectId().Query(projectId)

		if err != nil {
			fmt.Println("error query")
			return
		}
		defer rows.Close()

		tasks := make([]Task, 0)
		for rows.Next() {
			var id, projectId int
			var name, description string

			err := rows.Scan(&id, &name, &description, &projectId)
			if err != nil {
				fmt.Println("error scan")
				return
			}

			tasks = append(tasks, newTask(id, name, description, projectId))
		}

		js, err := json.Marshal(tasks)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(js)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}

	if request.Method == "POST" {
		var task Task

		body, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if err != nil {
			http.Error(writer, err.Error(), 500)
			return
		}

		err = json.Unmarshal(body, &task)
		if err != nil {
			http.Error(writer, err.Error(), 500)
			return
		}

		_, err = GetPreparedInsertIntoTasksStmt().Exec(task.Name, task.Description, task.ProjectID)
		if err != nil {
			http.Error(writer, "Bad request", http.StatusBadRequest)
			return
		}

		writer.WriteHeader(200)
	}
}
