package dto

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"testing"
)

func Test_given_spot_without_coordinates_should_return_error(t *testing.T) {
	request := SpotRequest{
		Name:        "Dworzec Glowny Krakow",
		Address:     "Pawia 5",
		Coordinates: nil,
		Lighting:    true,
		Friendly:    true,
		Verified:    true,
	}

	appError := request.Validate()

	assert.Equal(t, http.StatusUnprocessableEntity, appError.Code)
	assert.Regexp(t, regexp.MustCompile(`Coordinates not provided`), appError.Message)
}

func Test_given_valid_spot_should_return_success(t *testing.T) {
	request := SpotRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &SpotCoordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}

	appError := request.Validate()

	assert.Nil(t, appError)
}
