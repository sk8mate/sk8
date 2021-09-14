package places

import (
	"github.com/gorilla/mux"
	"net/http"
	"sk8.town/backside/spots/domain"
)

func Make(router *mux.Router) {
	config := getConfig()

	spotsDb := domain.NewSpotDb()
	spotsService :=
	handler := Handler{NewService(locationService)}

	router.
		HandleFunc("/spots", handler.GetPlacesAutocomplete).
		Methods(http.MethodPost).
		Name("AddNewSpot")
}
