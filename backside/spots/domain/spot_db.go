package domain

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sk8.town/backside/errs"
	"sk8.town/backside/logger"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/spot_repository.go -package=mocks sk8.town/backside/spots/domain SpotRepository
type SpotRepository interface {
	Add(*Spot) (*Spot, *errs.AppError)
	Get(int) (*Spot, *errs.AppError)
	GetAll() ([]*Spot, *errs.AppError)
	Update(int, *Spot) (*Spot, *errs.AppError)
	Delete(int) (int, *errs.AppError)
}

type SpotDb struct {
	client *gorm.DB
}

func (db SpotDb) Add(spot *Spot) (*Spot, *errs.AppError) {
	err := db.client.Create(spot).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	db.client.Find(&spot)
	return spot, nil
}

func (db SpotDb) Get(int) (*Spot, *errs.AppError) {
	return nil, nil
}

func (db SpotDb) GetAll() ([]*Spot, *errs.AppError) {
	var spots []*Spot
	err := db.client.Model(&Spot{}).Find(&spots).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return spots, nil
}

func (db SpotDb) Update(int, *Spot) (*Spot, *errs.AppError) {
	return nil, nil
}

func (db SpotDb) Delete(int) (int, *errs.AppError) {
	return 0, nil
}

func NewSpotDb(host, port, dbName, user, password string) SpotDb {
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

	return SpotDb{db}
}
