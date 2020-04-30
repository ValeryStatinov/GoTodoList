package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todolist/src/models"
	"todolist/src/server/ctxkeys"
	"todolist/src/server/middlewares"
	"todolist/src/store"
	"todolist/src/systemlogger"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func (s *server) configureRouter() {
	s.router.Use(s.middlewares.LogRequest())
	s.router.Use(s.middlewares.CORS())

	s.router.HandleFunc("/api/login/", s.handleLogin()).Methods("POST")
	s.router.HandleFunc("/api/register/", s.handleRegister()).Methods("POST")

	withAuth := s.router.NewRoute().Subrouter()
	withAuth.Use(s.middlewares.Auth())
	withAuth.HandleFunc("/api/projects/", s.handleGetProjects()).Methods("GET")
	withAuth.HandleFunc("/api/projects/", s.handleCreateProject()).Methods("POST")
	withAuth.HandleFunc("/api/projects/{projectId}/tasks/", s.handleGetTasks()).Methods("GET")
	withAuth.HandleFunc("/api/projects/{projectId}/tasks/", s.handleCreateTask()).Methods("POST")
	withAuth.HandleFunc("/api/projects/{projectId}/tasks/{taskId}/", s.handleUpdateTask()).Methods("PUT")

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
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No project with id=%d", projectId)
			return
		}

		if !user.HaveAccesToProject(project) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		tasks, ok := s.store.Tasks().GetByProjectId(projectId)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJson(w, tasks)
	}
}

func (s *server) handleCreateProject() func(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		requestData := &request{}
		user := r.Context().Value(ctxkeys.CtxUser).(*models.User)
		err := json.NewDecoder(r.Body).Decode(requestData)
		if err != nil {
			systemlogger.Log(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		project := &models.Project{Name: requestData.Name}
		if !project.Validate() {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Project name cannot be empty")
			return
		}

		if !s.store.Projects().Create(project, user) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJson(w, project)
	}
}

func (s *server) handleCreateTask() func(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Priority    uint8  `json:"priority"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		requestData := &request{}
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
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No project with id=%d", projectId)
			return
		}

		if !user.HaveAccesToProject(project) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		err = json.NewDecoder(r.Body).Decode(requestData)
		if err != nil {
			systemlogger.Log(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid request body")
			return
		}

		task := &models.Task{
			Name:        requestData.Name,
			Description: requestData.Description,
			Priority:    requestData.Priority,
		}

		if !task.Validate() {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "All fields (name, description, priority) must be filled; priority is int from 1 to 3")
			return
		}

		task, ok = s.store.Tasks().Create(task, projectId)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJson(w, task)
	}
}

func (s *server) handleUpdateTask() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkeys.CtxUser).(*models.User)
		vars := mux.Vars(r)

		projectId, err := strconv.Atoi(vars["projectId"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Invalid project id=%s", vars["projectId"])
			return
		}
		taskId, err := strconv.Atoi(vars["taskId"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Invalid task id=%s", vars["taskId"])
			return
		}

		project, ok := s.store.Projects().GetById(projectId)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No project with id=%d", projectId)
			return
		}

		if !user.HaveAccesToProject(project) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if !s.store.Tasks().BelongsToProject(taskId, projectId) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No task with id=%d in projectId=%d", taskId, projectId)
			return
		}

		task := &models.Task{}

		err = json.NewDecoder(r.Body).Decode(task) // TODO completed might be missed
		if err != nil {
			systemlogger.Log(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}

		if !task.Validate() {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "All fields (name, description, priority, completed, projectId) must be filled; priority is int from 1 to 3")
			return
		}

		_, err = s.store.Tasks().Update(task, taskId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) handleLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials = &store.Credentials{}

		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, ok := s.store.Users().GetByName(credentials.Login)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid login or password")
			return
		}

		if models.CheckPassword(user, credentials.Login, credentials.Password) {
			claims := &store.JWTPayload{
				UserId:         user.Id,
				StandardClaims: jwt.StandardClaims{},
				Time:           time.Now().Unix(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenString, err := token.SignedString(middlewares.JwtKey)
			if err != nil {
				// If there is an error in creating the JWT return an internal server error
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			type ResponseWithToken struct {
				Token string `json:"token"`
			}

			respondJson(w, ResponseWithToken{tokenString})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid login or password")
		}
	}
}

func (s *server) handleRegister() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials = &store.Credentials{}
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(credentials.Password) < 6 || len(credentials.Login) < 4 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Login should contain >= 4 symbols, password >= 6")
			return
		}

		_, ok := s.store.Users().GetByName(credentials.Login)
		if ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "This login name is already taken")
			return
		}

		user, ok := s.store.Users().Create(credentials.Login, credentials.Password)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		claims := &store.JWTPayload{
			UserId:         user.Id,
			StandardClaims: jwt.StandardClaims{},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(middlewares.JwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		type ResponseWithToken struct {
			Token string `json:"token"`
		}

		respondJson(w, ResponseWithToken{tokenString})
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
