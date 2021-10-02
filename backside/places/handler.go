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
		appError := errs.NewBadRequestError("")
		utils.WriteResponse(writer, appError.Code, appError.AsMessage())
		return
	}

	response, appError := handler.Service.GetPlaces(&placesRequest)
	if appError != nil {
		utils.WriteResponse(writer, appError.Code, appError.AsMessage())
	} else {
		utils.WriteResponse(writer, http.StatusOK, response)
	}
}
