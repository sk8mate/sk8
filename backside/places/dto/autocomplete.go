package dto

import (
	"fmt"
	"strings"

	"sk8.town/backside/errs"
)

type AutocompleteRequest struct {
	Search   string `json:"search"`
	Language string `json:"language"`
}

func validateLanguage(language string) string {
	if language == "" {
		return "Field \"language\" is required."
	}

	var languages = []string{"pl", "en"} // TODO: should be declared in some global scope when introducing i18n
	for _, l := range languages {
		if l == language {
			return ""
		}
	}

	return fmt.Sprintf("Field \"language\" must be one of: %s.", strings.Join(languages, ", "))
}

func (request AutocompleteRequest) Validate() *errs.AppError {
	if request.Search == "" {
		return errs.NewValidationError("Field \"search\" is required.")
	}

	if err := validateLanguage(request.Language); err != "" {
		return errs.NewValidationError(err)
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
