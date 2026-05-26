package internal

import (
	"todo-app/internal/handlers"
	"todo-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {

	//public routes-anyone can hit these//
	r.POST("/register", handlers.HandleRegister)
	r.POST("/login", handlers.HandleLogin)

	//protected routes- must have valid JWT//
	//protected ensuring each routes is been protected like a firewall//
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{

		protected.POST("/createtodo", handlers.CreateTodo)
		protected.GET("/gettodo", handlers.GetATodo)
		protected.GET("/todo", handlers.GetAlltodos)
		protected.PUT("/todo/:id", handlers.UpdateTodo)
		protected.DELETE("/delete/:id", handlers.DeleteTodo)

	}

}
