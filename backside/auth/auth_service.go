package auth

import (
	"sk8.town/backside/errs"
)

type LoginData struct {
	OAuthIdToken string
}

type LoggedInData struct {
	Token string
}

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/auth_service.go -package=mocks sk8.town/backside/auth AuthService
type AuthService interface {
	Login(*LoginData) (*LoggedInData, *errs.AppError)
}