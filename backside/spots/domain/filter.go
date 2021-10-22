package domain

import (
	"sk8.town/backside/spots/dto"
	"strconv"
)

type Filter struct {
	ID           int `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string
	Type         string
	FilterValues []FilterValue
}

type FilterValue struct {
	ID       int `gorm:"primary_key;AUTO_INCREMENT"`
	Value    string
	FilterID int
}

func (f Filter) ToDto() *dto.FilterData {
	var dtoFilterValues []*dto.FilterValueData

	for _, filterValue := range f.FilterValues {
		dtoFilterValues = append(dtoFilterValues, filterValue.ToDto())
	}

	return &dto.FilterData{
		Id:     strconv.Itoa(f.ID),
		Name:   f.Name,
		Type:   f.Type,
		Values: dtoFilterValues,
	}
}

func (fv FilterValue) ToDto() *dto.FilterValueData {
	return &dto.FilterValueData{
		Id:   strconv.Itoa(fv.ID),
		Name: fv.Value,
	}
}
