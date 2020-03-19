package server

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
