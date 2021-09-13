package places

import (
	"encoding/json"
	"net/http"
	"sk8.town/backside/errs"

	"sk8.town/backside/places/dto"
)

type Handler struct {
	Service Service
}

func (handler Handler) GetPlacesAutocomplete(writer http.ResponseWriter, request *http.Request) {
	var placesRequest dto.AutocompleteRequest
	if err := json.NewDecoder(request.Body).Decode(&placesRequest); err != nil {
		appError := errs.NewBadRequestError("")
		writeResponse(writer, appError.Code, appError.AsMessage())
		return
	}

	response, appError := handler.Service.GetPlaces(&placesRequest)
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