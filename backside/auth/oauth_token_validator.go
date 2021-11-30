package auth

import (
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/oauth_token_validator.go -package=mocks sk8.town/backside/auth OAuthTokenValidator
type OAuthTokenValidator interface {
	Verify(token string) (*OAuthClientClaims, *errs.AppError)
}
