package handlers

import (
	"net/http"
	"todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetAlltodos(c *gin.Context) {
	//Authorization//
	//Retrieve authenticated user ID from JWT middleware context//

	//to check is id is actually passed in the middleware//
	value, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	//if the id then exist in middleware then convert to int//
	//type assertion used here//
	id, ok := value.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id type in context",
		})
		return
	}

	//call repository for querying database list of todos//
	todos, err := repository.SelectQueryAllTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch todos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}
