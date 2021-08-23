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

func (places GetPlacesResponse) ToDto() []dto.AutocompleteEntryResponse {
	responseEntries := make([]dto.AutocompleteEntryResponse, 0)

	for _, place := range places.Results {
		if place.Type == "Geography" {
			entry := dto.AutocompleteEntryResponse{
				Coordinates: struct {
					Lat  float64 `json:"lat"`
					Long float64 `json:"long"`
				}{place.Position.Lat, place.Position.Lon},
				Name:    place.Address.FreeformAddress,
				Address: "",
			}
			responseEntries = append(responseEntries, entry)
		} else if place.Type == "POI" {
			entry := dto.AutocompleteEntryResponse{
				Coordinates: struct {
					Lat  float64 `json:"lat"`
					Long float64 `json:"long"`
				}{place.Position.Lat, place.Position.Lon},
				Name:    place.Poi.Name,
				Address: place.Address.FreeformAddress,
			}
			responseEntries = append(responseEntries, entry)
		}
	}
	return responseEntries
}
