package places

import (
	"sk8.town/backside/errs"
	"sk8.town/backside/spots/domain"
	"sk8.town/backside/spots/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/spot_service.go -package=mocks sk8.town/backside/spots SpotService
type SpotService interface {
	Add(*dto.SpotRequest) (*dto.SpotResponse, *errs.AppError)
}

type DefaultSpotService struct {
	spotDb domain.SpotDb
}

func (s DefaultSpotService) Add(request *dto.SpotRequest) (*dto.SpotResponse, *errs.AppError) {
	spotFromRequest := domain.Spot{
		Id:          "",
		Name:        request.Name,
		Address:     request.Address,
		Coordinates: domain.Coordinates{Lat: request.Coordinates.Lat, Long: request.Coordinates.Long},
		Lighting:    request.Lighting,
		Friendly:    request.Friendly,
		Verified:    request.Verified,
	}

	createdSpot, err := s.spotDb.Add(spotFromRequest)
	if err != nil {
		return nil, err
	}

	spotDtoResponse := createdSpot.ToResponseDto()
	return &spotDtoResponse, nil
}

func NewService(db domain.SpotDb) DefaultSpotService {
	return DefaultSpotService{
		spotDb: db,
	}
}
