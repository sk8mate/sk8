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
	var spotsRequest dto.SpotsAddRequest
	if err := json.NewDecoder(request.Body).Decode(&spotsRequest); err != nil {
		utils.WriteError(writer, errs.NewBadRequestError(""))
		return
	}

	spot, appError := handler.service.Add(&spotsRequest)
	if appError != nil {
		utils.WriteError(writer, appError)
	} else {
		response := &dto.SpotsAddResponse{
			Status: "success",
			Data:   spot,
		}
		utils.WriteJSON(writer, http.StatusOK, response)
	}
}
