package todo

import (
	"log"
	"net/http"

	"github.com/MuhammadUsamaAwan/go-todo-app/internal/models"
	"github.com/MuhammadUsamaAwan/go-todo-app/pkg/response"
	"github.com/MuhammadUsamaAwan/go-todo-app/pkg/validate"
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
		response.JSONResponse(w, http.StatusInternalServerError, models.APIResponse{
			Message: "Internal server error",
		})
		return
	}

	response.JSONResponse(w, 200, models.APIResponse{
		Data:    todos,
		Message: "Resquest Successful",
	})

}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateTodoDTO
	err := validate.Validate(r, &dto)
	if err != nil {
		response.ValidationError(w, err)
		return
	}

	err = h.service.Create(r.Context(), dto)
	if err != nil {
		response.JSONResponse(w, http.StatusInternalServerError, models.APIResponse{
			Message: "Internal server error",
		})
		return
	}

	response.JSONResponse(w, http.StatusCreated, models.APIResponse{
		Message: "Todo created Successfully",
	})
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto UpdateTodoDTO
	err := validate.Validate(r, &dto)
	if err != nil {
		response.ValidationError(w, err)
		return
	}

	id, err := validate.ParseURLParams(r, "id")
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, models.APIResponse{
			Message: err.Error(),
		})
		return
	}

	err = h.service.Update(r.Context(), dto, id)
	if err != nil {
		log.Println(err)
		response.JSONResponse(w, http.StatusInternalServerError, models.APIResponse{
			Message: "Internal server error",
		})
		return
	}

	response.JSONResponse(w, http.StatusOK, models.APIResponse{
		Message: "Todo updated Successfully",
	})
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := validate.ParseURLParams(r, "id")
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, models.APIResponse{
			Message: err.Error(),
		})
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		response.JSONResponse(w, http.StatusInternalServerError, models.APIResponse{
			Message: "Internal server error",
		})
		return
	}

	response.JSONResponse(w, http.StatusOK, models.APIResponse{
		Message: "Todo deleted Successfully",
	})
}
