package handlers

import (
	"net/http"
	"strconv"
	"todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetATodo(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id ",
		})
		return
	}

	todo, err := repository.SelectQueryATodo(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})

}
