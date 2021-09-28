package dto

import (
	"fmt"
	"strings"

	"sk8.town/backside/errs"
)

var supportedLanguages = []string{"pl", "en"} // TODO: should be declared in some global scope when introducing i18n

func validateLanguage(language string) string {
	if language == "" {
		return `Field "language" is required.`
	}

	for _, supportedLanguage := range supportedLanguages {
		if language == supportedLanguage {
			return ""
		}
	}

	return fmt.Sprintf(`Field "language" must be one of: %s.`, strings.Join(supportedLanguages, ", "))
}

func (x *AutocompleteRequest) Validate() *errs.AppError {
	if x.Search == "" {
		return errs.NewValidationError(`Field "search" is required.`)
	}

	if err := validateLanguage(x.Language); err != "" {
		return errs.NewValidationError(err)
	}

	return nil
}
