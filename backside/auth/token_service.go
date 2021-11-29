package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"sk8.town/backside/auth/config"
	"sk8.town/backside/errs"
	"time"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/token_service.go -package=mocks sk8.town/backside/auth TokenService
type TokenService interface {
	CreateToken(email string, oauthId string) (string, *errs.AppError)
	ParseToken(token string) (*UserClaims, *errs.AppError)
}

type DefaultTokenService struct{}

func (DefaultTokenService) CreateToken(email string, oauthId string) (string, *errs.AppError) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
		Email: email,
		OAuthId: oauthId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signedToken, err := token.SignedString([]byte(config.GetConfig().SecretKey))
	if err != nil {
		return "", errs.NewUnexpectedError("could not sign token")
	}
	return signedToken, nil
}

func (DefaultTokenService) ParseToken(token string) (*UserClaims, *errs.AppError) {
	tokenAfterVerification, err := jwt.ParseWithClaims(token, &UserClaims{}, func(tokenBeforeVerification *jwt.Token) (interface{}, error) {
		if tokenBeforeVerification.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("signing method not matched")
		}

		return []byte(config.GetConfig().SecretKey), nil
	})

	if err != nil || !tokenAfterVerification.Valid {
		return nil, errs.NewUnexpectedError("token is not valid")
	}

	userClaims := tokenAfterVerification.Claims.(*UserClaims)

	return userClaims, nil
}
