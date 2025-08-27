package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pimuzzo/go-chi-api/internal/handler"
	"github.com/pimuzzo/go-chi-api/pkg/logger"
)

// @title Go Chi Template
// @version 0.1.0
// @description This is a template REST API using Chi
// @host localhost:8080
// @BasePath /api/v1
func main() {
	logger.Init()

	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger.Log}))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		handler.RegisterRoutes(r)
	})

	logger.Log.Info("Starting server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Log.WithError(err).Fatal("Server failed to start")
	}
}
