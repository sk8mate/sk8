package auth

import (
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/token_service.go -package=mocks sk8.town/backside/auth TokenService
type TokenService interface {
	CreateToken(email string) (string, *errs.AppError)
	ParseToken(token string) (string, error)
}

//func CreateToken(email string) (string, error) {
//	claims := MyClaims{
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
//		},
//		UserId: userId,
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
//	signedToken, err := token.SignedString([]byte(config.Get().SecretKey))
//	if err != nil {
//		return "", fmt.Errorf("could not sign token %v", err)
//	}
//	return signedToken, nil
//}
//
//func ParseToken(token string) (string, error) {
//	tokenAfterVerification, err := jwt.ParseWithClaims(token, &MyClaims{}, func(tokenBeforeVerification *jwt.Token) (interface{}, error) {
//		if tokenBeforeVerification.Method.Alg() != jwt.SigningMethodHS256.Alg() {
//			return nil, fmt.Errorf("signing method not matched")
//		}
//
//		return []byte(config.Get().SecretKey), nil
//	})
//
//	if err != nil || !tokenAfterVerification.Valid {
//		return "", fmt.Errorf("token is not valid")
//	}
//
//	claims := tokenAfterVerification.Claims.(*MyClaims)
//	return claims.UserId, nil
//}
