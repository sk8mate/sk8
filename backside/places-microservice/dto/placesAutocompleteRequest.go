package dto

type PlacesAutocompleteRequest struct {
	Search string `json:"search"`
	Language string `json:"language"`
}
