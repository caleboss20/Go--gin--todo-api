package repository

import (
	"database/sql"
	"fmt"
	"todo-app/internal/db"
	"todo-app/internal/models"
)

func CreatingTodoQuery(title string) error {
	_, err := db.DB.Exec(
		"INSERT INTO todos (title) VALUES($1)", title,
	)
	return err
}

func UpdatatingTodoQuery(title string, id int) error {
	_, err := db.DB.Exec(
		"UPDATE todos SET title=$1 WHERE id=$2", title, id,
	)
	return err
}

func DeleteTodoQuery(id int) error {
	_, err := db.DB.Exec(
		"DELETE FROM todos WHERE id=$1", id,
	)
	return err
}

// func SelectAllTodoQuery() ([]models.Todo, error) {
// 	rows, err := db.DB.Query(
// 		"SELECT id,title,done FROM todos",
// 	)
// 	if err != nil {
// 		fmt.Errorf(" failed to fetch todos %v", err)
// 	}
// 	defer rows.Close()

// 	var todos []models.Todo

// 	for rows.Next() {
// 		var t models.Todo
// 		rows.Scan(&t.ID, &t.Title, &t.Done)
// 		todos = append(todos, t)
// 	}
// 	return todos, nil

// }

func SelectQueryAllTodo() ([]models.Todo, error) {
	rows, err := db.DB.Query(
		"SELECT id,title,done FROM todos",
	)
	if err != nil {
		fmt.Errorf("failed to fetch todos :%v", err)
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

func SelectQueryATodo(id int) (models.Todo, error) {
	var t models.Todo
	row := db.DB.QueryRow("SELECT id,title,done FROM todos WHERE id=$1", id)
	err := row.Scan(&t.ID, &t.Title, &t.Done)

	if err == sql.ErrNoRows {
		return models.Todo{}, fmt.Errorf("todo with id %d not found", id)
	}
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to fetch todo %d :%v", id, err)
	}
	return t, nil
}
