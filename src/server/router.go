package server

import (
	"todolist/src/server/handlers"
	"todolist/src/server/middlewares"
)

func (s *server) configureRouter() {
	s.router.Use(middlewares.LogRequest)

	s.router.PathPrefix("/").HandlerFunc(handlers.Default)
}
