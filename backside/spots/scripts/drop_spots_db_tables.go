package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sk8.town/backside/logger"
	"sk8.town/backside/spots"
	"sk8.town/backside/spots/domain"
)

const (
	SPOTS_DB_NAME = "spotsDb"
)

func dropSpotsDbTables(host, port, user, password string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw",
		host, user, password, SPOTS_DB_NAME, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	logger.Info(fmt.Sprintf("Connected to database %s on %s:%s", SPOTS_DB_NAME, host, port))

	if db.Migrator().HasTable(&domain.Spot{}) {
		err := db.Migrator().DropTable(&domain.Spot{})
		if err != nil {
			logger.Error(err.Error())
			panic(err)
		}
		logger.Info("Dropped spots table")
	} else {
		logger.Info("Spots db does not contain spots table")
	}
}

func main () {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	config := spots.GetConfig()
	dropSpotsDbTables(config.DbHost, config.DbPort, config.DbUser, config.DbPassword)
}
