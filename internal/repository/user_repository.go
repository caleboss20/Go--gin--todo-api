package repository

import "todo-app/internal/db"

func RegisterUserQuery(email string, password string) error {
	_, err := db.DB.Exec(
		"INSERT INTO users (email,password)VALUES($1,$2)", email, password,
	)
	return err
}
