package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pimuzzo/go-chi-api/internal/handler"
)

// @title Go Chi Template
// @version 0.1.0
// @description This is a template REST API using Chi
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		handler.RegisterRoutes(r)
	})

	http.ListenAndServe(":8080", r)
}
