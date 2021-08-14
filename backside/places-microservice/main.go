package main

import (
	"github.com/sk8mate/sk8/backside/app"
	"github.com/sk8mate/sk8/backside/logger"
)

func main() {
	logger.Info("Starting places microservice...")
	app.Start()
}

