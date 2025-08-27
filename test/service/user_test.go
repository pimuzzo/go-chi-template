package service_test

import (
	"testing"

	"github.com/pimuzzo/go-chi-template/internal/model"
	"github.com/pimuzzo/go-chi-template/internal/service"
	"github.com/pimuzzo/go-chi-template/pkg/logger"
)

func setup() {
	logger.Init()
	service.ResetUsers()
}

func TestSaveUser(t *testing.T) {
	setup()

	u := model.User{
		ID:    1,
		Name:  "Simone",
		Email: "simone@example.com",
	}

	service.SaveUser(u)
	users := service.GetAllUsers()

	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}

	if users[0].Name != "Simone" {
		t.Errorf("expected user name to be 'Simone', got '%s'", users[0].Name)
	}
}

func TestGetAllUsersEmpty(t *testing.T) {
	setup()

	users := service.GetAllUsers()

	if len(users) != 0 {
		t.Fatalf("expected empty user list, got %d users", len(users))
	}
}
