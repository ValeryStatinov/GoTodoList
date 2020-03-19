package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
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
