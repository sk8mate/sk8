package domain

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sk8.town/backside/errs"
	"sk8.town/backside/logger"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/spot_repository.go -package=mocks sk8.town/backside/spots/domain SpotRepository
type SpotRepository interface {
	Add(*Spot) (*Spot, *errs.AppError)
}

type SpotDb struct {
	client *gorm.DB
}

func (db SpotDb) Add(spot *Spot) (*Spot, *errs.AppError) {
	db.client.Create(spot)
	db.client.Find(&spot)
	return spot, nil
}

func NewSpotDb(host, port, dbName, user, password string) SpotDb {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	logger.Info(fmt.Sprintf("Connected to database %s on %s:%s", dbName, host, port))

	if db.Migrator().HasTable(&Spot{}) {
		err := db.Migrator().DropTable(&Spot{})
		if err != nil {
			logger.Error(err.Error())
			panic(err)
		}
	}
	err = db.Debug().AutoMigrate(&Spot{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	return SpotDb{db}
}
