package domain

import "sk8.town/backside/spots/dto"

type Spot struct {
	Id          string      `gorm:"primary_key"`
	Name        string
	Address     string
	Coordinates Coordinates `gorm:"embedded"`
	Lighting    bool
	Friendly    bool
	Verified    bool
}

type Coordinates struct {
	Lat  float64 `gorm:"lat"`
	Long float64 `gorm:"long"`
}

func (s Spot) ToResponseDto() dto.SpotResponse{
	return dto.SpotResponse{Id: s.Id}
}
