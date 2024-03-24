package server

import (
	"github.com/damiisdandy/go-tiny-url/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	PORT = 8000
)

type server struct {
	Router *chi.Mux
	Port   int
	// TODO: Add a database connection here
}

func New() *server {
	r := chi.NewRouter()

	return &server{
		Router: r,
		Port:   PORT,
	}
}

func (s *server) MountMiddlewares() {
	// Logging
	s.Router.Use(middleware.Logger)
	// CORS
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}

func (s *server) MountHandlers() {
	api := chi.NewRouter()

	api.Get("/health", handlers.HealthCheck)

	api.Post("/shorten", handlers.CreateShortURL)
	api.Delete("/delete/{shortURL}", handlers.DeleteShortURL)
	api.Get("/{shortURL}", handlers.Redirect)

	s.Router.Mount("/v1", api)
}
