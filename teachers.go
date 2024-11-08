package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"schedule-management/pkg/models"
)

func TeachersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		teachers, err := models.GetAllTeachers(db)
		if err != nil {
			http.Error(w, "Ошибка при получении преподавателей", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("web/templates/teachers.html"))
		tmpl.Execute(w, teachers)
	}
}
