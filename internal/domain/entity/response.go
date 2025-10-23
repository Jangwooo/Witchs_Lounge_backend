package entity

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Details interface{} `json:"details,omitempty"`
}
