package utils

import (
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
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

func WriteProtoMessage(writer http.ResponseWriter, code int, data proto.Message) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(code)
	m := jsonpb.Marshaler{EmitDefaults: true}
	err := m.Marshal(writer, data)
	if err != nil {
		panic(err)
	}
}

func WriteStatus(writer http.ResponseWriter, code int) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
}

