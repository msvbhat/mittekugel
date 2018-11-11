package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// StartPage : handles Welcome page
func StartPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Welcome Earthlings!")
}

// PostNodeMetrics : handles the POST of single node metrics
func PostNodeMetrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodename := vars["nodename"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "This metric is from:", nodename)
}

// GetNodeMetrics : handles GET on average node metrics
func GetNodeMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	p := NodeMetrics{
		Timeslice:   60.0,
		NodeCPUTime: 42.0,
		NodeMemUsed: 8.0,
	}
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}
