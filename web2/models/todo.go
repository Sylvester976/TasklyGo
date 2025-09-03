package models

import (
	"context"
	"time"

	"web2/db"
)

type Task struct {
	ID          int
	Title       string
	Description string
	User_id     string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskWithUser struct {
	ID          int
	Title       string
	Description string
	UserID      int
	UserName    string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Insert task into DB
func (t *Task) Create() error {
	query := `
        INSERT INTO todos (title, description, user_id, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `
	return db.Pool.QueryRow(
		context.Background(),
		query,
		t.Title,
		t.Description,
		t.User_id,
		t.Status,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func GetAllTasks(ctx context.Context) ([]Task, error) {
	var tasks []Task

	query := `
		SELECT id, title, description, user_id, status, created_at, updated_at
		FROM todos
		ORDER BY created_at DESC
	`
	rows, err := db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.User_id,
			&t.Status,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}

func GetTasksByUserID(ctx context.Context, userID int) ([]Task, error) {
	var tasks []Task

	query := `
		SELECT id, title, description, user_id, status, created_at, updated_at
		FROM todos
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	rows, err := db.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.User_id,
			&t.Status,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetAllTasksWithUsers(ctx context.Context) ([]TaskWithUser, error) {
	var tasks []TaskWithUser

	query := `
		SELECT t.id, t.title, t.description, t.user_id, u.names, t.status, t.created_at, t.updated_at
		FROM todos t
		JOIN users u ON t.user_id = u.id
		ORDER BY t.created_at DESC
	`
	rows, err := db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t TaskWithUser
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.UserID,
			&t.UserName,
			&t.Status,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
