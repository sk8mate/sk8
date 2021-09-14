package spots

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TomtomApiKey string `required:"true" split_words:"true"`
}

func getConfig() Config {
	var config Config

	err := envconfig.Process("sk8_places", &config)

	if err != nil {
		log.Fatal(err.Error())
	}

	return config
}
