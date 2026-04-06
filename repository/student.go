package repository

import (
	"database/sql"
	"errors"
	"strings"
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
	query := `SELECT id, first_name, last_name, email, age, course, created_at
	FROM students
	WHERE id = ?`

	var s models.Student
	err := db.DB.QueryRow(query, id).Scan(
		&s.ID,
		&s.FirstName,
		&s.LastName,
		&s.Age,
		&s.Course,
		&s.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err

	}
	return &s, nil

}

func Create(req models.CreateStudentRequest) (*models.Student, error) {
	query := `INSERT INTO students (first_name, last_name, email, age, course)
	VALUES (?,?,?,?,?)`

	result, err := db.DB.Exec(query,
		req.FirstName,
		req.LastName,
		req.Email,
		req.Age,
		req.Course,
	)

	if err != nil {
		if isUniqueConstraintErr(err) {
			return nil, ErrEmailExists
		}
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return GetByID(int(id))

}

func Update(id int, req models.UpdateStudentRequest) (*models.Student, error) {
	_, err := GetByID(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE students
			SET first_name = ?, last_name = ?, email = ?, age = ?, course = ?
			WHERE id = ?`

	_, err = db.DB.Exec(query,
		req.FirstName,
		req.LastName,
		req.Email,
		req.Age,
		req.Course,
		id,
	)

	if err != nil {
		if isUniqueConstraintErr(err) {
			return nil, ErrEmailExists
		}
		return nil, err
	}
	return GetByID(id)
}

func Delete(id int) error {
	_, err := GetByID(id)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec("DELETE FROM students WHERE id = ?", id)
	return err
}

func isUniqueConstraintErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed")
}

// func constainsString(s, substr string) bool {
// 	return len(s) >= len(substr) && (s == substr ||
// 		len(s) > 0 && constainsSubstring(s, substr))
// }

// func constainsSubstring(s, substr string) bool {
// 	for i := 0; i <= len(s)-len(substr); i++ {
// 		if s[i:i+len(substr)] == substr {
// 			return true
// 		}
// 	}
// 	return false
// }
