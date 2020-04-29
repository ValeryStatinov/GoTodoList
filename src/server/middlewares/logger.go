package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"todolist/src/systemlogger"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.code = code
	rw.ResponseWriter.WriteHeader(code)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		record := fmt.Sprintf("%s %s %d %s", r.Method, r.RequestURI, rw.code, time.Now().Sub(start))

		path := os.Getenv("REQUEST_LOG_PATH")
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			systemlogger.Log(err.Error())
		}
		defer f.Close()

		logger := log.New(f, "", log.LstdFlags|log.LUTC)
		logger.Println(record)
	})
}
