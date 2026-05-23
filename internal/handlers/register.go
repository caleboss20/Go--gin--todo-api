package handlers

import (
	"net/http"
	"todo-app/internal/models"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}
	//for validation//
	err = services.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//send response to the user//
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
