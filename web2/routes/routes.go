package routes

import (
	"net/http"
	"web2/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/signup", handlers.AuthRegisterHandler)
}
