package service

import (
	"github.com/sk8mate/sk8/backside/places-microservice/domain"
	"github.com/sk8mate/sk8/backside/places-microservice/dto"
	"github.com/sk8mate/sk8/backside/places-microservice/errs"
)

//go:generate mockgen -destination=../mocks/service/mockPlacesAutocompleteService.go -package=service github.com/sk8mate/sk8/backside/places-microservice/service PlacesAutocompleteService
type PlacesAutocompleteService interface {
	GetPlaces(dto.PlacesAutocompleteRequest) ([]dto.PlacesAutocompleteResponseEntry, *errs.AppError)
}

type DefaultPlacesAutocompleteService struct {
	repository domain.PlacesAutocompleteRepository
}

func (s DefaultPlacesAutocompleteService) GetPlaces(request dto.PlacesAutocompleteRequest) ([]dto.PlacesAutocompleteResponseEntry, *errs.AppError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	places, err := s.repository.GetPlaces(request.Search, request.Language)
	if err != nil {
		return nil, err
	}

	return places.ToDto(), nil
}

func NewPlacesAutocompleteService(repository domain.PlacesAutocompleteRepository) DefaultPlacesAutocompleteService {
	return DefaultPlacesAutocompleteService{repository}
}
