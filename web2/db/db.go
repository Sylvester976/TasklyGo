package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	// Load .env only if it exists (for local development)
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file found but could not be loaded:", err)
		} else {
			log.Println(".env file loaded")
		}
	}

	// Get DATABASE_URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set. Make sure you have a .env file locally or set the env variable on Render")
	}

	// Connect to PostgreSQL
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	Pool = pool
	log.Println("Connected to PostgreSQL")

	// Auto-create roles table if it doesn't exist
	_, err = Pool.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS roles (
		id SERIAL PRIMARY KEY,
		role VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)
`)
	if err != nil {
		log.Fatal("Unable to create roles table:", err)
	}
	log.Println("Roles table checked/created")

	// Insert default roles if they don't exist
	_, err = Pool.Exec(context.Background(), `
	INSERT INTO roles (role)
	VALUES 
		('staff'),
		('manager')
	ON CONFLICT (role) DO NOTHING
`)
	if err != nil {
		log.Fatal("Unable to insert default roles:", err)
	}

	log.Println("Default roles checked/created")

	// Auto-create user table if it doesn't exist
	_, err = Pool.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		names VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,  -- added password column
		roles INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE, -- foreign key
		status BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)
`)
	if err != nil {
		log.Fatal("Unable to create USERS table:", err)
	}

	log.Println("USERS table checked/created")

	// Auto-create todos table if it doesn't exist
	_, err = Pool.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- foreign key
		status BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)
`)
	if err != nil {
		log.Fatal("Unable to create todos table:", err)
	}

	log.Println("Todos table checked/created")

}
