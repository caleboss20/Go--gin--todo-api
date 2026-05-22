package handlers

import (
	"net/http"
	"todo-app/internal/models"
	"todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

var todos []models.Todo

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title is required",
		})
		return
	}
	todos = append(todos, todo)

	//passing an argument title to the create todo repo so title is added to database via repo//
	err = repository.CreatingTodoQuery(todo.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}
