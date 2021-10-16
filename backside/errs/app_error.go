package errs

import (
	"net/http"

	"sk8.town/backside/config"
)

type AppError struct {
	Code    int
	Message string
	Data    interface{}
}

func NewNotFoundError(message string) *AppError {
	if message == "" {
		message = "Not found"
	}

	return &AppError{Message: message, Code: http.StatusNotFound}
}

func NewUnexpectedError(message string) *AppError {
	isProduction := config.GetEnv() == config.Env.Production

	if message == "" || isProduction {
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

func NewBadRequestError(message string) *AppError {
	if message == "" {
		message = "Bad request"
	}

	return &AppError{Message: message, Code: http.StatusBadRequest}
}
