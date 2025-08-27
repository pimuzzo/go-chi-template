package service

import (
	"github.com/pimuzzo/go-chi-api/internal/model"
	"github.com/pimuzzo/go-chi-api/pkg/logger"
)

var users = []model.User{}

func GetAllUsers() []model.User {
	logger.Log.Debug("Reading all users")
	return users
}

func SaveUser(u model.User) {
	logger.Log.Debug("Creating user")
	users = append(users, u)
	logger.Log.Info("User Created")
}
