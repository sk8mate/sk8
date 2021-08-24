package dto

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_given_empty_search_param_should_return_error(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "",
		Language: "pl",
	}

	appError := request.Validate()

	assert.Equal(t, appError.Code, http.StatusUnprocessableEntity)
	assert.Regexp(t, regexp.MustCompile("\"search\" is required"), appError.Message)
}

func Test_given_empty_language_param_should_return_error(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "warsaw",
		Language: "",
	}

	appError := request.Validate()

	assert.Equal(t, appError.Code, http.StatusUnprocessableEntity)
	assert.Regexp(t, regexp.MustCompile("\"language\" is required"), appError.Message)
}

func Test_given_invalid_language_should_return_error(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "warsaw",
		Language: "fr",
	}

	appError := request.Validate()

	assert.Equal(t, appError.Code, http.StatusUnprocessableEntity)
	assert.Regexp(t, regexp.MustCompile("must be one of: pl, en"), appError.Message)
}

func Test_given_valid_params_should_return_success(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "warszawa",
		Language: "pl",
	}

	appError := request.Validate()

	assert.Nil(t, appError)
}

func Test_given_valid_params_should_return_success_2(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "cracow",
		Language: "en",
	}

	appError := request.Validate()

	assert.Nil(t, appError)
}
