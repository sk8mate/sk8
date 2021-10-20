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
	Delete(int) *errs.AppError
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

func (db SpotDb) Get(id int) (*Spot, *errs.AppError) {
	var spot Spot
	err := db.client.Where("id = ?", id).Take(&spot).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &spot, nil
}

func (db SpotDb) GetAll() ([]*Spot, *errs.AppError) {
	var spots []*Spot
	err := db.client.Find(&spots).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return spots, nil
}

func (db SpotDb) Update(id int, spotToUpdateWith *Spot) (*Spot, *errs.AppError) {
	err := db.client.Where("id = ?", id).Take(&Spot{}).UpdateColumns(spotToUpdateWith).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return db.Get(id)
}

func (db SpotDb) Delete(id int) *errs.AppError {
	if _, err := db.Get(id); err != nil {
		return errs.NewBadRequestError(fmt.Sprintf(`Spot with id %d does not exist`, id))
	}

	err := db.client.Delete(&Spot{}, id).Error
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
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
