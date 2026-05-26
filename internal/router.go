package internal

import (
	"todo-app/internal/config"
	"todo-app/internal/handlers"
	"todo-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine, cfg *config.Config) {

	//public routes-anyone can hit these//
	r.POST("/register", handlers.HandleRegister)

	//closure bridges Gin's required func(c *gin.Context)signature//
	//with our handler that needs cfg injected as second parameter//
	r.POST("/login", func(c *gin.Context) {
		handlers.HandleLogin(c, cfg)
	})

	//protected routes- must have valid JWT//
	//protected ensuring each routes is been protected like a firewall//
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))
	{

		protected.POST("/createtodo", handlers.CreateTodo)
		protected.GET("/gettodo", handlers.GetATodo)
		protected.GET("/todo", handlers.GetAlltodos)
		protected.PUT("/todo/:id", handlers.UpdateTodo)
		protected.DELETE("/delete/:id", handlers.DeleteTodo)

	}

}
