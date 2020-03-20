package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist/src/models"
	"todolist/src/server/ctxkeys"
	"todolist/src/systemlogger"
)

func (s *server) configureRouter() {
	s.router.Use(s.middlewares.LogRequest())
	s.router.Use(s.middlewares.CORS())

	withAuth := s.router.NewRoute().Subrouter()
	withAuth.Use(s.middlewares.Auth())
	withAuth.HandleFunc("/api/projects/", s.handleGetProjects())

	s.router.PathPrefix("/").HandlerFunc(defaultHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func (s *server) handleGetProjects() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkeys.CtxUser).(*models.User)

		projects := s.store.Projects().GetByUserId(user.Id)

		respondJson(w, projects)
	}
}

func respondJson(w http.ResponseWriter, data interface{}) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		systemlogger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
