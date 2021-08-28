package places

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var parser LanguageParser

func Test_given_supported_language_should_return_parsed_language_1(t *testing.T) {
	language := "pl"

	parsedLanguage, err := parser.ParseLanguage(language)

	assert.Nil(t, err)
	assert.Equal(t, "pl-PL", parsedLanguage)
}

func Test_given_supported_language_should_return_parsed_language_2(t *testing.T) {
	language := "en"

	parsedLanguage, err := parser.ParseLanguage(language)

	assert.Nil(t, err)
	assert.Equal(t, "en-US", parsedLanguage)
}

func Test_given_not_supported_language_should_return_error(t *testing.T) {
	languageAbbreviation := "fr"

	_, err := parser.ParseLanguage(languageAbbreviation)

	assert.Equal(t, http.StatusInternalServerError, err.Code)
}