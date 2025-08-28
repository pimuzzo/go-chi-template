package service

import (
	"github.com/pimuzzo/go-chi-template/internal/model"
	"github.com/pimuzzo/go-chi-template/pkg/logger"
)

var fakeDb = []model.User{}

func GetAllUsers() []model.User {
	logger.Log.Debug("Reading all users")
	return fakeDb
}

func SaveUser(u model.User) {
	logger.Log.Debug("Creating user")
	fakeDb = append(fakeDb, u)
	logger.Log.Info("User Created")
}

// ResetUsers clears the in-memory user store.
// Used only in tests to ensure isolation between runs.
func ResetUsers() {
	fakeDb = []model.User{}
}
