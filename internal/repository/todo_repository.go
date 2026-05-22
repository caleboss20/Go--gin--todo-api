package repository

import (
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
		fmt.Errorf("failed to fetch todos %v", err)
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
