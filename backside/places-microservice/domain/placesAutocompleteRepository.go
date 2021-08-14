package domain

import (
	"encoding/json"
	"fmt"
	"github.com/sk8mate/sk8/backside/places-microservice/errs"
	"github.com/sk8mate/sk8/backside/places-microservice/logger"
	"net/http"
)

type PlacesAutocompleteRepository struct {
	apiKey string
}

func (repo PlacesAutocompleteRepository) GetPlaces(search string, language string) (*Places, *errs.AppError) {
	url := fmt.Sprintf("https://api.tomtom.com/search/2/search/%s.json?key=%s&typeahead=true&language=%s", search, repo.apiKey, language)
	response, err := http.Get(url)
	if response == nil {
		return nil, nil
	}
	var places Places
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&places)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil
	}

	return &places, nil
}

func NewPlacesAutocompleteRepository(apiKey string) PlacesAutocompleteRepository {
	return PlacesAutocompleteRepository{apiKey}
}
