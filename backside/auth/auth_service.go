package auth

import (
	"sk8.town/backside/auth/interfaces"
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/auth_service.go -package=mocks sk8.town/backside/auth AuthService
type AuthService interface {
	Login(*interfaces.LoginData) (*interfaces.LoggedInData, *errs.AppError)
}

type DefaultAuthService struct {
	tokenService TokenService
}

func (s DefaultAuthService) Login(loginData *interfaces.LoginData) (*interfaces.LoggedInData, *errs.AppError) {


	//email, ok := OauthConnections[googleId]
	//if !ok {
	//	signedToken, err := domain.CreateToken(googleId)
	//	if err != nil {
	//		http.Error(c.Writer, "could not create token in oauth google receive", http.StatusInternalServerError)
	//		return
	//	}
	//
	//	urlValues := url.Values{}
	//	urlValues.Add("signedToken", signedToken)
	//	urlValues.Add("name", googleResponse.Name)
	//	urlValues.Add("email", googleResponse.Email)
	//	http.Redirect(c.Writer, c.Request, "/partial-register?"+urlValues.Encode(), http.StatusSeeOther)
	//	return
	//}
	//
	//err = CreateSession(email, c.Writer)
	//if err != nil {
	//	http.Error(c.Writer, "can not create user session", http.StatusInternalServerError)
	//}
	//
	//msg := url.QueryEscape("logged in " + email)
	//http.Redirect(c.Writer, c.Request, "/?msg="+msg, http.StatusSeeOther)

	return nil, nil
}

func NewAuthService(tokenServiceInit TokenService) DefaultAuthService {
	return DefaultAuthService{tokenService: tokenServiceInit}
}
