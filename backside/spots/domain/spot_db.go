package domain

import (
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

func NewSpotDb() SpotDb {
	dsn := "host=localhost user=root password=root dbname=spotsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	logger.Info("Connected to database spotsdb on localhost:5432")

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
