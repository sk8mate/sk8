package places

import (
	"encoding/json"
	"net/http"

	"sk8.town/backside/places/dto"
)

type Handler struct {
	Service Service
}

func (handler Handler) GetPlacesAutocomplete(writer http.ResponseWriter, request *http.Request) {
	var placesRequest dto.AutocompleteRequest
	err := json.NewDecoder(request.Body).Decode(&placesRequest)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		response, appError := handler.Service.GetPlaces(placesRequest)
		if appError != nil {
			writeResponse(writer, appError.Code, appError.Message)
		} else {
			writeResponse(writer, http.StatusOK, response)
		}
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
