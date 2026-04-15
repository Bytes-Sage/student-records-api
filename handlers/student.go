package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"students-api/models"
	"students-api/repository"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, models.APIResponse{
		Success: false,
		Message: message,
	})
}

func parseID(r *http.Request) (int, error) {
	//URL: /api/students/3
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		return 0, errors.New("missing ID in Url")
	}
	return strconv.Atoi(parts[len(parts)-1])
}

// get all
func GetAll(w http.ResponseWriter, r *http.Request) {
	students, err := repository.GetAll()

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch students ")
		return
	}

	writeJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    students,
	})
}

// Get Id
func GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	student, err := repository.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			writeError(w, http.StatusNotFound, "Student not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to fetch student")

	}
	writeJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    student,
	})
}

// create
func Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateStudentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validateCreateRequest(req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	student, err := repository.Create(req)
	if err != nil {
		if errors.Is(err, repository.ErrEmailExists) {
			writeError(w, http.StatusConflict, "Email already exists")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to create student")
		return
	}
	writeJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Data:    student,
		Message: "Student create successfully",
	})

}

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	var req models.UpdateStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validateUpdateRequest(req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	student, err := repository.Update(id, req)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			writeError(w, http.StatusNotFound, "Student not found")
			return
		}
		if errors.Is(err, repository.ErrEmailExists) {
			writeError(w, http.StatusConflict, "Email already exists")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to update student")
		return
	}

	writeJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Student updated successfully",
		Data:    student,
	})
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	err = repository.Delete(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			writeError(w, http.StatusNotFound, "Student not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "Failed to delete student")
		return
	}

	writeJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Student deleted successfully",
	})
}

//validations

func validateCreateRequest(req models.CreateStudentRequest) error {
	if strings.TrimSpace(req.FirstName) == "" {
		return errors.New("first_name is required")
	}
	if strings.TrimSpace(req.LastName) == "" {
		return errors.New("last_name is required")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if !strings.Contains(req.Email, "@") {
		return errors.New("email is invalid")
	}
	if req.Age <= 0 || req.Age > 70 {
		return errors.New("age must be between 1 and 70")
	}
	if strings.TrimSpace(req.Course) == "" {
		return errors.New("course is required")
	}
	return nil
}

func validateUpdateRequest(req models.UpdateStudentRequest) error {
	return validateCreateRequest(models.CreateStudentRequest(req))
}
