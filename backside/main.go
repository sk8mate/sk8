package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"sk8.town/backside/auth"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"sk8.town/backside/config"
	"sk8.town/backside/places"
	"sk8.town/backside/spots"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	cfg := config.Get()
	env := config.GetEnv()
	router := mux.NewRouter()

	places.Make(router)
	spots.Make(router)
	auth.Make(router)

	handler := cors.Default().Handler(router)
	fmt.Printf("Starting server on %s:%s in %s mode.\n", cfg.Address, cfg.Port, env)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Address, cfg.Port), handler))
}
