package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"todo-app/internal/db"
	"todo-app/internal/models"
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

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id ,email,password FROM user WHERE email=$1`
	rows := db.DB.QueryRow(query, email)
	err := rows.Scan(&user.Id, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
