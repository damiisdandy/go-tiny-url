package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code < 499 {
		log.Printf("Error: %s", message)
		return
	}
	type errorResponse struct {
		Error  string `json:"error"`
		Status bool   `json:"status"`
	}

	RespondWithJSON(w, code, errorResponse{Error: message, Status: false})
}
