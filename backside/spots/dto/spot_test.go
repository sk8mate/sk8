package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_given_spot_without_name_should_return_error(t *testing.T) {
	request := &SpotsAddRequest{
		Address:     "Pawia 5",
		Coordinates: nil,
		Lighting:    true,
		Friendly:    true,
		Verified:    true,
	}

	err := request.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "invalid SpotsAddRequest.Name: value length must be at least 1 runes", err.Error())
}

func Test_given_spot_without_address_should_return_error(t *testing.T) {
	request := &SpotsAddRequest{
		Name:        "Dworzec Glowny Krakow",
		Coordinates: nil,
		Lighting:    true,
		Friendly:    true,
		Verified:    true,
	}

	err := request.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "invalid SpotsAddRequest.Address: value length must be at least 1 runes", err.Error())
}

func Test_given_spot_without_coordinates_should_return_error(t *testing.T) {
	request := &SpotsAddRequest{
		Name:        "Dworzec Glowny Krakow",
		Address:     "Pawia 5",
		Coordinates: nil,
		Lighting:    true,
		Friendly:    true,
		Verified:    true,
	}

	err := request.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "invalid SpotsAddRequest.Coordinates: value is required", err.Error())
}

func Test_given_valid_spot_should_return_success(t *testing.T) {
	request := SpotsAddRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &SpotsAddRequest_Coordinates{
			Lat:  40,
			Long: 60,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}

	err := request.Validate()

	assert.Nil(t, err)
}
