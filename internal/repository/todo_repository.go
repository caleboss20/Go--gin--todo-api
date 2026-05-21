package repository

import "todo-app/internal/db"

//to query the database==for inserting or adding a todo to the database//
func QueryCreatingTodo(title string) error {
	_, err := db.DB.Exec(
		"INSERT INTO todos(title)VALUES($1)", title,
	)
	return err

}

//to update a single todo//
func UpdatingTodo(title string, id int) error {
	_, err := db.DB.Exec(
		"UPDATE todos SET title=$1 WHERE id=$2", title, id,
	)
	return err
}

func DeletingTodo(id int) error {
	_, err := db.DB.Exec(
		"DELETE FROM todos WHERE id=$1", id,
	)
	return err
}
