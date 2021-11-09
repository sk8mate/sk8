package domain

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sk8.town/backside/logger"
)

func Connect(host, port, dbName, user, password string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw",
		host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	logger.Info(fmt.Sprintf("Connected to database %s on %s:%s", dbName, host, port))

	if os.Getenv("SK8_DB_DROP_ALL_TABLES") == "true" {
		dropTables(db)
	}
	autoMigrate(db)

	return db
}
