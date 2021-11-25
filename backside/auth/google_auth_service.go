package auth

import (
	"fmt"
	"sk8.town/backside/errs"
)

type GoogleAuthService struct {
	tokenService TokenService
}

func (s GoogleAuthService) Login(loginData *LoginData) (*LoggedInData, *errs.AppError) {
	userClaims, err := s.tokenService.Verify(loginData.OAuthIdToken)
	if err != nil{
		return nil, err
	}

	fmt.Println(userClaims)
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

func NewGoogleAuthService(tokenServiceInit TokenService) GoogleAuthService {
	return GoogleAuthService{tokenService: tokenServiceInit}
}

