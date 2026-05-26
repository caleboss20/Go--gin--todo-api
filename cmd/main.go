package main

import (
	"fmt"
	"todo-app/internal"
	"todo-app/internal/config"
	"todo-app/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {

	//load all environment variables first//
	//everything else depends on this//
	cfg := config.Load()

	//pass config to database connection//
	db.Connect(cfg)

	r := gin.Default()

	//fetch all the endpoints from router.go//
	internal.SetUpRouter(r, cfg)
	fmt.Println("server running on port:" + cfg.AppPort)

	//use prt from config//
	r.Run(":" + cfg.AppPort)

}
