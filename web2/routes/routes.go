package routes

import (
	"net/http"
	"web2/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/signup", handlers.AuthRegisterHandler)
	http.HandleFunc("/login", handlers.AuthLoginHandler)
	http.HandleFunc("/task", handlers.staffTaskHandler)
	http.HandleFunc("/manager", handlers.managerTaskHandler)

}
