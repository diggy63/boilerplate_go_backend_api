package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/diggy63/boilerplate_go_api/cmd/api"
	"github.com/diggy63/boilerplate_go_api/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	DB_URL := os.Getenv("DB_HOST")
	// PORT := os.Getenv("PORT")
	PORT := ":8080"
	db, err := db.NewPSQLStorage(DB_URL) // Assuming this returns a *sql.DB and error
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	initStorage(db)
	server := api.NewAPIServer(PORT, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	log.Println("Connected to the database")

}
