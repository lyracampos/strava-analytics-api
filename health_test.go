package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthZ(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseWriter := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthZ)
	handler.ServeHTTP(responseWriter, request)

	obtainedStatus := responseWriter.Code
	expectedStatus := http.StatusOK

	if obtainedStatus != expectedStatus {
		t.Errorf("handler returned unexpected status code: got %v want %v", obtainedStatus, expectedStatus)
	}

	obtainedBody := responseWriter.Body.String()
	expectedBody := "api v3 is running"

	if obtainedBody != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", obtainedBody, expectedBody)
	}
}
