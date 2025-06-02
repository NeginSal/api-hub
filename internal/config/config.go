package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JwtSecret []byte

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using default secret")
		JwtSecret = []byte("default_secret")
	} else {
		JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	}
}
