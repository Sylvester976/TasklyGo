package main

import (
	"log"
	"net/http"
	"web2/config"
	"web2/db"
	"web2/routes"
)

func main() {
	// Load environment variables
	config.Load()

	// Connect to DB
	db.ConnectDB()
	defer db.Pool.Close()

	// Use port from config (fallback handled in config.Load)
	port := config.Port

	// Setup routes
	routes.SetupRoutes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
