package domain

import (
	"fmt"

	"gorm.io/gorm"

	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/spots_repository.go -package=mocks sk8.town/backside/spots/domain SpotsRepository
type SpotsRepository interface {
	Add(*Spot) (*Spot, *errs.AppError)
	Get(int) (*Spot, *errs.AppError)
	GetAll() ([]*Spot, *errs.AppError)
	Update(int, *Spot) (*Spot, *errs.AppError)
	Delete(int) *errs.AppError
}

type SpotsDb struct {
	client *gorm.DB
}

func (db SpotsDb) Add(spot *Spot) (*Spot, *errs.AppError) {
	err := db.client.Create(spot).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	db.client.Find(&spot)
	return spot, nil
}

func (db SpotsDb) Get(id int) (*Spot, *errs.AppError) {
	var spot Spot
	err := db.client.Where("id = ?", id).Take(&spot).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &spot, nil
}

func (db SpotsDb) GetAll() ([]*Spot, *errs.AppError) {
	var spots []*Spot
	err := db.client.Find(&spots).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return spots, nil
}

func (db SpotsDb) Update(id int, spotToUpdateWith *Spot) (*Spot, *errs.AppError) {
	err := db.client.Where("id = ?", id).Take(&Spot{}).UpdateColumns(spotToUpdateWith).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return db.Get(id)
}

func (db SpotsDb) Delete(id int) *errs.AppError {
	if _, err := db.Get(id); err != nil {
		return errs.NewBadRequestError(fmt.Sprintf(`Spot with id %d does not exist`, id))
	}

	err := db.client.Delete(&Spot{}, id).Error
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func NewSpotsDb(db *gorm.DB) SpotsDb {
	return SpotsDb{db}
}
