package domain

import (
	"github.com/sk8mate/sk8/backside/places-microservice/dto"
	"github.com/sk8mate/sk8/backside/places-microservice/errs"
)

type Places struct {
	Summary struct {
		Query        string `json:"query"`
		QueryType    string `json:"queryType"`
		QueryTime    int    `json:"queryTime"`
		NumResults   int    `json:"numResults"`
		Offset       int    `json:"offset"`
		TotalResults int    `json:"totalResults"`
		FuzzyLevel   int    `json:"fuzzyLevel"`
	} `json:"summary"`
	Results []struct {
		Type       string  `json:"type"`
		ID         string  `json:"id"`
		Score      float64 `json:"score"`
		EntityType string  `json:"entityType,omitempty"`
		Address    struct {
			Municipality                string `json:"municipality"`
			CountrySecondarySubdivision string `json:"countrySecondarySubdivision"`
			CountrySubdivision          string `json:"countrySubdivision"`
			CountryCode                 string `json:"countryCode"`
			Country                     string `json:"country"`
			CountryCodeISO3             string `json:"countryCodeISO3"`
			FreeformAddress             string `json:"freeformAddress"`
		} `json:"address"`
		Position struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"position"`
		Viewport struct {
			TopLeftPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"topLeftPoint"`
			BtmRightPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"btmRightPoint"`
		} `json:"viewport"`
		BoundingBox struct {
			TopLeftPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"topLeftPoint"`
			BtmRightPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"btmRightPoint"`
		} `json:"boundingBox,omitempty"`
		DataSources struct {
			Geometry struct {
				ID string `json:"id"`
			} `json:"geometry"`
		} `json:"dataSources,omitempty"`
		Info string `json:"info,omitempty"`
		Poi  struct {
			Name        string `json:"name"`
			Phone       string `json:"phone"`
			CategorySet []struct {
				ID int `json:"id"`
			} `json:"categorySet"`
			URL             string   `json:"url"`
			Categories      []string `json:"categories"`
			Classifications []struct {
				Code  string `json:"code"`
				Names []struct {
					NameLocale string `json:"nameLocale"`
					Name       string `json:"name"`
				} `json:"names"`
			} `json:"classifications"`
		} `json:"poi,omitempty"`
		EntryPoints []struct {
			Type     string `json:"type"`
			Position struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"position"`
		} `json:"entryPoints,omitempty"`
	} `json:"results"`
}

func (places Places) ToDto() []dto.PlacesAutocompleteResponseEntry {
	responseEntries := make([]dto.PlacesAutocompleteResponseEntry, 0)

	for _, place := range places.Results {
		if place.Type == "Geography" {
			entry := dto.PlacesAutocompleteResponseEntry{
				Coordinates: struct {
					Lat  float64 `json:"lat"`
					Long float64 `json:"long"`
				}{place.Position.Lat, place.Position.Lon},
				Name:    place.Address.FreeformAddress,
				Address: "",
			}
			responseEntries = append(responseEntries, entry)
		} else if place.Type == "POI" {
			entry := dto.PlacesAutocompleteResponseEntry{
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

//go:generate mockgen -destination=../mocks/domain/mockPlacesRepository.go -package=domain github.com/sk8mate/sk8/backside/places-microservice/domain PlacesRepository
type PlacesRepository interface {
	GetPlaces(search string, language string) (*Places, *errs.AppError)
}
