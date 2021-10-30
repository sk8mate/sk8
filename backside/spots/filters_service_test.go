package spots

import (
	"testing"

	"sk8.town/backside/spots/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"sk8.town/backside/errs"
	"sk8.town/backside/spots/domain"
	"sk8.town/backside/spots/dto"
)

func Test_get_all_filters_should_propagate_an_error_from_db(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockFiltersRepository(ctrl)
	service := NewFilterService(mockDb)
	expectedError := errs.NewNotFoundError("not found error")
	mockDb.EXPECT().GetAll().Return(nil, expectedError)

	_, appError := service.GetAll()

	assert.Equal(t, expectedError, appError)
}

func Test_get_all_filters_should_return_filters_response_when_filters_retrieved_successfully(t *testing.T) {
	filterValues := []*domain.Filter{
		{
			ID:   1,
			Name: "filter1",
			Type: "type1",
		},
		{
			ID:   2,
			Name: "filter2",
			Type: "type2",
			FilterValues: []domain.FilterValue{
				{
					ID:    1,
					Value: "value1",
				},
				{
					ID:    2,
					Value: "value2",
				},
			},
		},
	}
	expectedFilterDtoData := []*dto.FilterData{
		{
			Id:   "1",
			Name: "filter1",
			Type: "type1",
		},
		{
			Id:   "2",
			Name: "filter2",
			Type: "type2",
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
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockFiltersRepository(ctrl)
	service := NewFilterService(mockDb)
	mockDb.EXPECT().GetAll().Return(filterValues, nil)

	filtersGetAllData, appError := service.GetAll()

	assert.Nil(t, appError)
	assert.Equal(t, expectedFilterDtoData, filtersGetAllData)
}
