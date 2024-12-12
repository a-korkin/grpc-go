package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file: %s", err)
	}
}

func GetEnv(variable string) string {
	return os.Getenv(variable)
}
