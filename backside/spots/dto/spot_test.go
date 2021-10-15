package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_given_add_spot_request_without_name_should_return_error(t *testing.T) {
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

func Test_given_add_spot_request_without_address_should_return_error(t *testing.T) {
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

func Test_given_add_spot_request_without_coordinates_should_return_error(t *testing.T) {
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

func Test_given_add_spot_request_with_invalid_coordinates_should_return_error(t *testing.T) {
	request := &SpotsAddRequest{
		Name:        "Dworzec Glowny Krakow",
		Address:     "Pawia 5",
		Coordinates: &SpotsAddRequest_Coordinates{
			Lat:  700,
			Long: 100,
		},
		Lighting:    true,
		Friendly:    true,
		Verified:    true,
	}

	err := request.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "invalid SpotsAddRequest.Coordinates: embedded message failed validation | caused by: invalid SpotsAddRequest_Coordinates.Lat: value must be inside range [-90, 90]", err.Error())
}

func Test_given_valid_add_spot_request_should_return_success(t *testing.T) {
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

func Test_given_update_spot_request_with_invalid_coordinates_should_return_error(t *testing.T) {
	request := &SpotsUpdateRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &SpotsUpdateRequest_Coordinates{
			Lat:  700,
			Long: 100,
		},
		Lighting: true,
		Friendly: true,
		Verified: true,
	}

	err := request.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "invalid SpotsUpdateRequest.Coordinates: embedded message failed validation | caused by: invalid SpotsUpdateRequest_Coordinates.Lat: value must be inside range [-90, 90]", err.Error())
}

func Test_given_valid_update_spot_request_should_return_success(t *testing.T) {
	request := SpotsUpdateRequest{
		Name:    "Dworzec Glowny Krakow",
		Address: "Pawia 5",
		Coordinates: &SpotsUpdateRequest_Coordinates{
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
