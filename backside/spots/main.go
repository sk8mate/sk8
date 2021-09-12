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
		HandleFunc("/spots", handler.GetPlacesAutocomplete).
		Methods(http.MethodPost).
		Name("AddNewSpot")
}
