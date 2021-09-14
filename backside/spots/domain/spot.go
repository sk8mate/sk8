package domain

import "sk8.town/backside/spots/dto"

type Spot struct {
	Id          string      `gorm:"primary_key"`
	Name        string      `gorm:"name"`
	Address     string      `gorm:"address"`
	Coordinates Coordinates `gorm:"coordinates"`
	Lighting    bool        `gorm:"lighting"`
	Friendly    bool        `gorm:"friendly"`
	Verified    bool        `gorm:"verified"`
}

type Coordinates struct {
	Lat  float64 `gorm:"lat"`
	Long float64 `gorm:"long"`
}

func (s Spot) ToResponseDto() dto.SpotResponse{
	return dto.SpotResponse{Id: s.Id}
}
