package data

type ErrorResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}
