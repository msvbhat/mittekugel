package main

import (
	"fmt"
	"net/http"
)

// StartPage : handles Welcome page
func StartPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Welcome Earthlings!")
}
