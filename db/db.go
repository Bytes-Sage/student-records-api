package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(filepath string) error {
	var err error

	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return err
	}

	log.Println("Database connected successfully")

	return createTables()
}

func createTables() error {
	query := `
	
	CREATE TABLE IF NOT EXISTS students (
	id  INTEGER PRIMARY KEY AUTOINCREMENT,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	age INTEGER NOT NULL,
	course TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Tables ready")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
