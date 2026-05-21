package db

import (
	"database/sql"
	"log"
)

// var DB *sql.DB

// func Connect() {
// 	var err error
// 	//DSN==data sourcce name,how to find and connect to the database//s
// 	dsn := "postgres://caleboss:__67myboss$12@localhost:5432/todoapp?sslmode=disable"
// 	DB, err = sql.Open("pgx", dsn)
// 	if err != nil {
// 		log.Fatal("Database failed to open")

// 	}

// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatal("database failed to connect")

// 	}

// 	log.Println("database connected successfully")

// }
var DB *sql.DB

func Connect() {
	var err error
	dsn := "postgres://caleboss:password@localhost:5432/todoapp?sslmode=disable"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("database failed to open")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("database failed to connect")
	}
	log.Println("database connected successfully")
}
