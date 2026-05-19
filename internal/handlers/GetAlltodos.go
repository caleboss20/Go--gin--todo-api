package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlltodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}
