package places

import (
	"sk8.town/backside/errors"
	"sk8.town/backside/places/dto"
)

type Service interface {
	GetPlaces(dto.AutocompleteRequest) ([]dto.AutocompleteEntryResponse, *errors.AppError)
}

type DefaultService struct {
	repository PlacesRepository
}

func (s DefaultService) GetPlaces(request dto.AutocompleteRequest) ([]dto.AutocompleteEntryResponse, *errors.AppError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	places, err := s.repository.GetPlaces(request.Search, request.Language)
	if err != nil {
		return nil, err
	}

	return places.ToDto(), nil
}

func NewService(repository PlacesRepository) DefaultService {
	return DefaultService{repository}
}
