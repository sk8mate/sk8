package places

import (
	"regexp"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"sk8.town/backside/errors"
	"sk8.town/backside/places/domain"
	"sk8.town/backside/places/dto"
	"sk8.town/backside/places/mocks"
)

func Test_should_return_an_error_when_request_is_not_valid(t *testing.T) {
	request := dto.AutocompleteRequest{
		Search:   "",
		Language: "",
	}
	service := NewService(nil)
	_, appError := service.GetPlaces(request)

	assert.Regexp(t, regexp.MustCompile("\"search\" is required"), appError.Message)

}

func Test_should_propagate_an_error_from_places_repository(t *testing.T) {
	request := dto.AutocompleteRequest{
		Search:   "warsz",
		Language: "pl",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockPlacesRepository(ctrl)
	service := NewService(mockRepo)
	expectedError := errors.NewNotFoundError("not found error")
	mockRepo.EXPECT().GetPlaces(request.Search, request.Language).Return(nil, expectedError)

	_, appError := service.GetPlaces(request)

	assert.Equal(t, appError, expectedError)
}

func Test_should_return_places_response_when_places_retrieved_successfully(t *testing.T) {
	request := dto.AutocompleteRequest{
		Search:   "search",
		Language: "pl",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockPlacesRepository(ctrl)
	service := NewService(mockRepo)
	places := domain.GetPlacesResponse{}
	mockRepo.EXPECT().GetPlaces(request.Search, request.Language).Return(&places, nil)

	actualPlaces, appError := service.GetPlaces(request)

	assert.NotNil(t, actualPlaces)
	assert.Nil(t, appError)
}

// //TODO: add test with filled places (not able to initialize that complex struct for now)
