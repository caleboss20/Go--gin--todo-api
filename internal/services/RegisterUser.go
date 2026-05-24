package services

import (
	"errors"
	"fmt"
	"todo-app/internal/models"
	"todo-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user models.User) error {
	minlength := 8
	if user.Email == "" {
		return errors.New("Email is required")
	}
	if len(user.Password) < minlength {
		return errors.New("weak password try another")
	}

	//hashing user password before adding to database//
	//14 is the cost factor:controls how slow the hashing is
	// higher=slower=harder to brute force attack//
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}

	//saving user to database
	err = repository.RegisterUserQuery(user.Email, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("failed to register user :%w", err)
	}

	return nil
}

func Try(user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	err = repository.RegisterUserQuery(user.Email, string(hash))
	if err != nil {
		return errors.New("failed to save user")
	}
	return nil
}
