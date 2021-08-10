package app

import (
	"encoding/json"
	"github.com/sk8mate/sk8/backside/dto"
	"github.com/sk8mate/sk8/backside/service"
	"net/http"
)

type PlacesHandlers struct {
	service service.PlacesAutocompleteService
}

func (placesHandlers *PlacesHandlers) getPlacesByAutocomplete(writer http.ResponseWriter, request *http.Request) {
	var placesAutocompleteRequest dto.PlacesAutocompleteRequest
	err := json.NewDecoder(request.Body).Decode(&placesAutocompleteRequest)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		response, appError := placesHandlers.service.GetPlaces(placesAutocompleteRequest.Search, placesAutocompleteRequest.Language)
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
