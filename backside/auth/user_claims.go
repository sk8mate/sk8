package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	OAuthId string `json:"oauth_id,omitempty"`
	Email   string `json:"email"`
}
