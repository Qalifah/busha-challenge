package handler

type Response struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(message string, data interface{}) *Response {
	return &Response{
		Status:  StatusSuccess,
		Message: message,
		Data:    data,
	}
}

type ErrorResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Status: StatusError, Message: err.Error()}
}

const (
	StatusError   = "error"
	StatusSuccess = "success"
)
