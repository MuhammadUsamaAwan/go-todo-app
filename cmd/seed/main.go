package main

import (
	"context"
	"log"
	"time"

	"github.com/MuhammadUsamaAwan/go-todo-app/internal/db"
	"github.com/MuhammadUsamaAwan/go-todo-app/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	ctx := context.Background()
	pool := db.CreateDbPool(cfg.DbURL)
	defer pool.Close()

	seeds := []string{
		`INSERT INTO todos (title, completed) VALUES ('Learn go', true) ON CONFLICT DO NOTHING;`,
		`INSERT INTO todos (title, completed) VALUES ('Build a project', false) ON CONFLICT DO NOTHING;`,
	}

	start := time.Now()

	for _, query := range seeds {
		_, err := pool.Exec(ctx, query)
		if err != nil {
			log.Panicf("Error while seeding data: %v", err)
		}
	}

	duration := time.Since(start)

	log.Printf("Seeding data successfully in %s", duration)
}
