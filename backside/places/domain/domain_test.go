package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sk8.town/backside/places/dto"
)

func Test_given_get_places_response_with_geography_type_should_set_name_as_free_form_address_and_empty_address(t *testing.T) {
	domainResponse := GetPlacesResponse{
		Results: []Result{
			{
				Type: "Geography",
				Position: Position{
					Lat: 20,
					Lon: 30,
				},
				Address: Address{
					FreeformAddress: "Free form address",
				},
			},
		},
	}

	actualDtoResponse := domainResponse.ToDto()

	var expectedDtoResponse = &dto.AutocompleteEntry{
		Coordinates: &dto.Coordinates{
			Lat:  20,
			Long: 30,
		},
		Name:    "Free form address",
		Address: "",
	}
	assert.Equal(t, 1, len(actualDtoResponse))
	assert.Equal(t, expectedDtoResponse, actualDtoResponse[0])
}

func Test_given_get_places_response_with_poi_type_should_set_name_as_poi_name_and_address_as_address(t *testing.T) {
	domainResponse := GetPlacesResponse{
		Results: []Result{
			{
				Type: "POI",
				Position: Position{
					Lat: 3,
					Lon: 4,
				},
				Address: Address{
					FreeformAddress: "Free form address",
				},
				Poi: Poi{
					Name: "Poi name",
				},
			},
		},
	}

	actualDtoResponse := domainResponse.ToDto()

	var expectedDtoResponse = &dto.AutocompleteEntry{
		Coordinates: &dto.Coordinates{
			Lat:  3,
			Long: 4,
		},
		Name:    "Poi name",
		Address: "Free form address",
	}
	assert.Equal(t, 1, len(actualDtoResponse))
	assert.Equal(t, expectedDtoResponse, actualDtoResponse[0])
}
