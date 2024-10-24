package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateJSONContentType(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be 'application/json'", http.StatusUnsupportedMediaType)
		return false
	}
	return true
}

func ValidateAndDecode(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	validate := validator.New()

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return false
	}

	if err := validate.Struct(v); err != nil {
		http.Error(w, "Validation Failed: "+err.Error(), http.StatusBadRequest)
		return false
	}

	return true
}

func SendJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
