package domain

import (
	"gorm.io/gorm"

	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/filter_repository.go -package=mocks sk8.town/backside/spots/domain FilterRepository
type FilterRepository interface {
	GetAllFilterValues() ([]*FilterValue, *errs.AppError)
}

type FilterDb struct {
	client *gorm.DB
}

func (db FilterDb) GetAllFilterValues() ([]*FilterValue, *errs.AppError) {
	var filterValues []*FilterValue
	err := db.client.Preload("Filter").Find(&filterValues).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	if len(filterValues) == 0 {
		return nil, errs.NewNotFoundError(err.Error())
	}
	return filterValues, nil

	//mapping:= make(map[Filter][]*FilterValue)
	//for _, filterValue := range filterValues {
	//	mapping[filterValue.Filter] = append(mapping[filterValue.Filter], &FilterValue{
	//		ID:       filterValue.ID,
	//		Value:    filterValue.Value,
	//	})
	//}
	//
	//for k, v:=range mapping{
	//	fmt.Println(k)
	//	fmt.Println(v[0])
	//	fmt.Println(v[1])
	//}
	//
	////var filterWithFilterValues []*FilterWithFilterValues
	//
	//return nil, nil
}

func NewFilterDb(db *gorm.DB) FilterDb {
	return FilterDb{db}
}
