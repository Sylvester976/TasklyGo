package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web2/models"
	"web2/session"
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form submission", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := models.GetUserByEmailAndPassword(r.Context(), email, password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Create or retrieve session
	sess, err := session.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Could not create session", http.StatusInternalServerError)
		return
	}

	// Save user info in session
	sess.Values["userID"] = user.ID
	sess.Values["userName"] = user.Names
	sess.Values["userEmail"] = user.Email

	// Commit session
	if err := sess.Save(r, w); err != nil {
		log.Printf("Session save error: %v", err)
		http.Error(w, "Could not save session", http.StatusInternalServerError)
		return
	}

	// Redirect based on role
	switch user.Roles {
	case 1:
		http.Redirect(w, r, "/task", http.StatusSeeOther)
	case 2:
		http.Redirect(w, r, "/manager", http.StatusSeeOther)
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func staffTaskHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := session.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Could not retrieve session", http.StatusInternalServerError)
		return
	}

	//userID, ok := sess.Values["userID"].(int)
	//if !ok {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	//tasks, err := models.GetTasksByUserID(r.Context(), userID)
	//if err != nil {
	//	http.Error(w, "Could not retrieve tasks", http.StatusInternalServerError)
	//	return
	//}

	tmpl, err := template.ParseFiles("./templates/staff_tasks.html")
	if err != nil {
		log.Println("Template parse error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		UserName string
	}{
		UserName: sess.Values["userName"].(string),
	}

	//data := struct {
	//	UserName string
	//	Tasks    []models.Task
	//}{
	//	UserName: sess.Values["userName"].(string),
	//	Tasks:    tasks,
	//}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Template execute error:", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
		return
	}
}

func managerTaskHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := session.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Could not retrieve session", http.StatusInternalServerError)
		return
	}

	//userID, ok := sess.Values["userID"].(int)
	//if !ok {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	//tasks, err := models.GetAllTasks(r.Context())
	//if err != nil {
	//	http.Error(w, "Could not retrieve tasks", http.StatusInternalServerError)
	//	return
	//}

	tmpl, err := template.ParseFiles("./templates/manager_tasks.html")
	if err != nil {
		log.Println("Template parse error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		UserName string
	}{
		UserName: sess.Values["userName"].(string),
	}

	//data := struct {
	//	UserName string
	//	Tasks    []models.Task
	//}{
	//	UserName: sess.Values["userName"].(string),
	//	Tasks:    tasks,
	//}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Template execute error:", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
		return
	}
}
