package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test(t *testing.T, req *http.Request, controller func(w http.ResponseWriter, r *http.Request), expectedStatus int, expectedErrorMessage string) {
	/*
		Generic testing function
	*/

	rr := httptest.NewRecorder()
	handler := http.Handler(http.HandlerFunc(controller))

	handler.ServeHTTP(rr, req)

	// Checking error code
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}

	actual := strings.TrimSpace(rr.Body.String())

	// Checking the right error message
	if expectedErrorMessage != "" && actual != expectedErrorMessage {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'", actual, expectedErrorMessage)
	}

}
