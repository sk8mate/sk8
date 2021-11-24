package auth

import (
	"sk8.town/backside/errs"
	"sk8.town/backside/auth/dto"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/service.go -package=mocks sk8.town/backside/auth Service
type Service interface {
	Login(request *dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}

type DefaultService struct {
}

func (s DefaultService) Login(request *dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	return nil, nil
}

func NewService() DefaultService {
	return DefaultService{}
}
