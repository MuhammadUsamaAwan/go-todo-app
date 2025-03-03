package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDbPool(dbURL string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return pool
}
