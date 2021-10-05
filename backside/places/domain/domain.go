package domain

import (
	"sk8.town/backside/places/dto"
)

type Position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Poi struct {
	Name string `json:"name"`
}

type Address struct {
	FreeformAddress string `json:"freeformAddress"`
}

type Result struct {
	Type     string   `json:"type"`
	Address  Address  `json:"address"`
	Position Position `json:"position"`
	Poi      Poi      `json:"poi,omitempty"`
}

type GetPlacesResponse struct {
	Results []Result `json:"results"`
}

func (places GetPlacesResponse) ToDto() []*dto.AutocompleteEntry {
	responseEntries := make([]*dto.AutocompleteEntry, 0)

	for _, place := range places.Results {
		coordinates := dto.Coordinates{
			Lat:  place.Position.Lat,
			Long: place.Position.Lon,
		}

		if place.Type == "Geography" {
			entry := dto.AutocompleteEntry{
				Coordinates: &coordinates,
				Name:        place.Address.FreeformAddress,
				Address:     "",
			}
			responseEntries = append(responseEntries, &entry)
		} else if place.Type == "POI" {
			entry := dto.AutocompleteEntry{
				Coordinates: &coordinates,
				Name:        place.Poi.Name,
				Address:     place.Address.FreeformAddress,
			}
			responseEntries = append(responseEntries, &entry)
		}
	}
	return responseEntries
}
