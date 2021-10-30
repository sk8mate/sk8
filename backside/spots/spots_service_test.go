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

func Test_given_invalid_add_request_should_return_unprocessable_entity(t *testing.T) {
	request := dto.SpotsAddRequest{
		Address: "Pawia 5",
		Coordinates: &dto.SpotsAddRequest_Coordinates{
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
	service := NewSpotsService(mockDb)
	expectedError := errs.NewValidationError("invalid SpotsAddRequest.Name: value length must be at least 1 runes")

	_, appError := service.Add(&request)

	assert.Equal(t, expectedError, appError)
}

func Test_add_request_should_propagate_an_error_from_db(t *testing.T) {
	request := dto.SpotsAddRequest{
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
	service := NewSpotsService(mockDb)
	expectedError := errs.NewNotFoundError("not found error")
	mockDb.EXPECT().Add(&spotToAdd).Return(nil, expectedError)

	_, appError := service.Add(&request)

	assert.Equal(t, expectedError, appError)
}

func Test_add_request_should_return_spots_response_when_spot_added_successfully(t *testing.T) {
	request := dto.SpotsAddRequest{
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
	service := NewSpotsService(mockDb)
	mockDb.EXPECT().Add(&spotToAdd).Return(&createdSpot, nil)

	spotsAddData, appError := service.Add(&request)

	assert.Nil(t, appError)
	assert.Equal(t, strconv.Itoa(createdSpot.Id), spotsAddData.Id)
}

func Test_get_request_should_propagate_an_error_from_db(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotsService(mockDb)
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
	expectedSpotDtoData := &dto.SpotsGetData{
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
	service := NewSpotsService(mockDb)
	id := 4
	mockDb.EXPECT().Get(id).Return(&spot, nil)

	spotsGetData, appError := service.Get(id)

	assert.Nil(t, appError)
	assert.Equal(t, expectedSpotDtoData, spotsGetData)
}

func Test_get_all_request_should_propagate_an_error_from_db(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotsService(mockDb)
	expectedError := errs.NewNotFoundError("not found error")
	mockDb.EXPECT().GetAll().Return(nil, expectedError)

	_, appError := service.GetAll()

	assert.Equal(t, expectedError, appError)
}

func Test_get_all_request_should_return_spots_response_when_spots_retrieved_successfully(t *testing.T) {
	spots := []*domain.Spot{
		{
			Id:      1,
			Name:    "Dworzec Glowny Krakow",
			Address: "Pawia 5",
			Coordinates: domain.Coordinates{
				Lat:  40,
				Long: 60,
			},
			Lighting: true,
			Friendly: true,
			Verified: true,
		},
		{
			Id:      2,
			Name:    "Brama Zuraw",
			Address: "Szeroka 67/68, 22-100 Gdansk",
			Coordinates: domain.Coordinates{
				Lat:  54.35,
				Long: 18.66,
			},
			Lighting: true,
			Friendly: false,
			Verified: true,
		},
	}
	expectedSpotsDtoData := []*dto.SpotsGetData{
		{
			Id:      "1",
			Name:    "Dworzec Glowny Krakow",
			Address: "Pawia 5",
			Coordinates: &dto.ResponseCoordinates{
				Lat:  40,
				Long: 60,
			},
			Lighting: true,
			Friendly: true,
			Verified: true,
		},
		{
			Id:      "2",
			Name:    "Brama Zuraw",
			Address: "Szeroka 67/68, 22-100 Gdansk",
			Coordinates: &dto.ResponseCoordinates{
				Lat:  54.35,
				Long: 18.66,
			},
			Lighting: true,
			Friendly: false,
			Verified: true,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotsService(mockDb)
	mockDb.EXPECT().GetAll().Return(spots, nil)

	spotsGetAllData, appError := service.GetAll()

	assert.Nil(t, appError)
	assert.Equal(t, expectedSpotsDtoData, spotsGetAllData)
}

func Test_given_invalid_update_request_should_return_unprocessable_entity(t *testing.T) {
	request := dto.SpotsUpdateRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.SpotsUpdateRequest_Coordinates{
			Lat:  400,
			Long: 6000,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotsService(mockDb)
	expectedError := errs.NewValidationError("invalid SpotsUpdateRequest.Coordinates: embedded message failed validation | caused by: invalid SpotsUpdateRequest_Coordinates.Lat: value must be inside range [-90, 90]")
	id := 4

	_, appError := service.Update(id, &request)

	assert.Equal(t, expectedError, appError)
}

func Test_update_request_should_propagate_an_error_from_db(t *testing.T) {
	spot := domain.Spot{
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
	request := dto.SpotsUpdateRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &dto.SpotsUpdateRequest_Coordinates{
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
	service := NewSpotsService(mockDb)
	expectedError := errs.NewNotFoundError("not found error")
	id := 4
	mockDb.EXPECT().Update(id, &spot).Return(nil, expectedError)

	_, appError := service.Update(id, &request)

	assert.Equal(t, expectedError, appError)
}

func Test_update_request_should_return_updated_spot_response_when_spot_updated_successfully(t *testing.T) {
	spot := domain.Spot{
		Name: "Rynek Krakow",
	}
	updatedSpot := domain.Spot{
		Id:      10,
		Name:    "Rynek Krakow",
		Address: "Pawia 5",
		Coordinates: domain.Coordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}
	request := dto.SpotsUpdateRequest{
		Name: "Rynek Krakow",
	}
	expectedSpotsDtoData := &dto.SpotsUpdateData{
		Id:      "10",
		Name:    "Rynek Krakow",
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
	service := NewSpotsService(mockDb)
	id := 4
	mockDb.EXPECT().Update(id, &spot).Return(&updatedSpot, nil)

	spotUpdatedData, appError := service.Update(id, &request)

	assert.Nil(t, appError)
	assert.Equal(t, expectedSpotsDtoData, spotUpdatedData)
}

func Test_delete_request_should_propagate_an_error_from_db(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotsService(mockDb)
	expectedError := errs.NewNotFoundError("not found error")
	id := 4
	mockDb.EXPECT().Delete(id).Return(expectedError)

	appError := service.Delete(id)

	assert.Equal(t, expectedError, appError)
}

func Test_delete_request_should_return_no_error_when_spot_deleted_successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockSpotRepository(ctrl)
	service := NewSpotsService(mockDb)
	id := 4
	mockDb.EXPECT().Delete(id).Return(nil)

	appError := service.Delete(id)

	assert.Nil(t, appError)
}
