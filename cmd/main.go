package main

import (
	"todo-app/internal/db"
	"todo-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()
	r.POST("/createtodo", handlers.CreateTodo)
	r.GET("/todo", handlers.GetAlltodos)
	r.PUT("/todo/:id", handlers.UpdateTodo)
	r.DELETE("/delete/:id", handlers.DeleteTodo)

	r.Run(":5000")

}
