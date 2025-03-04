package todo

import (
	"net/http"

	"github.com/MuhammadUsamaAwan/go-todo-app/internal/models"
	"github.com/MuhammadUsamaAwan/go-todo-app/pkg/response"
)

type TodoHandler struct {
	service *TodoService
}

func NewTodoHandler(service *TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAll(r.Context())
	if err != nil {
		response.JSONResponse(w, 500, models.APIResponse{
			Message: "Internal server error",
		})
		return
	}

	response.JSONResponse(w, 200, models.APIResponse{
		Data:    todos,
		Message: "Todos fetched successfully",
	})

}
