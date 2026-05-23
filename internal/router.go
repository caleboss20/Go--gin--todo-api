package internal

import (
	"todo-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	r.POST("/createtodo", handlers.CreateTodo)
	r.GET("/gettodo", handlers.GetATodo)
	r.GET("/todo", handlers.GetAlltodos)
	r.PUT("/todo/:id", handlers.UpdateTodo)
	r.DELETE("/delete/:id", handlers.DeleteTodo)
}
