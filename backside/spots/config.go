package spots

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbHost     string `required:"true" split_words:"true"`
	DbPort     string `required:"true" split_words:"true"`
	DbUser     string `required:"true" split_words:"true"`
	DbPassword string `required:"true" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
}

func GetConfig() Config {
	var config Config

	err := envconfig.Process("sk8_spots", &config)

	if err != nil {
		log.Fatal(err.Error())
	}

	return config
}
