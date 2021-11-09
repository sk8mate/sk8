package spots

import (
	"sk8.town/backside/errs"
	"sk8.town/backside/spots/domain"
	"sk8.town/backside/spots/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/filters_service.go -package=mocks sk8.town/backside/spots FiltersService
type FiltersService interface {
	GetAll() ([]*dto.FilterData, *errs.AppError)
}

type DefaultFilterService struct {
	filtersRepo domain.FiltersRepository
}

func (s DefaultFilterService) GetAll() ([]*dto.FilterData, *errs.AppError) {
	allFilters, err := s.filtersRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var dtoAllFilters []*dto.FilterData
	for _, filter := range allFilters {
		dtoAllFilters = append(dtoAllFilters, filter.ToDto())
	}

	return dtoAllFilters, nil
}

func NewFilterService(repo domain.FiltersRepository) DefaultFilterService {
	return DefaultFilterService{
		filtersRepo: repo,
	}
}
