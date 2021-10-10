package utils

import (
	"encoding/json"
	"net/http"

	"sk8.town/backside/errs"
)

func WriteError(writer http.ResponseWriter, appError *errs.AppError) {
	type ErrorResponse struct {
		Status  string      `json:"status"`
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message,omitempty"`
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(appError.Code)

	err := json.NewEncoder(writer).Encode(ErrorResponse{
		Status:  "error",
		Data:    appError.Data,
		Message: appError.Message,
	})

	if err != nil {
		panic(err)
	}
}

func WriteJSON(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		panic(err)
	}
}
