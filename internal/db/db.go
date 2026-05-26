package db

import (
	"database/sql"
	"log"
	"todo-app/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect(cfg *config.Config) {
	var err error
	//use cfg.DBurl instead of os.Getenv("DB_URL")
	DB, err := sql.Open("pgx", cfg.DBUrl)
	if err != nil {
		log.Fatalf("failed to open database %v", err)
	}

	//ping to confirm connection is actually alive//

	err = DB.Ping()
	if err != nil {
		log.Fatalf("failed to connect to database %v", err)
	}
}
