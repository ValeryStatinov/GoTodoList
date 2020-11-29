package server

import (
	"net/http"
	"os"
	"todolist/src/server/middlewares"
	"todolist/src/store"
	"todolist/src/systemlogger"

	"github.com/gorilla/mux"
)

type server struct {
	router      *mux.Router
	store       *store.Store
	middlewares *middlewares.Middlewares
}

func newServer() *server {
	store := store.New()
	srv := &server{
		router:      mux.NewRouter(),
		store:       store,
		middlewares: middlewares.New(store),
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

	port := ":" + os.Getenv("PORT")

	err := http.ListenAndServe(port, srv)
	if err != nil {
		systemlogger.Log("Failed to run server")
	}
}
