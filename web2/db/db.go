package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"web2/config"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	// Get DATABASE_URL from config
	dbURL := config.DatabaseURL
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set in config")
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

	// Auto-create users table
	_, err = Pool.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		names VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		roles INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
		status BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)
`)
	if err != nil {
		log.Fatal("Unable to create USERS table:", err)
	}
	log.Println("Users table checked/created")

	// Auto-create todos table
	_, err = Pool.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
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
