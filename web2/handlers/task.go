package handlers

import (
	"html/template"
	"log"
	"net/http"
	"web2/session"
)

func StaffTaskHandler(w http.ResponseWriter, r *http.Request) {
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

func ManagerTaskHandler(w http.ResponseWriter, r *http.Request) {
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
