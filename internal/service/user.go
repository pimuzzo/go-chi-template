package service

import "github.com/pimuzzo/go-chi-api/internal/model"

var users = []model.User{}

func GetAllUsers() []model.User {
	return users
}

func SaveUser(u model.User) {
	users = append(users, u)
}
