package places

import (
	"sk8.town/backside/errs"
	"sk8.town/backside/places/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/service.go -package=mocks sk8.town/backside/places Service
type Service interface {
	GetPlaces(*dto.AutocompleteRequest) ([]*dto.AutocompleteEntryResponse, *errs.AppError)
}

type DefaultService struct {
	locationService LocationService
	languageParser LanguageParser
}

func (s DefaultService) GetPlaces(request *dto.AutocompleteRequest) ([]*dto.AutocompleteEntryResponse, *errs.AppError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	language, languageErr := s.languageParser.Parse(request.Language)
	if languageErr != nil {
		return nil, languageErr
	}

	places, locationServiceErr := s.locationService.Search(request.Search, language)
	if locationServiceErr != nil {
		return nil, locationServiceErr
	}

	return places.ToDto(), nil
}

func NewService(locationService LocationService) DefaultService {
	return DefaultService{
		locationService: locationService,
	}
}
