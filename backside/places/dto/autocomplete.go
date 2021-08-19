package dto

import (
	"sk8.town/backside/errors"
)

type AutocompleteRequest struct {
	Search   string `json:"search"`
	Language string `json:"language"`
}

func (request AutocompleteRequest) Validate() *errors.AppError {
	if request.Search == "" {
		return errors.NewValidationError("Search field not provided")
	}

	if request.Language == "" {
		return errors.NewValidationError("Language field not provided")
	}

	return nil
}

type AutocompleteEntryResponse struct {
	Coordinates struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	} `json:"coordinates"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
