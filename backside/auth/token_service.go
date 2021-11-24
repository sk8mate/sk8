package auth

import (
	"context"
	"google.golang.org/api/idtoken"
	"sk8.town/backside/auth/config"
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/token_service.go -package=mocks sk8.town/backside/auth TokenService
type TokenService interface {
	Verify(token string) (*UserClaims, *errs.AppError)
}

type DefaultTokenService struct {
}

func (s DefaultTokenService) Verify(token string) (*UserClaims, *errs.AppError) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, token, config.GetConfig().GoogleClientId)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	googleClaims:= UserClaims{
		Id:            payload.Subject,
		Email:         payload.Claims["email"].(string),
		EmailVerified: payload.Claims["email_verified"].(bool),
		FirstName:     payload.Claims["given_name"].(string),
		LastName:      payload.Claims["family_name"].(string),
	}

	return &googleClaims, nil
}

func NewTokenService() DefaultTokenService {
	return DefaultTokenService{}
}
