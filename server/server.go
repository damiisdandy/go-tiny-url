package server

import (
	"database/sql"
	"strconv"

	"github.com/damiisdandy/go-tiny-url/internal/database"
	"github.com/damiisdandy/go-tiny-url/utils"
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
	DB     *database.Queries
}

func New() *server {
	r := chi.NewRouter()

	port := utils.GetEnv("PORT")
	p, err := strconv.Atoi(port)

	if err != nil {
		p = PORT
	}

	return &server{
		Router: r,
		Port:   p,
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

	api.Get("/health", HealthCheck)

	api.Post("/shorten", s.CreateShortURL)
	api.Delete("/delete/{shortURL}", s.DeleteShortURL)
	api.Get("/{shortURL}", s.Redirect)

	s.Router.Mount("/v1", api)
}

func (s *server) ConnectDB() {
	conn, err := sql.Open("postgres", utils.GetEnv("DB_URL"))
	if err != nil {
		panic("Failed to connect to database")
	}

	s.DB = database.New(conn)
}
