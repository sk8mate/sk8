package auth

import (
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/token_service.go -package=mocks sk8.town/backside/auth TokenService
type TokenService interface {
	Verify(token string) (*UserClaims, *errs.AppError)
}
