package todo

import (
	"encoding/json"
	"net/http"
)

type TodoHandler struct {
	service *TodoService
}

func NewTodoHandler(service *TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetTodos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
