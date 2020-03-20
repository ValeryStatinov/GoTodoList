package server

import (
	"net/http"
	"todolist/src/store"
	"todolist/src/systemlogger"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  *store.Store
}

func newServer() *server {
	srv := &server{
		router: mux.NewRouter(),
		store:  store.New(),
	}

	srv.configureRouter()

	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func StartServer() {
	systemlogger.Log("Starting server...")

	srv := newServer()

	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		systemlogger.Log("Failed to run server")
	}
}
