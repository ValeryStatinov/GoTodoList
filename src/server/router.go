package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todolist/src/models"
	"todolist/src/server/ctxkeys"
	"todolist/src/systemlogger"

	"github.com/gorilla/mux"
)

func (s *server) configureRouter() {
	s.router.Use(s.middlewares.LogRequest())
	s.router.Use(s.middlewares.CORS())

	withAuth := s.router.NewRoute().Subrouter()
	withAuth.Use(s.middlewares.Auth())
	withAuth.HandleFunc("/api/projects/", s.handleGetProjects())
	withAuth.HandleFunc("/api/projects/{projectId}/tasks/", s.handleGetTasks())

	s.router.PathPrefix("/").HandlerFunc(defaultHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func (s *server) handleGetProjects() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkeys.CtxUser).(*models.User)

		projects, ok := s.store.Projects().GetByUserId(user.Id)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
		}

		respondJson(w, projects)
	}
}

func (s *server) handleGetTasks() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkeys.CtxUser).(*models.User)
		vars := mux.Vars(r)
		projectId, err := strconv.Atoi(vars["projectId"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid project id=%s", vars["projectId"])
			return
		}

		project, ok := s.store.Projects().GetById(projectId)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "No project with id=%d", projectId)
			return
		}

		if project.UserId != user.Id {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		tasks, ok := s.store.Tasks().GetByProjectId(projectId)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respondJson(w, tasks)
	}
}

func respondJson(w http.ResponseWriter, data interface{}) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		systemlogger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
