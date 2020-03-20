package middlewares

import (
	"context"
	"net/http"
	"todolist/src/models"
	"todolist/src/server/ctxkeys"
)

func (md *Middlewares) auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Header.Get("Token")
			if userId == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var user *models.User
			var ok bool
			user, ok = md.store.Users().GetByName(userId)

			if !ok {
				user, ok = md.store.Users().Create(userId)
				if !ok {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			ctx := context.WithValue(r.Context(), ctxkeys.CtxUser, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
