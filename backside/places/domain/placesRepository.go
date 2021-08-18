package domain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sk8.town/backside/places/errs"
	"sk8.town/backside/places/logger"
)

type DefaultPlacesRepository struct {
	apiKey string
}

func (repo DefaultPlacesRepository) GetPlaces(search string, language string) (*Places, *errs.AppError) {
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

func NewPlacesRepository(apiKey string) PlacesRepository {
	return DefaultPlacesRepository{apiKey}
}
