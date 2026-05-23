package services

import (
	"errors"
	"fmt"
	"todo-app/internal/models"
	"todo-app/internal/repository"
)

func RegisterUser(user models.User) error {
	minlength := 8
	if user.Email == "" {
		return errors.New("Email is required")
	}
	if len(user.Password) < minlength {
		return errors.New("weak password try another")
	}

	err := repository.RegisterUserQuery(user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to register user :%w", err)
	}

	return nil
}
