package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StartPage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("StartPage returned wrong status code")
	}

	expected := "Welcome Earthlings!"
	if rr.Body.String() != expected {
		t.Errorf("StartPage returned wrong Body: got %v expected %v",
			rr.Body.String(), expected)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "text/plain" {
		t.Errorf("StartPage content-type does not match: got %v expected %v",
			ctype, "text/plain")
	}

}

func TestGetNodeMetrics(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/analytics/nodes/average", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetNodeMetrics)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetNodeMetrics returned wrong status code")
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("GetNodeMetrics Content-Type don't match: got %v expected %v",
			ctype, "application/json")
	}

}

func TestPostNodeMetrics(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/metrics/node/mittekugel", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostNodeMetrics)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("PostNodeMetrics returned wrong status code")
	}

}
