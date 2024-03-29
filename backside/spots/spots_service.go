package spots

import (
	"sk8.town/backside/errs"
	"sk8.town/backside/spots/domain"
	"sk8.town/backside/spots/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/spots_service.go -package=mocks sk8.town/backside/spots SpotsService
type SpotsService interface {
	Add(*dto.SpotsAddRequest) (*dto.SpotsAddData, *errs.AppError)
	Get(int) (*dto.SpotsGetData, *errs.AppError)
	GetAll() ([]*dto.SpotsGetData, *errs.AppError)
	Update(int, *dto.SpotsUpdateRequest) (*dto.SpotsUpdateData, *errs.AppError)
	Delete(int) *errs.AppError
}

type DefaultSpotService struct {
	spotsRepo domain.SpotsRepository
}

func (s DefaultSpotService) Add(request *dto.SpotsAddRequest) (*dto.SpotsAddData, *errs.AppError) {
	if validationError := request.Validate(); validationError != nil {
		return nil, errs.NewValidationError(validationError.Error())
	}

	spotFromRequest := domain.Spot{
		Name:        request.Name,
		Address:     request.Address,
		Coordinates: domain.Coordinates{Lat: request.Coordinates.Lat, Long: request.Coordinates.Long},
		Lighting:    request.Lighting,
		Friendly:    request.Friendly,
		Verified:    request.Verified,
	}

	createdSpot, err := s.spotsRepo.Add(&spotFromRequest)
	if err != nil {
		return nil, err
	}

	spotDtoResponse := createdSpot.ToAddSpotResponseDto()
	return &spotDtoResponse, nil
}

func (s DefaultSpotService) Get(id int) (*dto.SpotsGetData, *errs.AppError) {
	spot, err := s.spotsRepo.Get(id)
	if err != nil {
		return nil, err
	}

	spotDtoResponse := spot.ToGetSpotResponseDto()
	return &spotDtoResponse, nil
}

func (s DefaultSpotService) GetAll() ([]*dto.SpotsGetData, *errs.AppError) {
	spots, err := s.spotsRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var spotsDtoData []*dto.SpotsGetData

	for _, spot := range spots {
		spotDto := spot.ToGetSpotResponseDto()
		spotsDtoData = append(spotsDtoData, &spotDto)
	}

	return spotsDtoData, nil
}

func (s DefaultSpotService) Delete(id int) *errs.AppError {
	return s.spotsRepo.Delete(id)
}

func (s DefaultSpotService) Update(id int, request *dto.SpotsUpdateRequest) (*dto.SpotsUpdateData, *errs.AppError) {
	if validationError := request.Validate(); validationError != nil {
		return nil, errs.NewValidationError(validationError.Error())
	}

	spotFromRequest := domain.Spot{
		Name:     request.Name,
		Address:  request.Address,
		Lighting: request.Lighting,
		Friendly: request.Friendly,
		Verified: request.Verified,
	}
	if request.Coordinates != nil {
		spotFromRequest.Coordinates.Lat = request.Coordinates.Lat
		spotFromRequest.Coordinates.Long = request.Coordinates.Long
	}

	updatedSpot, err := s.spotsRepo.Update(id, &spotFromRequest)
	if err != nil {
		return nil, err
	}

	spotDtoResponse := updatedSpot.ToUpdateSpotResponseDto()
	return &spotDtoResponse, nil
}

func NewSpotsService(repo domain.SpotsRepository) DefaultSpotService {
	return DefaultSpotService{
		spotsRepo: repo,
	}
}
