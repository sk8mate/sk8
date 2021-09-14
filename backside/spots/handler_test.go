package spots

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"sk8.town/backside/errs"
	"sk8.town/backside/spots/dto"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"sk8.town/backside/spots/mocks"
)

var router *mux.Router
var handler Handler
var serviceMock *mocks.MockSpotService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	serviceMock = mocks.NewMockSpotService(ctrl)
	handler = Handler{serviceMock}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_given_correct_request_should_return_spot_id_with_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	addSpotRequest := dto.SpotRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.SpotCoordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	addedSpotResponse := dto.SpotResponse{
		Id: "random id",
	}
	serviceMock.EXPECT().Add(&addSpotRequest).Return(&addedSpotResponse, nil)
	router.HandleFunc("/spots", handler.AddSpot)
	var jsonStr = []byte(`{"name":"Dworzec Glowny Krakow","address":"Pawia 5","coordinates":{"lat":40,"long":60},"lighting":true,"friendly":true,"verified":true}`)
	request, _ := http.NewRequest(http.MethodPost, "/spots", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse := `{"id":"random id"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_given_bad_request_should_return_400(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/spots", handler.AddSpot)
	var jsonStr = []byte(``)
	request, _ := http.NewRequest(http.MethodPost, "/spots", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	expectedResponse := `{"message":"Bad request"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_given_service_error_should_return_service_error(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	addSpotRequest := dto.SpotRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.SpotCoordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	serviceMock.EXPECT().Add(&addSpotRequest).Return(nil, errs.NewNotFoundError(""))
	router.HandleFunc("/spots", handler.AddSpot)
	var jsonStr = []byte(`{"name":"Dworzec Glowny Krakow","address":"Pawia 5","coordinates":{"lat":40,"long":60},"lighting":true,"friendly":true,"verified":true}`)
	request, _ := http.NewRequest(http.MethodPost, "/spots", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expectedResponse := `{"message":"Not found"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}
