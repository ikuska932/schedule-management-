package models

import "database/sql"

type Teacher struct {
	ID          int
	Name        string
	Qualification string
}

func GetAllTeachers(db *sql.DB) ([]Teacher, error) {
	rows, err := db.Query("SELECT id, name, qualification FROM teachers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []Teacher
	for rows.Next() {
		var teacher Teacher
		if err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Qualification); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}
