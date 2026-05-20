package handlers

import (
	"net/http"
	"todo-app/internal/models"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
	}
	var todo models.Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "todo deleted",
	})
}
