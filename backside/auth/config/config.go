package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var once sync.Once

type Config struct {
	GoogleClientId     string `required:"true" split_words:"true"`
	GoogleClientSecret string `required:"true" split_words:"true"`
}

var config Config

func GetConfig() Config {
	once.Do(func() {
		err := envconfig.Process("sk8_auth", &config)

		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return config
}
