package main

import (
	"todo-app/internal"
	"todo-app/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()
	internal.SetUpRouter(r)
	r.Run(":5000")

}
