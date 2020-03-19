package handlers

import (
	"fmt"
	"net/http"
)

func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
