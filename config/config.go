package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("can't read .env file. Err: %s", err)
	}
}
