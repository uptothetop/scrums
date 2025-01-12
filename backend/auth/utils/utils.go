package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, data interface{}, statusCode int) {
	// Set response header content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Encode the payload and write it
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
	}
}
