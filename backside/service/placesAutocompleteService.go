package service

import (
	"github.com/sk8mate/sk8/backside/domain"
	"github.com/sk8mate/sk8/backside/dto"
	"github.com/sk8mate/sk8/backside/errs"
)

type PlacesAutocompleteService interface {
	GetPlaces(search string, language string) ([]dto.PlacesAutocompleteResponseEntry, *errs.AppError)
}

type DefaultPlacesAutocompleteService struct {
	repository domain.PlacesAutocompleteRepository
}

func (s DefaultPlacesAutocompleteService) GetPlaces(search string, language string) ([]dto.PlacesAutocompleteResponseEntry, *errs.AppError) {
	places, err := s.repository.GetPlaces(search, language)
	if err != nil {
		return nil, err
	}

	return places.ToDto(), nil
}

func NewPlacesAutocompleteService(repository domain.PlacesAutocompleteRepository) DefaultPlacesAutocompleteService {
	return DefaultPlacesAutocompleteService{repository}
}
