package db

import (
	"database/sql"
	"log"
)

//	func Connect() {
//		var err error
//		dsn := "postgres://username:password@localhost:5432/todoapp?sslmode=disable"
//		DB, err = sql.Open("pgx", dsn)
//		if err != nil {
//			log.Fatalf("database failed to open: %s, %v", dsn, err)
//		}
//		err = DB.Ping()
//		if err != nil {
//			log.Fatalf("failed to reach database at: %s, %v", dsn, err)
//		}
//		fmt.Println("Database connected successfully!")
//	}
var DB *sql.DB

func Connect() {
	var err error
	dsn := "postgres://username:password@localhost:5432/todoapp?sslmode=disable"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open database %v", err)

	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("failed to reach database %v", err)

	}

}
