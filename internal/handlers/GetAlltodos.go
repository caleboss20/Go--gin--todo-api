package handlers

import (
	"net/http"
	"todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetAlltodos(c *gin.Context) {
	//call repository for querying database list of todos//
	todos, err := repository.SelectQueryAllTodo()
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
