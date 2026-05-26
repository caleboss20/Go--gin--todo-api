package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// config holds all environment variables for the app
// instead calling os.Getenv scattered everywhere//
type Config struct {
	AppPort   string
	DBUrl     string
	SECRETKEY string
}

func Load() *Config {
	//read the .env file and load all variables into memory//
	//if .env file is missing crash immediately with clear message//

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	//create a new Config struct and fill each field
	//with its matching alue from the .env file
	//& means return the memory address not a copy//

	return &Config{
		AppPort:   os.Getenv("APP_PORT"),
		DBUrl:     os.Getenv("DB_URL"),
		SECRETKEY: os.Getenv("JWT_SECRET"),
	}

}
