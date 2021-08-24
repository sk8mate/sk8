package places

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sk8.town/backside/errs"
	"sk8.town/backside/logger"
	"sk8.town/backside/places/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/repository.go -package=mocks sk8.town/backside/places Repository
type Repository interface {
	GetPlaces(search string, language string) (*domain.GetPlacesResponse, *errs.AppError)
}

type DefaultRepository struct {
	tomtomApiKey string
}

func (repo DefaultRepository) ParseLanguage(language string) (string, *errs.AppError) {
	switch language {
	case "pl":
		return "pl-PL", nil
	case "en":
		return "en-US", nil
	default:
		logger.Error(fmt.Sprintf("Could not parse language \"%s\".", language))
		return "", errs.NewUnexpectedError("")
	}
}

func (repo DefaultRepository) GetPlaces(search string, language string) (*domain.GetPlacesResponse, *errs.AppError) {
	language, languageErr := repo.ParseLanguage(language)
	if languageErr != nil {
		return nil, languageErr
	}

	url := fmt.Sprintf("https://api.tomtom.com/search/2/search/%s.json?key=%s&typeahead=true&language=%s", search, repo.tomtomApiKey, language)
	response, err := http.Get(url)

	if err != nil {
		return nil, errs.NewUnexpectedError("")
	}

	var places domain.GetPlacesResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&places)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("")
	}

	return &places, nil
}

func NewRepository(tomtomApiKey string) Repository {
	return DefaultRepository{tomtomApiKey}
}
