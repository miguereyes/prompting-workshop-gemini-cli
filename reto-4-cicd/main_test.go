package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTrackVisitHandler_MissingUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/visit", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(trackVisitHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler devolvió el código de estado incorrecto: obtenido %v, se esperaba %v",
			status, http.StatusBadRequest)
	}
}
