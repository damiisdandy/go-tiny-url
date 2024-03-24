package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/damiisdandy/go-tiny-url/internal/database"
	"github.com/damiisdandy/go-tiny-url/utils"
	"github.com/go-chi/chi/v5"
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
		utils.RespondWithError(w, http.StatusBadRequest, "Internal server error")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, url)
}

func (s *server) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	url, err := s.DB.GetURL(r.Context(), shortURL)
	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching URL: %q", shortURL))
		return
	}
	http.Redirect(w, r, url.OriginalUrl, http.StatusMovedPermanently)
}

func (s *server) DeleteShortURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	_, err := s.DB.DeleteURL(r.Context(), shortURL)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting URL: %q", shortURL))
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, nil)
}
