package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route : A struct to hold the route details
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes : A list of Routes to be registered
type Routes []Route

// NewRouter : Function which registers routes in the Routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		Name:        "StartPage",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: StartPage,
	},
}
