package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type apiResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func HttpResponse(w http.ResponseWriter, status int, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(apiResponse{
		Data:    data,
		Message: message,
	})
}

func HttpError(w http.ResponseWriter, status int, message string) {
	HttpResponse(w, status, nil, message)
}

func Validate(r *http.Request, input any) error {
	json.NewDecoder(r.Body).Decode(input)
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(input)
}

func ParseURLParam(r *http.Request, key string) (int, error) {
	value := chi.URLParam(r, key)
	return strconv.Atoi(value)
}
