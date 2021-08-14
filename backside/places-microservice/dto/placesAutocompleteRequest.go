package dto

import (
	"github.com/sk8mate/sk8/backside/errs"
)

type PlacesAutocompleteRequest struct {
	Search string `json:"search"`
	Language string `json:"language"`
}

func (request PlacesAutocompleteRequest) Validate() *errs.AppError {
	if request.Search == "" {
		return errs.NewValidationError("Search field not provided")
	}

	if request.Language == "" {
		return errs.NewValidationError("Language field not provided")
	}

	return nil
}
