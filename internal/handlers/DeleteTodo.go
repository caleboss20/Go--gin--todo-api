package handlers

import (
	"net/http"
	"strconv"
	"todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}
	//converting the id to string//
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	//==calling the repository delete query==//
	err = repository.DeleteTodoQuery(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo deleted",
	})
}
