package places

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sk8.town/backside/errors"
	"sk8.town/backside/logger"
	"sk8.town/backside/places/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/repository.go -package=mocks sk8.town/backside/places PlacesRepository
type PlacesRepository interface {
	GetPlaces(search string, language string) (*domain.GetPlacesResponse, *errors.AppError)
}

type Repository struct {
	apiKey string
}

func (repo Repository) ParseLanguage(language string) (string, *errors.AppError) {
	switch language {
	case "pl":
		return "pl-PL", nil
	case "en":
		return "en-US", nil
	default:
		logger.Error(fmt.Sprintf("Could not parse language \"%s\".", language))
		return "", errors.NewUnexpectedError("")
	}
}

func (repo Repository) GetPlaces(search string, language string) (*domain.GetPlacesResponse, *errors.AppError) {
	language, languageErr := repo.ParseLanguage(language)
	if languageErr != nil {
		return nil, languageErr
	}

	url := fmt.Sprintf("https://api.tomtom.com/search/2/search/%s.json?key=%s&typeahead=true&language=%s", search, repo.apiKey, language)
	response, err := http.Get(url)

	if err != nil {
		return nil, errors.NewUnexpectedError("")
	}

	var places domain.GetPlacesResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&places)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.NewUnexpectedError("")
	}

	return &places, nil
}

func NewRepository() Repository {
	return Repository{config.TomtomApiKey}
}
