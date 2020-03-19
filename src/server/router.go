package server

import (
	"todolist/src/server/middlewares"
)

func (s *server) configureRouter() {
	s.router.Use(middlewares.LogRequest)

	s.router.PathPrefix("/").HandlerFunc(defaultHandler)
}
