package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"sk8.town/backside/places/domain"
	"sk8.town/backside/places/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("TOM_TOM_API_KEY") == "" {
		log.Fatal("Environment variables not defined")
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	placesApiKey := os.Getenv("TOM_TOM_API_KEY")
	placesAutocompleteRepository := domain.NewPlacesRepository(placesApiKey)
	placesHandlers := PlacesHandlers{service.NewPlacesService(placesAutocompleteRepository)}
	router.HandleFunc("/places/autocomplete", placesHandlers.getPlacesByAutocomplete).Methods(http.MethodGet).Name("GetPlacesAutocomplete")

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
