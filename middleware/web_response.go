package middleware

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func OkResponse() StatusResponse {
	return StatusResponse{Status: "ok"}
}
