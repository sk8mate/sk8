package utils

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"

	"sk8.town/backside/errs"
)

type ErrorResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func WriteError(writer http.ResponseWriter, appError *errs.AppError) {
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

func WriteProtoJSON(writer http.ResponseWriter, code int, data proto.Message) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(code)
	marshalOptions := protojson.MarshalOptions{EmitUnpopulated: true}
	jsonData, err := marshalOptions.Marshal(data)
	if err != nil {
		panic(err)
	}
	_, err = writer.Write(jsonData)
	if err != nil {
		panic(err)
	}
}

func WriteStatus(writer http.ResponseWriter, code int) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
}
