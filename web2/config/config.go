package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseURL string
	Port        string
	HashKey     []byte
	BlockKey    []byte
	SessionLife int
)

func Load() {
	// Load .env if present
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file found but could not be loaded:", err)
		}
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	Port = os.Getenv("PORT")

	// Load session keys
	hash := os.Getenv("SESSION_HASH_KEY")
	block := os.Getenv("SESSION_BLOCK_KEY")
	if hash == "" {
		log.Fatal("SESSION_HASH_KEY is required")
	}
	if block == "" {
		log.Fatal("SESSION_BLOCK_KEY is required")
	}
	HashKey = []byte(hash)
	BlockKey = []byte(block)

	// Session lifetime
	life, err := strconv.Atoi(os.Getenv("SESSION_LIFETIME"))
	if err != nil {
		life = 3600 // default in seconds
	}
	SessionLife = life

	if DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
}
