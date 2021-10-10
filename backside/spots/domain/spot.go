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

func (s Spot) ToAddSpotResponseDto() dto.SpotsAddData {
	return dto.SpotsAddData{Id: strconv.Itoa(s.Id)}
}

func (s Spot) ToGetSpotResponseDto() dto.SpotsGetData {
	return dto.SpotsGetData{
		Id:      strconv.Itoa(s.Id),
		Name:    s.Name,
		Address: s.Address,
		Coordinates: &dto.ResponseCoordinates{
			Lat:  s.Coordinates.Lat,
			Long: s.Coordinates.Long,
		},
		Lighting: s.Lighting,
		Friendly: s.Friendly,
		Verified: s.Verified,
	}
}

func (s Spot) ToUpdateSpotResponseDto() dto.SpotsUpdateData {
	return dto.SpotsUpdateData{
		Id:      strconv.Itoa(s.Id),
		Name:    s.Name,
		Address: s.Address,
		Coordinates: &dto.ResponseCoordinates{
			Lat:  s.Coordinates.Lat,
			Long: s.Coordinates.Long,
		},
		Lighting: s.Lighting,
		Friendly: s.Friendly,
		Verified: s.Verified,
	}
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
