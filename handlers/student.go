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
