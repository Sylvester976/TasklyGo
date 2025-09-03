package handlers

import (
	"html/template"
	"log"
	"net/http"
	"web2/models"
	"web2/session"
)

func StaffTaskHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := session.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Could not retrieve session", http.StatusInternalServerError)
		return
	}

	userID, ok := sess.Values["userID"].(int)
	if !ok || userID < 1 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//modify later to work with alerts
	userRoles, ok := sess.Values["userRoles"].(int)
	if !ok || userRoles != 1 {
		http.Redirect(w, r, "/manager", http.StatusSeeOther)
		return
	}

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

func ManagerTaskHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := session.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Could not retrieve session", http.StatusInternalServerError)
		return
	}

	//modify later to work with alerts
	userID, ok := sess.Values["userID"].(int)
	if !ok || userID < 1 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userRoles, ok := sess.Values["userRoles"].(int)
	if !ok || userRoles != 2 {
		http.Redirect(w, r, "/task", http.StatusSeeOther)
		return
	}

	allUsers, _ := models.GetAllUsers(r.Context())
	all_tasks, _ := models.GetAllTasks(r.Context())

	tmpl, err := template.ParseFiles("./templates/manager_tasks.html")
	if err != nil {
		log.Println("Template parse error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		UserName string
		AllUsers interface{}
		AllTasks interface{}
		userID   int
	}{
		UserName: sess.Values["userName"].(string),
		AllUsers: allUsers,
		AllTasks: all_tasks,
		userID:   userID,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Template execute error:", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
		return
	}

}
