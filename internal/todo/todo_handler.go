package todo

import (
	"net/http"

	"github.com/MuhammadUsamaAwan/go-todo-app/pkg/utils"
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
		utils.HttpError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.HttpResponse(w, http.StatusOK, todos, "Todos fetched Successfully")
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateTodoDTO
	err := utils.Validate(r, &dto)
	if err != nil {
		utils.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Create(r.Context(), dto)
	if err != nil {
		utils.HttpError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.HttpResponse(w, http.StatusCreated, nil, "Todo created Successfully")
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto UpdateTodoDTO
	err := utils.Validate(r, &dto)
	if err != nil {
		utils.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := utils.ParseURLParam(r, "id")
	if err != nil {
		utils.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Update(r.Context(), dto, id)
	if err != nil {
		utils.HttpError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.HttpResponse(w, http.StatusOK, nil, "Todo updated Successfully")
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseURLParam(r, "id")
	if err != nil {
		utils.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		utils.HttpError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.HttpResponse(w, http.StatusOK, nil, "Todo deleted Successfully")
}
