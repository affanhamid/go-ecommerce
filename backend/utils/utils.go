package utils

import (
	"encoding/json"
	"log"
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

func SendJSON(w http.ResponseWriter, status int, jsonBytes []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)
}

func SendJSONMessage(w http.ResponseWriter, status int, message string) {
	message_json, err := json.Marshal(map[string]string{"message": message})
	if err != nil {
		log.Fatalf("something went wrong: %v", err.Error())
	}
	SendJSON(w, status, message_json)
}
