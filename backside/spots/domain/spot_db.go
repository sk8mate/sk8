package domain

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/spot_repository.go -package=mocks sk8.town/backside/spots SpotRepository
type SpotRepository interface {
	Add(Spot) (*Spot, *errs.AppError)
}

type SpotDb struct {
	client *gorm.DB
}

func (repo SpotDb) Add(spot Spot) (*Spot, *errs.AppError) {
	return nil, nil
}

func NewSpotDb() SpotDb {
	dsn := "host=localhost user=root password=root dbname=spotsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	return SpotDb{db}
}

