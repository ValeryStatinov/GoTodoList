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
	UserID      int    `json:"userId"`
}

func newTask(id int, name string, description string, projectID int, userID int) *Task {
	task := Task{id, name, description, projectID, userID}
	return &task
}

func HandleTasks(writer http.ResponseWriter, req *http.Request) {
	var task Task = Task{UserID: 1}
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	fmt.Println(task)
}
