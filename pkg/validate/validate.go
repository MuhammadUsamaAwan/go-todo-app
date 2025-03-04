package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func Validate(r *http.Request, input any) error {
	json.NewDecoder(r.Body).Decode(input)
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(input)
}

func ParseURLParams(r *http.Request, key string) (int, error) {
	fmt.Println(chi.URLParam(r, "id"))
	value := chi.URLParam(r, key)
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("invalid path parameter")
	}
	return num, nil
}
