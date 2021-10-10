package domain

import (
	"strconv"

	"sk8.town/backside/spots/dto"
)

type Spot struct {
	Id          int `gorm:"primary_key;AUTO_INCREMENT"`
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

func (s Spot) ToResponseDto() dto.SpotsAddData {
	return dto.SpotsAddData{Id: strconv.Itoa(s.Id)}
}

func (s Spot) toMap() map[string]interface{} {
	return map[string]interface{}{
		"name":     s.Name,
		"address":  s.Address,
		"lat":      s.Coordinates.Lat,
		"long":     s.Coordinates.Long,
		"lighting": s.Lighting,
		"friendly": s.Friendly,
		"verified": s.Verified,
	}
}
