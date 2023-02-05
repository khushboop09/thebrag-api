package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvMongoURI() string {
	if os.Getenv("MONGOURI") == "" {
		loadEnv()
	}
	return os.Getenv("MONGOURI")
}
