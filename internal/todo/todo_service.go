package todo

import (
	"context"
)

type TodoService struct {
	repo *TodoRepository
}

func NewTodoService(repo *TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetTodos(ctx context.Context) ([]Todo, error) {
	return s.repo.GetTodos(ctx)
}
