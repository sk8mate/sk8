package places

import (
	"regexp"
	"testing"

	"sk8.town/backside/places/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"sk8.town/backside/errs"
	"sk8.town/backside/places/domain"
	"sk8.town/backside/places/dto"
)

func Test_should_return_an_error_when_request_is_not_valid(t *testing.T) {
	request := dto.AutocompleteRequest{
		Search:   "",
		Language: "",
	}
	service := NewService(nil)
	_, appError := service.GetPlaces(&request)

	assert.Regexp(t, regexp.MustCompile("\"search\" is required"), appError.Message)
}

func Test_should_propagate_an_error_from_http_service(t *testing.T) {
	request := dto.AutocompleteRequest{
		Search:   "warsz",
		Language: "pl",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLocationService := mocks.NewMockLocationService(ctrl)
	service := NewService(mockLocationService)
	expectedError := errs.NewNotFoundError("not found error")
	mockLocationService.EXPECT().Search("warsz", "pl-PL").Return(nil, expectedError)

	_, appError := service.GetPlaces(&request)

	assert.Equal(t, expectedError, appError)
}

func Test_should_return_places_response_when_places_retrieved_successfully(t *testing.T) {
	request := dto.AutocompleteRequest{
		Search:   "search",
		Language: "pl",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLocationService := mocks.NewMockLocationService(ctrl)
	service := NewService(mockLocationService)
	places := domain.GetPlacesResponse{
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
		},
	}
	mockLocationService.EXPECT().Search("search", "pl-PL").Return(&places, nil)

	actualPlaces, appError := service.GetPlaces(&request)

	var expectedPlace = &dto.AutocompleteEntry{
		Coordinates: &dto.Coordinates{
			Lat:  1,
			Long: 2,
		},
		Name:    "Free form address",
		Address: "",
	}
	assert.Equal(t, 1, len(actualPlaces))
	assert.Equal(t, expectedPlace, actualPlaces[0])
	assert.Nil(t, appError)
}
