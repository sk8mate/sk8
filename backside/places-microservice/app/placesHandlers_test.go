package app

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sk8mate/sk8/backside/places-microservice/dto"
	"github.com/sk8mate/sk8/backside/places-microservice/errs"
	"github.com/sk8mate/sk8/backside/places-microservice/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var placesHandlers PlacesHandlers
var mockService *service.MockPlacesService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockPlacesService(ctrl)
	placesHandlers = PlacesHandlers{mockService}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_given_correct_request_and_correct_response_from_service_should_return_places_with_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	dummyPlacesAutocompleteRequest := dto.PlacesAutocompleteRequest{
		Search:   "xxx",
		Language: "pl",
	}
	dummyPlaces := []dto.PlacesAutocompleteResponseEntry{
		{
			Coordinates: struct {
				Lat  float64 `json:"lat"`
				Long float64 `json:"long"`
			}{5, 6},
			Name:    "name1",
			Address: "address1",
		},
		{
			Coordinates: struct {
				Lat  float64 `json:"lat"`
				Long float64 `json:"long"`
			}{10, 20},
			Name:    "name2",
			Address: "address2",
		},
	}
	mockService.EXPECT().GetPlaces(dummyPlacesAutocompleteRequest).Return(dummyPlaces, nil)
	router.HandleFunc("/places/autocomplete", placesHandlers.getPlacesByAutocomplete)
	var jsonStr = []byte(`{"search":"xxx", "language":"pl"}`)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete", bytes.NewBuffer(jsonStr))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_given_correct_request_and_error_from_service_should_propagate_error_from_service(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	dummyPlacesAutocompleteRequest := dto.PlacesAutocompleteRequest{
		Search:   "xxx",
		Language: "pl",
	}
	mockService.EXPECT().GetPlaces(dummyPlacesAutocompleteRequest).Return(nil, errs.NewNotFoundError("not found"))
	router.HandleFunc("/places/autocomplete", placesHandlers.getPlacesByAutocomplete)
	var jsonStr = []byte(`{"search":"xxx", "language":"pl"}`)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete", bytes.NewBuffer(jsonStr))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_given_incorrect_request_should_return_400(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/places/autocomplete", placesHandlers.getPlacesByAutocomplete)
	var jsonStr = []byte(``)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete", bytes.NewBuffer(jsonStr))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Error("Failed while testing the status code")
	}
}
