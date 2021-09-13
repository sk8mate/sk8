package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := "host=localhost user=root password=root dbname=spotsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")

	config := getConfig()
	router := mux.NewRouter()

	places.Make(router)

	handler := cors.Default().Handler(router)

	fmt.Println("Starting server on ", config.Address, ":", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Address, config.Port), handler))
}
