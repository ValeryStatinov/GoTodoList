package server

import (
	"fmt"
	"net/http"
	"todolist/src/database"

	"github.com/gorilla/mux"
)

type server struct {
	router   *mux.Router
	database *database.Database
}

func newServer() *server {
	srv := &server{
		router:   mux.NewRouter(),
		database: database.New(),
	}

	srv.configureRouter()

	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func StartServer() {
	srv := newServer()

	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		fmt.Println("FAIL")
	}
}
