package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"

	"sk8.town/backside/places"
)

type Config struct {
	Address string `default:"localhost"`
	Port    int    `default:"8080"`
}

func getConfig() Config {
	var config Config
	err := envconfig.Process("sk8", &config)

	if err != nil {
		log.Fatal(err.Error())
	}

	return config
}

func main() {
	config := getConfig()
	router := mux.NewRouter()

	places.Make(router)

	fmt.Printf("Starting server on %s:%d", config.Address, config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Address, config.Port), router))
}
