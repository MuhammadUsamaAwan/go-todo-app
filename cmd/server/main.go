package main

import (
	"log"
	"net/http"

	"github.com/MuhammadUsamaAwan/go-todo-app/internal/db"
	"github.com/MuhammadUsamaAwan/go-todo-app/internal/todo"
	"github.com/MuhammadUsamaAwan/go-todo-app/pkg/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.LoadConfig()

	pool := db.CreateDbPool(cfg.DbURL)
	defer pool.Close()

	todoRepo := todo.NewTodoRepository(pool)
	todoService := todo.NewTodoService(todoRepo)
	todoHandler := todo.NewTodoHandler(todoService)

	r := chi.NewRouter()
	r.Get("/todos", todoHandler.GetTodos)

	log.Printf("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
