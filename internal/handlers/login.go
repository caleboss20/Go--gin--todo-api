package handlers

import (
	"net/http"
	"todo-app/internal/models"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	var loginInput models.LoginInput

	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//call login service to handle logic
	token, err := services.LoginUser(loginInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//return response to user//
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
