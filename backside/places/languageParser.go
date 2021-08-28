package places

import (
	"fmt"
	"sk8.town/backside/errs"
	"sk8.town/backside/logger"
)

type LanguageParser struct{}

func (s LanguageParser) ParseLanguage(language string) (string, *errs.AppError) {
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
