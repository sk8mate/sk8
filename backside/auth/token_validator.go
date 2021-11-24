package auth

import (
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/token_validator.go -package=mocks sk8.town/backside/auth TokenValidator
type TokenValidator interface {
	Validate(token string) (*UserClaims, *errs.AppError)
}
