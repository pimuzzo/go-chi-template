package model

import "github.com/go-playground/validator/v10"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
