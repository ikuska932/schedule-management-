package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" 
)

func InitDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./schedule.db")
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
