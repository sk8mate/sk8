package spots

import (
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

var filterRouter *mux.Router
var filtersHandler FiltersHandler
var filterServiceMock *mocks.MockFilterService

func initialize(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	filterServiceMock = mocks.NewMockFilterService(ctrl)
	filtersHandler = FiltersHandler{filterServiceMock}
	filterRouter = mux.NewRouter()
	return func() {
		filterRouter = nil
		defer ctrl.Finish()
	}
}

func Test_get_filters_retrieved_successfully_from_service_should_return_spots_with_status_200(t *testing.T) {
	teardown := initialize(t)
	defer teardown()
	filtersGetAllData := []*dto.FilterData{
		{
			Id:   "1",
			Name: "filter1",
			Type: "type1",
			Values: []*dto.FilterValueData{
				{
					Id:   "1",
					Name: "value1",
				},
				{
					Id:   "2",
					Name: "value2",
				},
			},
		},
		{
			Id:   "2",
			Name: "filter2",
			Type: "type2",
			Values: []*dto.FilterValueData{
				{
					Id:   "3",
					Name: "value3",
				},
			},
		},
	}
	filterServiceMock.EXPECT().GetAll().Return(filtersGetAllData, nil)
	filterRouter.HandleFunc("/spots/filters", filtersHandler.GetFilters)
	request, _ := http.NewRequest(http.MethodGet, "/spots/filters", nil)
	recorder := httptest.NewRecorder()

	filterRouter.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse := `{"status":"success","data":[{"id":"1","name":"filter1","type":"type1","values":[{"id":"1","name":"value1"},{"id":"2","name":"value2"}]},{"id":"2","name":"filter2","type":"type2","values":[{"id":"3","name":"value3"}]}]}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}

func Test_get_filters_failed_from_service_should_return_error(t *testing.T) {
	teardown := initialize(t)
	defer teardown()
	filterServiceMock.EXPECT().GetAll().Return(nil, errs.NewNotFoundError(""))
	filterRouter.HandleFunc("/spots/filters", filtersHandler.GetFilters)
	request, _ := http.NewRequest(http.MethodGet, "/spots/filters", nil)
	recorder := httptest.NewRecorder()

	filterRouter.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expectedResponse := `{"status":"error","message":"Not found"}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(recorder.Body.String()))
}
