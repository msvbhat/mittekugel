package main

import (
	"fmt"
	"net/http"
)

// StartPage : handles Welcome page
func StartPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Earthlings!")
}
