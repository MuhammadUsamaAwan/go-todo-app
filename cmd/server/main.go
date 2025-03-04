package main

import (
	"log"
	"net/http"

	infrastructure "github.com/MuhammadUsamaAwan/go-todo-app/internal/infrastruture"
	"github.com/MuhammadUsamaAwan/go-todo-app/internal/todo"
	"github.com/MuhammadUsamaAwan/go-todo-app/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	cfg := config.LoadConfig()

	pool := infrastructure.CreateDbPool(cfg.DbURL)
	defer pool.Close()

	todoRepo := todo.NewTodoRepository(pool)
	todoService := todo.NewTodoService(todoRepo)
	todoHandler := todo.NewTodoHandler(todoService)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	r.Get("/todos", todoHandler.GetAll)
	r.Post("/todos", todoHandler.Create)
	r.Put("/todos/{id}", todoHandler.Update)
	r.Delete("/todos/{id}", todoHandler.Delete)

	log.Printf("Starting server on " + cfg.Port)
	http.ListenAndServe(":"+cfg.Port, r)
}
