package main

import (
	"log"
	"net/http"
	"os"

	"web/db"
	"web/handlers"
)

func main() {
	// Connect to DB
	db.ConnectDB()
	defer db.Pool.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/todos", handlers.GetTodosHandler)
	http.HandleFunc("/todos/create", handlers.CreateTodoHandler)
	http.HandleFunc("/todos/update/", handlers.UpdateTodoHandler) // expects /todos/update/{id}
	http.HandleFunc("/todos/delete/", handlers.DeleteTodoHandler) // expects /todos/delete/{id}

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
