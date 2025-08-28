package config

import (
	"log"
	"os"
	"strconv"
)

var (
	DatabaseURL string
	Port        string
	SecretKey   string
	SessionLife int
)

func Load() {
	DatabaseURL = os.Getenv("DATABASE_URL")
	Port = os.Getenv("PORT")
	SecretKey = os.Getenv("SECRET_KEY")

	life, err := strconv.Atoi(os.Getenv("SESSION_LIFETIME"))
	if err != nil {
		life = 3600 // default
	}
	SessionLife = life

	if DatabaseURL == "" || SecretKey == "" {
		log.Fatal("Required env vars missing")
	}
}
