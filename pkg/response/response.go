package response

import (
	"encoding/json"
	"net/http"

	"github.com/MuhammadUsamaAwan/go-todo-app/internal/models"
	"github.com/go-playground/validator/v10"
)

func JSONResponse(w http.ResponseWriter, statusCode int, response models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func ValidationError(w http.ResponseWriter, err error) {
	var errorMessage string
	errors := err.(validator.ValidationErrors)
	for i, e := range errors {
		if i > 0 {
			errorMessage += ", "
		}
		errorMessage += e.Field() + " is " + e.Tag()
	}
	JSONResponse(w, http.StatusBadRequest, models.APIResponse{
		Message: errorMessage,
	})
}
