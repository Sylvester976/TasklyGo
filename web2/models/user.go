package models

import (
	"context"
	"errors"
	"time"

	"web2/db"
	"web2/utils"
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

func GetUserByEmailAndPassword(ctx context.Context, email, plainPassword string) (*User, error) {
	var u User

	query := `
		SELECT id, names, email, password, roles, status, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	err := db.Pool.QueryRow(ctx, query, email).Scan(
		&u.ID,
		&u.Names,
		&u.Email,
		&u.Password, // stored hashed password
		&u.Roles,
		&u.Status,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// compare stored hash with entered plain password
	if !utils.CheckPasswordHash(plainPassword, u.Password) {
		return nil, errors.New("invalid credentials")
	}

	return &u, nil
}
