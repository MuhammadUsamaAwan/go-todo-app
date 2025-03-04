package todo

type TodoDTO struct {
	ID        *string `json:"id"`
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

type CreateTodoDTO struct {
	Title     *string `json:"title" validate:"required"`
	Completed *bool   `json:"completed" validate:"required"`
}

type UpdateTodoDTO struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}
