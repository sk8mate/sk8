package auth

import (
	"context"
	"google.golang.org/api/idtoken"
	"sk8.town/backside/auth/config"
	"sk8.town/backside/errs"
	"time"
)

type GoogleTokenService struct {
}

func (s GoogleTokenService) verifyPayload(payload *idtoken.Payload) (*UserClaims, *errs.AppError) {
	if payload.Issuer != "accounts.google.com" && payload.Issuer != "https://accounts.google.com" {
		return nil, errs.NewUnexpectedError("iss is invalid")
	}

	if payload.Expires < time.Now().UTC().Unix() {
		return nil, errs.NewUnexpectedError("token is expired")
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return nil, errs.NewUnexpectedError("email field in payload is invalid")
	}

	emailVerified, ok := payload.Claims["email_verified"].(bool)
	if !ok {
		return nil, errs.NewUnexpectedError("email_verified field in payload is invalid")
	}

	firstName, ok := payload.Claims["given_name"].(string)
	if !ok {
		return nil, errs.NewUnexpectedError("given_name field in payload is invalid")
	}

	lastName, ok := payload.Claims["family_name"].(string)
	if !ok {
		return nil, errs.NewUnexpectedError("family_name field in payload is invalid")
	}

	userClaims := UserClaims{
		Id:            payload.Subject,
		Email:         email,
		EmailVerified: emailVerified,
		FirstName:     firstName,
		LastName:      lastName,
	}

	return &userClaims, nil
}

func (s GoogleTokenService) Verify(token string) (*UserClaims, *errs.AppError) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, token, config.GetConfig().GoogleClientId)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	googleUserClaims, appError := s.verifyPayload(payload)
	if appError != nil{
		return nil, appError
	}

	return googleUserClaims, nil
}

func NewGoogleTokenService() GoogleTokenService {
	return GoogleTokenService{}
}
