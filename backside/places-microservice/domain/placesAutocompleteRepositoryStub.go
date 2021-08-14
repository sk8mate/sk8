package domain

import (
	"github.com/sk8mate/sk8/backside/places-microservice/errs"
)

type PlacesAutocompleteRepositoryStub struct {
}

func (stub PlacesAutocompleteRepositoryStub) GetPlaces(search string, language string) (*Places, *errs.AppError) {
	var places Places
	return &places, nil
}

func NewPlacesAutocompleteRepositoryStub() PlacesAutocompleteRepositoryStub {
	return PlacesAutocompleteRepositoryStub{}
}