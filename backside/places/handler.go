package places

import (
	"net/http"

	"github.com/gorilla/schema"

	"sk8.town/backside/errs"
	"sk8.town/backside/places/dto"
	"sk8.town/backside/utils"
)

var decoder = schema.NewDecoder()

type Handler struct {
	Service Service
}

func (handler Handler) GetPlacesAutocomplete(writer http.ResponseWriter, request *http.Request) {
	var placesRequest dto.AutocompleteRequest

	if err := decoder.Decode(&placesRequest, request.URL.Query()); err != nil {
		utils.WriteError(writer, errs.NewBadRequestError(""))
		return
	}

	places, appError := handler.Service.GetPlaces(&placesRequest)
	if appError != nil {
		utils.WriteError(writer, appError)
	} else {
		response := &dto.AutocompleteResponse{
			Status: "success",
			Data:   places,
		}
		utils.WriteJSON(writer, http.StatusOK, response)
	}
}
