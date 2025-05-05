package structs

type SuccessResponse struct {
	Success bool   `json:"success"`
	Massage string `json:"message"`
	Data    any    `json:"data"`
}
