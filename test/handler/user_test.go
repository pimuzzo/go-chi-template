package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/pimuzzo/go-chi-template/internal/handler"
	"github.com/pimuzzo/go-chi-template/pkg/logger"
)

func setupRouter() http.Handler {
	logger.Init()
	r := chi.NewRouter()
	handler.RegisterRoutes(r)
	return r
}

func TestGetUsers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	setupRouter().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestCreateUser(t *testing.T) {
	body := `{"name": "Simone", "email": "simone@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	setupRouter().ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rr.Code)
	}
}
