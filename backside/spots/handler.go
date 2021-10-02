package spots

import (
	"encoding/json"
	"net/http"
	"sk8.town/backside/errs"

	"sk8.town/backside/spots/dto"
	"sk8.town/backside/utils"
)

type Handler struct {
	service SpotService
}

func (handler Handler) AddSpot(writer http.ResponseWriter, request *http.Request) {
	var spotRequest dto.SpotRequest
	if err := json.NewDecoder(request.Body).Decode(&spotRequest); err != nil {
		appError := errs.NewBadRequestError("")
		utils.WriteResponse(writer, appError.Code, appError.AsMessage())
		return
	}

	response, appError := handler.service.Add(&spotRequest)
	if appError != nil {
		utils.WriteResponse(writer, appError.Code, appError.AsMessage())
	} else {
		utils.WriteResponse(writer, http.StatusOK, response)
	}
}
