package places

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"sk8.town/backside/errs"
	"sk8.town/backside/places/dto"

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

	dummyPlaces := []*dto.AutocompleteEntry{
		{
			Coordinates: &dto.Coordinates{
				Lat:  5,
				Long: 6,
			},
			Name:    "name1",
			Address: "address1",
		},
		{
			Coordinates: &dto.Coordinates{
				Lat:  10,
				Long: 20,
			},
			Name:    "name2",
			Address: "address2",
		},
	}
	serviceMock.EXPECT().GetPlaces(&dummyPlacesAutocompleteRequest).Return(dummyPlaces, nil)
	router.HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete?search=xxx&language=pl", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse := `{"status":"success","data":[{"coordinates":{"lat":5,"long":6},"name":"name1","address":"address1"},{"coordinates":{"lat":10,"long":20},"name":"name2","address":"address2"}]}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_given_bad_request_should_return_400(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete?random_param=xxx", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	expectedResponse := `{"status":"error","message":"Bad request"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_given_service_error_should_return_service_error(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	dummyPlacesAutocompleteRequest := dto.AutocompleteRequest{
		Search:   "xxx",
		Language: "pl",
	}
	serviceMock.EXPECT().GetPlaces(&dummyPlacesAutocompleteRequest).Return(nil, errs.NewNotFoundError(""))
	router.HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete?search=xxx&language=pl", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expectedResponse := `{"status":"error","message":"Not found"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}
