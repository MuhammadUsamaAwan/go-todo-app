package response

import (
	"encoding/json"
	"net/http"

	"github.com/MuhammadUsamaAwan/go-todo-app/internal/models"
)

func JSONResponse(w http.ResponseWriter, statusCode int, response models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
