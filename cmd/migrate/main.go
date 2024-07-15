package main

import (
	"log"
	"os"

	"github.com/diggy63/boilerplate_go_api/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	DB_URL := os.Getenv("DB_HOST")
	//make a db connection
	db, err := db.NewPSQLStorage(DB_URL) // Assuming this returns a *sql.DB and error
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	//create a database driver for migrate
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Error creating database driver instance:", err)
	}

	//create a migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("Error creating migration instance:", err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Error running migration:", err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
