package spots

import (
	"sk8.town/backside/errs"
	"sk8.town/backside/spots/domain"
	"sk8.town/backside/spots/dto"
	"strconv"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/filter_service.go -package=mocks sk8.town/backside/spots FilterService
type FilterService interface {
	GetAll() ([]*dto.FilterData, *errs.AppError)
}

type DefaultFilterService struct {
	filterRepo domain.FilterRepository
}

func (s DefaultFilterService) GetAll() ([]*dto.FilterData, *errs.AppError) {
	allFilterValues, err := s.filterRepo.GetAllFilterValues()
	if err != nil {
		return nil, err
	}

	mapping := make(map[domain.Filter][]*domain.FilterValue)
	for _, filterValue := range allFilterValues {
		mapping[filterValue.Filter] = append(mapping[filterValue.Filter], &domain.FilterValue{
			ID:    filterValue.ID,
			Value: filterValue.Value,
		})
	}

	var result []*dto.FilterData

	for filter, filterValues := range mapping {
		filterData := &dto.FilterData{
			Id:     strconv.Itoa(filter.ID),
			Name:   filter.Name,
			Type:   filter.Type,
			Values: nil,
		}

		for _, filterValue := range filterValues {
			filterData.Values = append(filterData.Values, &dto.FilterValueData{
				Id:   strconv.Itoa(filterValue.ID),
				Name: filterValue.Value,
			})
		}

		result = append(result, filterData)
	}

	return result, nil
}

func NewFilterService(repo domain.FilterRepository) DefaultFilterService {
	return DefaultFilterService{
		filterRepo: repo,
	}
}
