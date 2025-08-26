package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pimuzzo/go-chi-api/internal/model"
	"github.com/pimuzzo/go-chi-api/internal/service"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", GetUsers)
		r.Post("/", CreateUser)
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	service.SaveUser(u)
	w.WriteHeader(http.StatusCreated)
}
