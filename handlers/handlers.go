package handlers

import (
	"net/http"

	"github.com/damiisdandy/go-tiny-url/utils"
)

type ResponseData struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, ResponseData{
		Message: "Service is up and running",
		Status:  true,
		Data:    nil,
	})
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateShortURL"))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Redirect"))
}

func DeleteShortURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteShortURL"))
}
