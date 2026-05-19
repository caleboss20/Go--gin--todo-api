package handlers

import (
	"fmt"
	"net/http"
	"todo-app/internal/models"

	"github.com/gin-gonic/gin"
)

var todos []models.Todo
var idCounter int = 1

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
	todo.ID = idCounter
	fmt.Println("ID", todo.ID)
	idCounter++

	todos = append(todos, todo)
	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
		// "id":   todo.ID,
	})
}
