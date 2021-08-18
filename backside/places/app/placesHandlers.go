package app

import (
	"encoding/json"
	"net/http"

	"sk8.town/backside/places/dto"
	"sk8.town/backside/places/service"
)

type PlacesHandlers struct {
	service service.PlacesService
}

func (placesHandlers *PlacesHandlers) getPlacesByAutocomplete(writer http.ResponseWriter, request *http.Request) {
	var placesAutocompleteRequest dto.PlacesAutocompleteRequest
	err := json.NewDecoder(request.Body).Decode(&placesAutocompleteRequest)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		response, appError := placesHandlers.service.GetPlaces(placesAutocompleteRequest)
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
