package error

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewInternalServerError(m string) *RestError {
	return &RestError{
		Message: m,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewBadRequestError(m string) *RestError {
	return &RestError{
		Message: m,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(m string) *RestError {
	return &RestError{
		Message: m,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
