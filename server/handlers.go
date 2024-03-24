package server

import (
	"encoding/json"
	"net/http"

	"github.com/damiisdandy/go-tiny-url/internal/database"
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

func (s *server) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	type RequestData struct {
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	requestData := RequestData{}
	err := decoder.Decode(&requestData)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, ResponseData{
			Message: "Invalid request body",
			Status:  false,
			Data:    nil,
		})
		return
	}
	url, err := s.DB.CreateURL(r.Context(), database.CreateURLParams{
		UrlID:       utils.GenerateRandomString(6),
		OriginalUrl: requestData.URL,
	})
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, ResponseData{
			Message: "Internal server error",
			Status:  false,
			Data:    nil,
		})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, url)
}

func (s *server) Redirect(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Redirect"))
}

func (s *server) DeleteShortURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteShortURL"))
}
