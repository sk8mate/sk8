package places

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"sk8.town/backside/places/domain"
	"sk8.town/backside/places/mocks"
)

var router *mux.Router
var handler Handler
var service DefaultService
var repositoryMock *mocks.MockPlacesRepository

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	repositoryMock = mocks.NewMockPlacesRepository(ctrl)
	service = NewService(repositoryMock)
	handler = Handler{service}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_given_correct_request_should_return_places_with_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockedResponse := domain.GetPlacesResponse{
		Results: []domain.Result{
			{
				Type: "Geography",
				Position: domain.Position{
					Lat: 1,
					Lon: 2,
				},
				Address: domain.Address{
					FreeformAddress: "Free form address",
				},
			},
			{
				Type: "POI",
				Position: domain.Position{
					Lat: 3,
					Lon: 4,
				},
				Address: domain.Address{
					FreeformAddress: "Free form address",
				},
				Poi: domain.Poi{
					Name: "Poi",
				},
			},
		},
	}

	repositoryMock.EXPECT().GetPlaces("xxx", "pl").Return(&mockedResponse, nil)
	router.HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete)
	var jsonStr = []byte(`{"search":"xxx", "language":"pl"}`)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete", bytes.NewBuffer(jsonStr))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Code, http.StatusOK)

	expectedResponse := `[{"coordinates":{"lat":1,"long":2},"name":"Free form address","address":""},{"coordinates":{"lat":3,"long":4},"name":"Poi","address":"Free form address"}]`
	assert.Equal(t, strings.TrimSpace(recorder.Body.String()), expectedResponse)
}

func Test_given_incorrect_request_should_return_400(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete)
	var jsonStr = []byte(``)
	request, _ := http.NewRequest(http.MethodGet, "/places/autocomplete", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Code, http.StatusBadRequest)
}
