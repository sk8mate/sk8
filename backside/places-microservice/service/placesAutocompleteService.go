package service

import (
	"github.com/sk8mate/sk8/backside/domain"
	"github.com/sk8mate/sk8/backside/dto"
	"github.com/sk8mate/sk8/backside/errs"
)

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
