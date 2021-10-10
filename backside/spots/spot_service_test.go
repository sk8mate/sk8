package spots

import (
	"strconv"
	"testing"

	"sk8.town/backside/spots/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"sk8.town/backside/errs"
	"sk8.town/backside/spots/domain"
	"sk8.town/backside/spots/dto"
)

func Test_add_request_should_propagate_an_error_from_db(t *testing.T) {
	request := dto.SpotsAddRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.RequestCoordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	spotToAdd := domain.Spot{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: domain.Coordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotService(mockDb)
	expectedError := errs.NewNotFoundError("not found error")
	mockDb.EXPECT().Add(&spotToAdd).Return(nil, expectedError)

	_, appError := service.Add(&request)

	assert.Equal(t, expectedError, appError)
}

func Test_add_request_should_return_spots_response_when_spot_added_successfully(t *testing.T) {
	request := dto.SpotsAddRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.RequestCoordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	spotToAdd := domain.Spot{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: domain.Coordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	createdSpot := domain.Spot{
		Id:      5,
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: domain.Coordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotService(mockDb)
	mockDb.EXPECT().Add(&spotToAdd).Return(&createdSpot, nil)

	spotsAddData, appError := service.Add(&request)

	assert.Nil(t, appError)
	assert.Equal(t, strconv.Itoa(createdSpot.Id), spotsAddData.Id)
}

func Test_get_request_should_propagate_an_error_from_db(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotService(mockDb)
	expectedError := errs.NewNotFoundError("not found error")
	id := 4
	mockDb.EXPECT().Get(id).Return(nil, expectedError)

	_, appError := service.Get(id)

	assert.Equal(t, expectedError, appError)
}

func Test_get_request_should_return_spot_response_when_spot_retrieved_successfully(t *testing.T) {
	spot := domain.Spot{
		Id:      4,
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: domain.Coordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	expectedSpotDtoData := dto.SpotsGetData{
		Id:      "4",
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.ResponseCoordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotService(mockDb)
	id := 4
	mockDb.EXPECT().Get(id).Return(&spot, nil)

	spotsGetData, appError := service.Get(id)

	assert.Nil(t, appError)
	assert.Equal(t, expectedSpotDtoData, *spotsGetData)
}
