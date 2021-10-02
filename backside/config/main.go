package config

import (
	"log"
	"os"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Env

var Env = struct {
	Development string
	Test        string
	Production  string
}{
	Development: "development",
	Test:        "test",
	Production:  "production",
}

func resolveEnv(e string) string {
	switch e {
	case Env.Development, "":
		return Env.Development
	case Env.Test:
		return Env.Test
	case Env.Production:
		return Env.Production
	default:
		panic("Invalid SK8_ENV value.")
	}
}

var env string

var onceEnv sync.Once

func GetEnv() string {
	onceEnv.Do(func() {
		env = resolveEnv(os.Getenv("SK8_ENV"))
	})

	return env
}

// Config

type Config struct {
	Address    string
	Port       string `default:"8080"`
	DbHost     string `required:"true" split_words:"true"`
	DbPort     string `required:"true" split_words:"true"`
	DbUser     string `required:"true" split_words:"true"`
	DbPassword string `required:"true" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
}

var config Config

var onceConfig sync.Once

func Get() Config {
	onceConfig.Do(func() {
		err := envconfig.Process("sk8", &config)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return config
}
