package main

import (
	"log"
	"net/http"
	"os"
	"web2/db"
	"web2/routes"
)

func main() {
	// Connect to DB
	db.ConnectDB()
	defer db.Pool.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	// Setup routes
	routes.SetupRoutes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
