package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Assuming cfg is a string containing the DSN for PostgreSQL
func NewPSQLStorage(cfg string) (*sql.DB, error) {
	// Use "postgres" as the driver name and cfg as the DSN
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err) // Consider changing this to return the error instead of terminating the program
	}

	return db, nil

}
