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
	//add the rate limiting middleware to the router before the route runs on a request//
	r.POST("/login", middleware.FixedWindowRateLimit(), func(c *gin.Context) {
		handlers.HandleLogin(c, cfg)

	})

	//protected routes- must have valid JWT//
	//protected ensuring each routes is been protected like a firewall//

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))

	{
		protected.GET("/api/todos", handlers.GetAlltodos)
		protected.GET("/api/todo/:id", handlers.GetATodo)
		protected.POST("/api/todos", middleware.LessStrictLimiter(), handlers.CreateTodo)
		protected.PUT("/api/todo/:id", handlers.UpdateTodo)
		protected.DELETE("api/todo/:id", handlers.DeleteTodo)

	}

}
