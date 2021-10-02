package places

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sk8.town/backside/errs"
	"sk8.town/backside/logger"
	"sk8.town/backside/places/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/location_service.go -package=mocks sk8.town/backside/places LocationService
type LocationService interface {
	Search(search, language string) (*domain.GetPlacesResponse, *errs.AppError)
}

const baseUrl = "https://api.tomtom.com"

type TomTomLocationService struct {
	tomtomApiKey string
}

func (s TomTomLocationService) createUrl(search, language string) string {
	return fmt.Sprintf("%s/search/2/search/%s.json?key=%s&typeahead=true&language=%s", baseUrl, search, s.tomtomApiKey, language)
}

func (s TomTomLocationService) decodePlacesFromResponse(response *http.Response) (*domain.GetPlacesResponse, *errs.AppError) {
	var places domain.GetPlacesResponse
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&places)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("")
	}
	return &places, nil
}

func (s TomTomLocationService) Search(search, language string) (*domain.GetPlacesResponse, *errs.AppError) {
	tomTomRequestUrl := s.createUrl(search, language)
	response, err := http.Get(tomTomRequestUrl)
	if err != nil {
		return nil, errs.NewUnexpectedError("")
	}

	places, decodeErr := s.decodePlacesFromResponse(response)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return places, nil
}

func NewLocationService(tomtomApiKey string) TomTomLocationService {
	return TomTomLocationService{tomtomApiKey: tomtomApiKey}
}
