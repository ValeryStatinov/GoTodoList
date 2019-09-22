package main

import (
	"fmt"
	"net/http"

	"github.com/valerystatinov/TodoList/database"
	"github.com/valerystatinov/TodoList/models/tasks"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("index")
	fmt.Fprintf(writer, "Hello world")
}

func main() {
	database.InitDB()
	database.ShowTasksTable()

	http.HandleFunc("/tasks/", tasks.HandleTasks)
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

// var lastInsertedID = -1
// err = db.QueryRow("INSERT INTO tasks (NAME) VALUES ('Kostya') RETURNING ID").Scan(&lastInsertedID)
// fmt.Println(lastInsertedID)
