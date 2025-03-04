package todo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository struct {
	pool *pgxpool.Pool
}

func NewTodoRepository(pool *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{pool: pool}
}

func (r *TodoRepository) GetAll(ctx context.Context) ([]TodoDTO, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []TodoDTO
	for rows.Next() {
		var todo TodoDTO
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) Create(ctx context.Context, dto *CreateTodoDTO) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO todos (title, completed) VALUES ($1, $2)", dto.Title, dto.Completed)
	return err
}

func (r *TodoRepository) Update(ctx context.Context, dto *UpdateTodoDTO, id int) error {
	_, err := r.pool.Exec(ctx, "UPDATE todos SET title = $1, completed = $2 WHERE id = $3", dto.Title, dto.Completed, id)
	return err
}

func (r *TodoRepository) Delete(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM todos WHERE id = $1", id)
	return err
}
