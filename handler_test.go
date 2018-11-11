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
