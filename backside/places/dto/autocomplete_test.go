package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_given_empty_search_param_should_return_error(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "",
		Language: "pl",
	}

	err := request.Validate()

	assert.Equal(t, "invalid AutocompleteRequest.Search: value length must be at least 1 runes", err.Error())
}

func Test_given_empty_language_param_should_return_error(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "warsaw",
		Language: "",
	}

	err := request.Validate()

	assert.Equal(t, "invalid AutocompleteRequest.Language: value must be in list [pl en]", err.Error())
}

func Test_given_invalid_language_should_return_error(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "warsaw",
		Language: "fr",
	}

	err := request.Validate()

	assert.Regexp(t, "", err.Error())
}

func Test_given_valid_params_should_return_success(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "warszawa",
		Language: "pl",
	}

	err := request.Validate()

	assert.Nil(t, err)
}

func Test_given_valid_params_should_return_success_2(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "cracow",
		Language: "en",
	}

	err := request.Validate()

	assert.Nil(t, err)
}
