package spots

import (
	"sk8.town/backside/errs"
	"sk8.town/backside/spots/domain"
	"sk8.town/backside/spots/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/filter_service.go -package=mocks sk8.town/backside/spots FilterService
type FilterService interface {
	GetAll() ([]*dto.FilterData, *errs.AppError)
}

type DefaultFilterService struct {
	filterRepo domain.FilterRepository
}

func (s DefaultFilterService) GetAll() ([]*dto.FilterData, *errs.AppError) {
	allFilters, err := s.filterRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var dtoAllFilters []*dto.FilterData
	for _, filter := range allFilters {
		dtoAllFilters = append(dtoAllFilters, filter.ToDto())
	}

	return dtoAllFilters, nil
}

func NewFilterService(repo domain.FilterRepository) DefaultFilterService {
	return DefaultFilterService{
		filterRepo: repo,
	}
}
