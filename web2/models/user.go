package models

import (
	"context"
	"time"

	"web2/db"
)

type User struct {
	ID        int
	Names     string
	Email     string
	Password  string
	Roles     int
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Insert user into DB
func (u *User) Create() error {
	query := `
        INSERT INTO users (names, email, password, roles, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `
	return db.Pool.QueryRow(
		context.Background(),
		query,
		u.Names,
		u.Email,
		u.Password,
		u.Roles,
		u.Status,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}
