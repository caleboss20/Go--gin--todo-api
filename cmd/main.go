package main

import (
	"fmt"
	"todo-app/internal"
	"todo-app/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.Connect()
	r := gin.Default()
	internal.SetUpRouter(r)
	fmt.Println("port sering on:5000")
	r.Run(":5000")

}
