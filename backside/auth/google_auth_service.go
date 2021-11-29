package auth

import (
	"sk8.town/backside/auth/domain"
	"sk8.town/backside/errs"
)

type GoogleAuthService struct {
	tokenValidator  GoogleTokenValidator
	usersRepository domain.UsersRepository
	tokenService    TokenService
}

func (s GoogleAuthService) Login(loginData *LoginData) (*LoggedInData, *errs.AppError) {
	oauthClientClaims, err := s.tokenValidator.Verify(loginData.OAuthIdToken)
	if err != nil {
		return nil, err
	}

	_, err = s.usersRepository.Get(oauthClientClaims.Email)
	if err != nil {
		newUser := domain.User{
			FirstName: oauthClientClaims.FirstName,
			LastName:  oauthClientClaims.LastName,
			Email:     oauthClientClaims.Email,
			OAuthId:   oauthClientClaims.Id,
		}

		if err = s.usersRepository.Add(&newUser); err != nil {
			return nil, errs.NewUnexpectedError(err.Message)
		}
	}

	token, err := s.tokenService.CreateToken(oauthClientClaims.Email, oauthClientClaims.Id)
	if err != nil {
		return nil, err
	}

	loggedInData := LoggedInData{Token: token}

	return &loggedInData, nil
}

func NewGoogleAuthService(usersRepositoryInit domain.UsersRepository, tokenServiceInit TokenService) GoogleAuthService {
	return GoogleAuthService{usersRepository: usersRepositoryInit, tokenService: tokenServiceInit}
}
