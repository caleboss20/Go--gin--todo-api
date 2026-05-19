package main

import (
	"todo-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/createtodo", handlers.CreateTodo)
	r.GET("/todo", handlers.GetAlltodos)
	// r.GET("/todo/:id", GetATodo)
	// r.PUT("/todo/:id", UpdateTodo)
	// r.DELETE("/delete/:id", DeleteATodo)

	r.Run(":5000")

}
