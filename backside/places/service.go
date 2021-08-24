package places

import (
	"sk8.town/backside/errs"
	"sk8.town/backside/places/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/service.go -package=mocks sk8.town/backside/places Service
type Service interface {
	GetPlaces(dto.AutocompleteRequest) ([]dto.AutocompleteEntryResponse, *errs.AppError)
}

type DefaultService struct {
	repository Repository
}

func (s DefaultService) GetPlaces(request dto.AutocompleteRequest) ([]dto.AutocompleteEntryResponse, *errs.AppError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	places, err := s.repository.GetPlaces(request.Search, request.Language)
	if err != nil {
		return nil, err
	}

	return places.ToDto(), nil
}

func NewService(repository Repository) DefaultService {
	return DefaultService{repository}
}
