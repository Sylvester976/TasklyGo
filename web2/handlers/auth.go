package handlers

import (
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"web2/models"
	"web2/utils"
)

type RegisterData struct {
	Roles []models.Role
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	roles, err := models.GetAllRoles(r.Context())
	if err != nil {
		http.Error(w, "Failed to load roles", http.StatusInternalServerError)
		return
	}

	data := RegisterData{
		Roles: roles,
	}

	tmpl, err := template.ParseFiles("./templates/register.html")
	if err != nil {
		log.Println("Template parse error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Template execute error:", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
		return
	}
}

func AuthRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Log the incorrect method
		log.Printf("Invalid method %s on %s", r.Method, r.URL.Path)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Parse form
	if err := r.ParseForm(); err != nil {
		// Log the incorrect method
		log.Printf("Invalid method %s on %s", r.Method, r.URL.Path)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	firstName := r.FormValue("firstName")
	surname := r.FormValue("surname")
	names := firstName + " " + surname
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")
	roleStr := r.FormValue("role")
	role, _ := strconv.Atoi(roleStr)

	// Check password strength
	if ok, msg := utils.CheckPasswordStrength(password); !ok {
		log.Printf(msg, r.Method, r.URL.Path)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Check if passwords match
	if confirmPassword != password {
		log.Printf("Passwords do not match on %s", r.URL.Path)
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Hash password
	hash, _ := utils.HashPassword(password)

	// Create user
	user := &models.User{
		Names:    names,
		Email:    email,
		Password: hash,
		Roles:    role,
		Status:   true, // default active
	}

	if err := user.Create(); err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	log.Printf("User %s registered with ID %d", user.Email, user.ID)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/login.html")
}

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Invalid method %s on %s", r.Method, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// Parse form
	if err := r.ParseForm(); err != nil {
		log.Printf("Invalid method %s on %s", r.Method, r.URL.Path)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validate credentials
	user, err := models.GetUserByEmailAndPassword(r.Context(), email, password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Create a session
	session, _ := store.Get(r, "session-name")

	// Save user info in session
	session.Values["userID"] = user.ID
	session.Values["userName"] = user.Names
	session.Values["userEmail"] = user.Email

	// Save session to the response
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Could not save session", http.StatusInternalServerError)
		return
	}

	// Redirect or respond
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
