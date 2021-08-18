package main

import (
	"sk8.town/backside/places/app"
	"sk8.town/backside/places/logger"
)

func main() {
	logger.Info("Starting places microservice...")
	app.Start()
}
