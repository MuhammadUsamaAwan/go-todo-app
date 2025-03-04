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

func (s *TodoService) GetAll(ctx context.Context) ([]TodoDTO, error) {
	return s.repo.GetAll(ctx)
}

func (s *TodoService) Create(ctx context.Context, dto CreateTodoDTO) error {
	return s.repo.Create(ctx, &dto)
}

func (s *TodoService) Update(ctx context.Context, dto UpdateTodoDTO, id int) error {
	return s.repo.Update(ctx, &dto, id)
}

func (s *TodoService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
