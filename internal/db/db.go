package db

import (
	"database/sql"
	"log"
)

// create a global database connection//
var DB *sql.DB

func Connect() {
	var err error
	dsn := "postgres://postgres:password@localhost:5432/todoapp?sslmode=disable"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to open database")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB not reachable")
	}
	log.Println("connection to database successful")

}
