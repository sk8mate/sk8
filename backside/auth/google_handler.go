package auth

import (
	"encoding/json"
	"net/http"
	"sk8.town/backside/auth/dto"
	"sk8.town/backside/errs"
	"sk8.town/backside/utils"
)

type GoogleHandler struct {
	authService AuthService
}

func (handler GoogleHandler) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequestBody dto.LoginRequestBody
	if err := json.NewDecoder(request.Body).Decode(&loginRequestBody); err != nil {
		utils.WriteError(writer, errs.NewBadRequestError(""))
		return
	}

	loggedInData, appError := handler.authService.Login(&LoginData{OAuthIdToken: loginRequestBody.IdToken})
	if appError != nil {
		utils.WriteError(writer, appError)
	} else {
		response := &dto.LoginResponseBody{
			Status: "success",
			Data:   &dto.LoginResponseData{Token: loggedInData.Token},
		}
		utils.WriteProtoJSON(writer, http.StatusOK, response)
	}
}
