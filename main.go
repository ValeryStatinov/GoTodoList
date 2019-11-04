package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valerystatinov/TodoList/database"
	"github.com/valerystatinov/TodoList/models/projects"
	"github.com/valerystatinov/TodoList/models/tasks"
)

func main() {
	database.InitDB()

	router := mux.NewRouter()

	router.HandleFunc("/projects/", projects.HandleProjects).Methods("GET", "POST")
	router.HandleFunc("/tasks/", tasks.HandleTasks).Methods("GET", "POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// router.NotFoundHandler = http.HandlerFunc(customNotFoundHandler)

	http.Handle("/", handlers.CORS(originsOk, headersOk, methodsOk)(router))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
