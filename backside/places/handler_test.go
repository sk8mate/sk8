package places

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sk8.town/backside/places/dto"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"sk8.town/backside/places/mocks"
)

var router *mux.Router
var handler Handler
var serviceMock *mocks.MockService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	serviceMock = mocks.NewMockService(ctrl)
	handler = Handler{serviceMock}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_given_correct_request_should_return_places_with_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	dummyPlacesAutocompleteRequest := dto.AutocompleteRequest{
		Search:   "xxx",
		Language: "pl",
	}
	dummyPlaces := []dto.AutocompleteEntryResponse{
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
	serviceMock.EXPECT().GetPlaces(dummyPlacesAutocompleteRequest).Return(dummyPlaces, nil)
	router.HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete)
	var jsonStr = []byte(`{"search":"xxx", "language":"pl"}`)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete", bytes.NewBuffer(jsonStr))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Code, http.StatusOK)

	expectedResponse := `[{"coordinates":{"lat":5,"long":6},"name":"name1","address":"address1"},{"coordinates":{"lat":10,"long":20},"name":"name2","address":"address2"}]`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_given_incorrect_request_should_return_400(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete)
	var jsonStr = []byte(``)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
