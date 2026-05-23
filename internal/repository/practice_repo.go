package repository

import (
	"database/sql"
	"fmt"
	"todo-app/internal/db"
	"todo-app/internal/models"
)

func CreateQuery(title string) error {
	_, err := db.DB.Exec(
		"INSERT INTO todos (title)VALUES($1)",
	)
	return err
}

func UpdateQuery(title string, id int) error {
	_, err := db.DB.Exec(
		"UPDATE todos SET title=$1 WHERE id=$2", title,
	)
	return err

}
func DeleteQuery(id int) error {
	_, err := db.DB.Exec(
		"DELETE FROM todos WHERE id=$1", id,
	)
	return err
}

func GetAllQuery() ([]models.Todo, error) {
	rows, err := db.DB.Query(
		"SELECT id,title,done FROM todos ",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get todos %v", err)
	}

	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		rows.Scan(&t.ID, &t.Title, &t.Done)
		todos = append(todos, t)
	}

	return todos, nil
}

func GetAllBookings() ([]models.Bookings, error) {
	rows, err := db.DB.Query(
		"SELECT id,price,booked FROM todos",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bookings %v", err)
	}
	defer rows.Close()

	var bookings []models.Bookings
	for rows.Next() {
		var b models.Bookings
		rows.Scan(&b.ID, &b.Price, &b.Booked)
		bookings = append(bookings, b)

	}
	return bookings, nil

}

func GetATodo(id int) (models.Todo, error) {
	var t models.Todo
	row := db.DB.QueryRow("SELECT id,title,done FROM todos WHERE id=$1", id)
	err := row.Scan(&t.ID, &t.Title, &t.Done)
	if err == sql.ErrNoRows {
		return models.Todo{}, fmt.Errorf("cannot find todo %d", id)
	}
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to fetch todo item %d", id)
	}

	return t, nil
}
