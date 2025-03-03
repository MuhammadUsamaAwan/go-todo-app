package main

import (
	"log"
	"time"

	"github.com/MuhammadUsamaAwan/go-todo-app/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config := config.LoadConfig()

	m, err := migrate.New(
		"file://migrations",
		config.DbURL)
	if err != nil {
		log.Fatalf("Could not initialize migrate: %v", err)
	}

	start := time.Now()
	err = m.Up()
	duration := time.Since(start)

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Printf("Migrations applied successfully in %s", duration)
}
