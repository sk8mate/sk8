package domain

import (
	"gorm.io/gorm"

	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/filter_repository.go -package=mocks sk8.town/backside/spots/domain FilterRepository
type FilterRepository interface {
	GetAll() ([]*Filter, *errs.AppError)
}

type FilterDb struct {
	client *gorm.DB
}

func (db FilterDb) GetAll() ([]*Filter, *errs.AppError) {
	var filters []*Filter

	err := db.client.Preload("FilterValues").Find(&filters).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if len(filters) == 0 {
		return nil, errs.NewNotFoundError(err.Error())
	}

	return filters, nil
}

func NewFilterDb(db *gorm.DB) FilterDb {
	return FilterDb{db}
}
