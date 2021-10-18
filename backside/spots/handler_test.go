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

func Test_add_spot_given_correct_request_should_return_spot_id_with_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	addSpotRequest := dto.SpotsAddRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.SpotsAddRequest_Coordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	addedSpotsAddData := dto.SpotsAddData{
		Id: "random id",
	}
	serviceMock.EXPECT().Add(&addSpotRequest).Return(&addedSpotsAddData, nil)
	router.HandleFunc("/spots", handler.AddSpot)
	var jsonStr = []byte(`{"name":"Dworzec Glowny Krakow","address":"Pawia 5","coordinates":{"lat":40,"long":60},"lighting":true,"friendly":true,"verified":true}`)
	request, _ := http.NewRequest(http.MethodPost, "/spots", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse := `{"status":"success","data":{"id":"random id"}}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_add_spot_given_bad_request_should_return_400(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/spots", handler.AddSpot)
	var jsonStr = []byte(``)
	request, _ := http.NewRequest(http.MethodPost, "/spots", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	expectedResponse := `{"status":"error","message":"Bad request"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_add_spot_given_service_error_should_return_service_error(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	addSpotRequest := dto.SpotsAddRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.SpotsAddRequest_Coordinates{
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
	expectedResponse := `{"status":"error","message":"Not found"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_get_spot_given_correct_request_should_return_spot_with_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	spotGetData := dto.SpotsGetData{
		Id:      "5",
		Name:    "Dom",
		Address: "Grzegorzecka 79f Krakow",
		Coordinates: &dto.ResponseCoordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	id := 5
	serviceMock.EXPECT().Get(id).Return(&spotGetData, nil)
	router.HandleFunc("/spots/{id:[0-9]+}", handler.GetSpot)
	request, _ := http.NewRequest(http.MethodGet, "/spots/5", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse := `{"status":"success","data":{"id":"5","name":"Dom","address":"Grzegorzecka 79f Krakow","coordinates":{"lat":40,"long":60},"lighting":true,"friendly":true,"verified":true}}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_get_spot_given_service_error_should_return_service_error(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	id := 5
	serviceMock.EXPECT().Get(id).Return(nil, errs.NewNotFoundError(""))
	router.HandleFunc("/spots/{id:[0-9]+}", handler.GetSpot)
	request, _ := http.NewRequest(http.MethodPost, "/spots/5", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expectedResponse := `{"status":"error","message":"Not found"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_get_spots_given_correct_request_should_return_spots_with_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	spotsGetAllData := []*dto.SpotsGetData{
		{Id: "5",
			Name:    "Dom",
			Address: "Grzegorzecka 79f Krakow",
			Coordinates: &dto.ResponseCoordinates{
				Lat:  40,
				Long: 60,
			},
			Lighting: true,
			Friendly: true,
			Verified: true,
		},
		{Id: "6",
			Name:    "Sasiad",
			Address: "Grzegorzecka 79e Krakow",
			Coordinates: &dto.ResponseCoordinates{
				Lat:  40,
				Long: 60,
			},
			Lighting: false,
			Friendly: false,
			Verified: false,
		},
	}
	serviceMock.EXPECT().GetAll().Return(spotsGetAllData, nil)
	router.HandleFunc("/spots", handler.GetSpots)
	request, _ := http.NewRequest(http.MethodGet, "/spots", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse := `{"status":"success","data":[{"id":"5","name":"Dom","address":"Grzegorzecka 79f Krakow","coordinates":{"lat":40,"long":60},"lighting":true,"friendly":true,"verified":true},{"id":"6","name":"Sasiad","address":"Grzegorzecka 79e Krakow","coordinates":{"lat":40,"long":60},"lighting":false,"friendly":false,"verified":false}]}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_get_spots_given_service_error_should_return_service_error(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	serviceMock.EXPECT().GetAll().Return(nil, errs.NewNotFoundError(""))
	router.HandleFunc("/spots", handler.GetSpots)
	request, _ := http.NewRequest(http.MethodPost, "/spots", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expectedResponse := `{"status":"error","message":"Not found"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_update_spot_given_correct_request_should_return_updated_spot_with_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	updateSpotRequest := dto.SpotsUpdateRequest{
		Name: "Pizza",
	}
	updatedSpotData := dto.SpotsUpdateData{
		Id:      "5",
		Name:    "Pizza",
		Address: "Pawia 5",
		Coordinates: &dto.ResponseCoordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: false,
		Friendly: false,
		Verified: false,
	}
	id := 5
	serviceMock.EXPECT().Update(id, &updateSpotRequest).Return(&updatedSpotData, nil)
	router.HandleFunc("/spots/{id:[0-9]+}", handler.UpdateSpot)
	var jsonStr = []byte(`{"name":"Pizza"}`)
	request, _ := http.NewRequest(http.MethodPut, "/spots/5", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse := `{"status":"success","data":{"id":"5","name":"Pizza","address":"Pawia 5","coordinates":{"lat":40,"long":60},"lighting":false,"friendly":false,"verified":false}}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_update_spot_given_bad_request_should_return_400(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/spots/{id:[0-9]+}", handler.UpdateSpot)
	var jsonStr = []byte(``)
	request, _ := http.NewRequest(http.MethodPut, "/spots/5", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	expectedResponse := `{"status":"error","message":"Bad request"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_update_spot_given_service_error_should_return_service_error(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	updateSpotRequest := dto.SpotsUpdateRequest{
		Name: "Pizza",
	}
	id := 5
	serviceMock.EXPECT().Update(id, &updateSpotRequest).Return(nil, errs.NewNotFoundError(""))
	router.HandleFunc("/spots/{id:[0-9]+}", handler.UpdateSpot)
	var jsonStr = []byte(`{"name":"Pizza"}`)
	request, _ := http.NewRequest(http.MethodPut, "/spots/5", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expectedResponse := `{"status":"error","message":"Not found"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_delete_spot_given_correct_request_should_return_status_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	id := 5
	serviceMock.EXPECT().Delete(id).Return(nil)
	router.HandleFunc("/spots/{id:[0-9]+}", handler.DeleteSpot)
	request, _ := http.NewRequest(http.MethodDelete, "/spots/5", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_delete_spot_given_service_error_should_return_service_error(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	id := 5
	serviceMock.EXPECT().Delete(id).Return(errs.NewNotFoundError(""))
	router.HandleFunc("/spots/{id:[0-9]+}", handler.DeleteSpot)
	request, _ := http.NewRequest(http.MethodDelete, "/spots/5", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expectedResponse := `{"status":"error","message":"Not found"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}
