package utils

type ResponseData struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
