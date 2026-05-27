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

//*Config means this function returns a pointer to Config-a memory address,not a copy.

func Load() *Config {
	//read the .env file and load all variables into memory//
	//if .env file is missing crash immediately with clear message//
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	//create a new Config struct and fill each field
	//with its matching value from the .env file
	//& means return the memory address not a copy//

	return &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		DBUrl:     MustGetEnv("DB_URL"),
		SECRETKEY: MustGetEnv("JWT_SECRET"),
	}

}

//optional-returns default value if missing//

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// mandatory-crashes app immediately if missing//
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("required envirinment variable missing :" + key)
	}
	return value
}
