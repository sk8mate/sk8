package main

import (
	"github.com/sk8mate/sk8/backside/places-microservice/app"
	"github.com/sk8mate/sk8/backside/places-microservice/logger"
)

func main() {
	logger.Info("Starting places microservice...")
	app.Start()
}

