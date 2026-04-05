package main

import (
	"fmt"
	"log"
	"net/http"
	"students-api/db"
)

func main() {
	fmt.Println("Starting Student API server...")

	if err := db.Init("students.db"); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	defer db.Close() // close cleanly when the program exits

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(": 8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
