package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"

	"sk8.town/backside/places"
)

type Config struct {
	Address string
	Port    string `default:"8080"`
}

func getConfig() Config {
	var config Config
	err := envconfig.Process("sk8", &config)

	if err != nil {
		log.Println(err.Error())
	}

	return config
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	config := getConfig()
	router := mux.NewRouter()

	places.Make(router)
	spots.Make(router)

	handler := cors.Default().Handler(router)

	fmt.Println("Starting server on ", config.Address, ":", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Address, config.Port), handler))
}
