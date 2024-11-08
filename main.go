package main

import (
	"log"
	"net/http"
	"schedule-management/pkg/database"
	"schedule-management/pkg/handlers"
	
)

func main() {
	
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS teachers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			qualification TEXT
		)
	`)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы:", err)
	}

	http.HandleFunc("/teachers", handlers.TeachersHandler(db))

	log.Println("Сервер запущен на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
