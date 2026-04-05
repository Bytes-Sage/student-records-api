package repository

import (
	"errors"
	"students-api/db"
	"students-api/models"
	// "students-api/models"
)

var (
	ErrNotFound    = errors.New("student not found")
	ErrEmailExists = errors.New("email already exists")
)

func GetAll() ([]models.Student, error) {
	query := `SELECT id, first_name, last_name, email, age, course, created_at 
	FROM students
	ORDER BY created_at DESC`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []models.Student{}
	for rows.Next() {
		var s models.Student
		err := rows.Scan(
			&s.ID,
			&s.FirstName,
			&s.LastName,
			&s.Email,
			&s.Age,
			&s.Course,
			&s.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

func GetByID(id int) (*models.Student, error) {
	query := `SELECT id, first_name`
}
