package server

import (
	"fmt"
	"net/http"
	"todolist/src/server/middlewares"
)

func (s *server) configureRouter() {
	s.router.Use(middlewares.LogRequest)

	s.router.PathPrefix("/").HandlerFunc(defaultHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
