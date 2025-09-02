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
	http.HandleFunc("/task", handlers.StaffTaskHandler)
	http.HandleFunc("/manager", handlers.ManagerTaskHandler)
	http.HandleFunc("/logout", handlers.AuthLogoutHandler)

}
