package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	appErrors "app/pkg/errors"
)

func RespondWithError(w http.ResponseWriter, err *appErrors.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	if err := json.NewEncoder(w).Encode(map[string]string{"error": err.Message}); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
