package structs

type ErrorResponse struct {
	Success bool              `json:"success"`
	Massage string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}
