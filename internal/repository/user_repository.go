package repository

import (
	"fmt"
	"todo-app/internal/db"
)

func RegisterUserQuery(email string, password string) error {
	_, err := db.DB.Exec(
		"INSERT INTO users (email,password)VALUES($1,$2) ", email, password,
	)
	if err != nil {
		return fmt.Errorf("failed to save user %v", err)
	}
	return nil
}
