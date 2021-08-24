package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

func NewNotFoundError(message string) *AppError {
	if message == "" {
		message = "Not found"
	}

	return &AppError{Message: message, Code: http.StatusNotFound}
}

func NewUnexpectedError(message string) *AppError {
	if message == "" {
		message = "Internal server error"
	}

	return &AppError{Message: message, Code: http.StatusInternalServerError}
}

func NewValidationError(message string) *AppError {
	if message == "" {
		message = "Validation failed"
	}

	return &AppError{Message: message, Code: http.StatusUnprocessableEntity}
}
