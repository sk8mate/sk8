package places

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TomtomApiKey string `required:"true" split_words:"true"`
}

var config Config

func init() {
	err := envconfig.Process("sk8_places", &config)

	if err != nil {
		log.Fatal(err.Error())
	}
}
