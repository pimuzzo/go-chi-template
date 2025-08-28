package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pimuzzo/go-chi-template/internal/model"
	"github.com/pimuzzo/go-chi-template/internal/service"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", GetUsers)
		r.Post("/", CreateUser)
	})
}

// GetUsers godoc
// @Summary Get users
// @Tags User
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

// CreateUser godoc
// @Summary Create user
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.User true "User to create"
// @Success 201
// @Failure 400 {string} string "Bad request"
// @Failure 422 {string} string "Validation error"
// @Router /users [post]
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
