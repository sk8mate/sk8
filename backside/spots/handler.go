package spots

import (
	"encoding/json"
	"net/http"
	"sk8.town/backside/errs"

	"sk8.town/backside/spots/dto"
)

type Handler struct {
	service SpotService
}

func (handler Handler) AddSpot(writer http.ResponseWriter, request *http.Request) {
	var spotRequest dto.SpotRequest
	if err := json.NewDecoder(request.Body).Decode(&spotRequest); err != nil {
		appError := errs.NewBadRequestError("")
		writeResponse(writer, appError.Code, appError.AsMessage())
		return
	}

	response, appError := handler.service.Add(&spotRequest)
	if appError != nil {
		writeResponse(writer, appError.Code, appError.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, response)
	}
}

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		panic(err)
	}
}
