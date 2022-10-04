package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	restErr := RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}

	return &restErr
}

func NewNotFoundError(message string) *RestErr {
	restErr := RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}

	return &restErr
}
