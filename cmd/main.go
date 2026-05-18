package main

import (
	"todo-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/createtodo", handlers.CreateTodo)
	r.Run(":8080")

}
