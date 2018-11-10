package main

import (
	"fmt"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Earthlings!")
}
