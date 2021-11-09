package domain

import (
	"gorm.io/gorm"

	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/filters_repository.go -package=mocks sk8.town/backside/spots/domain FiltersRepository
type FiltersRepository interface {
	GetAll() ([]*Filter, *errs.AppError)
}

type FiltersDb struct {
	client *gorm.DB
}

func (db FiltersDb) GetAll() ([]*Filter, *errs.AppError) {
	var filters []*Filter

	err := db.client.Preload("FilterValues").Find(&filters).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return filters, nil
}

func NewFiltersDb(db *gorm.DB) FiltersDb {
	return FiltersDb{db}
}
