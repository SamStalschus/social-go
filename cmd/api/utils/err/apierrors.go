package apierrors

import (
	"fmt"
	"net/http"
)

type ApiError interface {
	Message() string
	Code() string
	Status() uint
	Error() string
}

type apiErr struct {
	ErrorMessage string `json:"message"`
	ErrorCode    string `json:"error"`
	ErrorStatus  uint   `json:"status"`
}

const INTERNAL_SERVER_ERROR = "Internal server error"

func (err apiErr) Code() string {
	return err.ErrorCode
}

func (err apiErr) Error() string {
	return fmt.Sprintf("Message: %s; Error Code: %s; Status: %d", err.ErrorMessage, err.ErrorCode, err.ErrorStatus)
}

func (err apiErr) Status() uint {
	return err.ErrorStatus
}

func (err apiErr) Message() string {
	return err.Message()
}

func NewApiError(message string, error string, status uint) ApiError {
	return apiErr{message, error, status}
}

func NewNotFoundApiError(message string) ApiError {
	return apiErr{message, "not_found", http.StatusNotFound}
}

func NewBadRequestApiError(message string) ApiError {
	return apiErr{message, "bad_request", http.StatusBadRequest}
}

func NewInternalServerError() ApiError {
	return apiErr{INTERNAL_SERVER_ERROR, "internal_error", http.StatusInternalServerError}
}
