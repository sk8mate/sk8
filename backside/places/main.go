package places

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Make(router *mux.Router) {
	config := getConfig()

	locationService := NewLocationService(config.TomtomApiKey)
	handler := Handler{NewService(locationService)}

	router.
		HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete).
		Methods(http.MethodGet).
		Name("GetPlacesAutocomplete")
}
