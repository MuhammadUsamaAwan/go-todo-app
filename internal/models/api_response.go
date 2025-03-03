package models

type APIResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
