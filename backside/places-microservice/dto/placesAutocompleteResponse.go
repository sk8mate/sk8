package dto

type PlacesAutocompleteResponseEntry struct {
	Coordinates struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	} `json:"coordinates"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
