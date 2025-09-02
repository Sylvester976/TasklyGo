package models

import (
	"context"
	"time"

	"web/db"
)

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Fetch all todos
func GetTodos() ([]Todo, error) {
	rows, err := db.Pool.Query(context.Background(),
		"SELECT id, title, description, status, created_at, updated_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

// Create a new task
func CreateTodo(title, description string) (Todo, error) {
	var t Todo
	err := db.Pool.QueryRow(context.Background(),
		"INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING id, title, description, status, created_at, updated_at",
		title, description).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

// Update a task status/title/description
func UpdateTodo(id int, title, description string, status bool) (Todo, error) {
	var t Todo
	err := db.Pool.QueryRow(context.Background(),
		"UPDATE todos SET title=$1, description=$2, status=$3, updated_at=NOW() WHERE id=$4 RETURNING id, title, description, status, created_at, updated_at",
		title, description, status, id).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

// Delete a todo
func DeleteTodo(id int) error {
	_, err := db.Pool.Exec(context.Background(), "DELETE FROM todos WHERE id=$1", id)
	return err
}
