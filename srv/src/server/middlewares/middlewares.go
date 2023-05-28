package middlewares

import (
	"net/http"
	"todolist/src/store"
)

type Middlewares struct {
	store *store.Store
}

func New(store *store.Store) *Middlewares {
	return &Middlewares{store}
}

func (md *Middlewares) LogRequest() func(http.Handler) http.Handler {
	return logRequest
}

func (md *Middlewares) CORS() func(http.Handler) http.Handler {
	return cors
}

func (md *Middlewares) Auth() func(http.Handler) http.Handler {
	return md.auth()
}
