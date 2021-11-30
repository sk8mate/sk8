package auth

import (
	"net/http"
	"sk8.town/backside/auth/domain"

	"github.com/gorilla/mux"
)

func Make(router *mux.Router) {
	usersDb := domain.NewUsersDb()
	tokenService := NewTokenService()
	googleTokenValidator := NewGoogleTokenValidator()
	googleAuthService := NewGoogleAuthService(usersDb, tokenService, googleTokenValidator)
	handler := GoogleHandler{googleAuthService}

	router.
		HandleFunc("/oauth/google/login", handler.Login).
		Methods(http.MethodPost).
		Name("LoginWithGoogleOAuth")
}