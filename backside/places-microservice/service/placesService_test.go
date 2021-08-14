package service

import (
	"github.com/golang/mock/gomock"
	realdomain "github.com/sk8mate/sk8/backside/places-microservice/domain"
	"github.com/sk8mate/sk8/backside/places-microservice/dto"
	"github.com/sk8mate/sk8/backside/places-microservice/errs"
	"github.com/sk8mate/sk8/backside/places-microservice/mocks/domain"
	"testing"
)

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T){
	request := dto.PlacesAutocompleteRequest{
		Search:   "",
		Language: "",
	}
	service := NewPlacesService(nil)

	_, appError := service.GetPlaces(request)

	if appError == nil {
		t.Error("failed while testing places autocomplete request validation")
	}
}

func Test_should_return_an_error_from_the_server_side_if_can_not_get_places(t *testing.T){
	request := dto.PlacesAutocompleteRequest{
		Search:   "search",
		Language: "language",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := domain.NewMockPlacesRepository(ctrl)
	service := NewPlacesService(mockRepo)
	mockRepo.EXPECT().GetPlaces("search", "language").Return(nil, errs.NewNotFoundError("not found error"))

	_, appError := service.GetPlaces(request)

	if appError == nil {
		t.Error("failed while validating error for places autocomplete")
	}
}

func Test_should_return_places_response_when_places_retrieved_successfully(t *testing.T){
	request := dto.PlacesAutocompleteRequest{
		Search:   "search",
		Language: "language",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := domain.NewMockPlacesRepository(ctrl)
	service := NewPlacesService(mockRepo)
	places := realdomain.Places{
		Summary: struct {
			Query        string `json:"query"`
			QueryType    string `json:"queryType"`
			QueryTime    int    `json:"queryTime"`
			NumResults   int    `json:"numResults"`
			Offset       int    `json:"offset"`
			TotalResults int    `json:"totalResults"`
			FuzzyLevel   int    `json:"fuzzyLevel"`
		}{},
		Results: nil,
	}
	mockRepo.EXPECT().GetPlaces("search", "language").Return(&places, nil)

	actualPlaces, appError := service.GetPlaces(request)

	if appError != nil {
		t.Error("failed while validating error for places autocomplete")
	}

	if len(actualPlaces) != 0 {
			t.Error("failed while matching retrieved places")
	}
}

//TODO: add test with filled places (not able to initialize that complex struct for now)
