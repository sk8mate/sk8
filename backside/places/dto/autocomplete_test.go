package dto

import (
	"net/http"
	"testing"
)

func Test_given_empty_search_field_should_return_error(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "",
		Language: "language",
	}

	appError := request.Validate()

	if appError.Message != "Search field not provided" {
		t.Error("Invalid message while testing places autocomplete request - language field")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing places autocomplete request - language field")
	}
}

func Test_given_empty_language_field_should_return_error(t *testing.T) {
	request := AutocompleteRequest{
		Search:   "warsaw",
		Language: "",
	}

	appError := request.Validate()

	if appError.Message != "Language field not provided" {
		t.Error("Invalid message while testing places autocomplete request - language field")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing places autocomplete request - language field")
	}
}
