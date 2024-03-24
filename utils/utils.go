package utils

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
)

const (
	ALLOWED_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
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
	}
	type errorResponse struct {
		Error  string `json:"error"`
		Status bool   `json:"status"`
	}

	RespondWithJSON(w, code, errorResponse{Error: message, Status: false})
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable " + key + " is not set.")
	}
	return value
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = ALLOWED_CHARS[rand.Intn(len(ALLOWED_CHARS))]
	}
	return string(b)
}
