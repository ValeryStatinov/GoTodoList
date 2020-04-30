package middlewares

import (
	"context"
	"net/http"
	"os"
	"todolist/src/server/ctxkeys"
	"todolist/src/store"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func (md *Middlewares) auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Token")
			claims := &store.JWTPayload{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			})
			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			user, ok := md.store.Users().GetById(claims.UserId)

			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ctxkeys.CtxUser, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
