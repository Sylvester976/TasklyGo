package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"web/models"
)

// Get all todos
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// Create a new todo
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	todo, err := models.CreateTodo(input.Title, input.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// Update a todo
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/update/")
	id, _ := strconv.Atoi(idStr)

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      bool   `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	todo, err := models.UpdateTodo(id, input.Title, input.Description, input.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// Delete a todo
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/delete/")
	id, _ := strconv.Atoi(idStr)

	err := models.DeleteTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
