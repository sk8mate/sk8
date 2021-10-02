package places

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var once sync.Once

type Config struct {
	TomtomApiKey string `required:"true" split_words:"true"`
}

var config Config

func getConfig() Config {
	once.Do(func() {
		err := envconfig.Process("sk8_places", &config)

		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return config
}
