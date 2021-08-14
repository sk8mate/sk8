package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sk8mate/sk8/backside/places-microservice/domain"
	"github.com/sk8mate/sk8/backside/places-microservice/service"
	"log"
	"net/http"
	"os"
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
