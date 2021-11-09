package spots

import (
	"net/http"

	"sk8.town/backside/spots/dto"
	"sk8.town/backside/utils"
)

type FiltersHandler struct {
	service FiltersService
}

func (handler FiltersHandler) GetFilters(writer http.ResponseWriter, request *http.Request) {
	filters, appError := handler.service.GetAll()
	if appError != nil {
		utils.WriteError(writer, appError)
	} else {
		response := &dto.FilterGetAllResponse{
			Status: "success",
			Data:   filters,
		}
		utils.WriteJSON(writer, http.StatusOK, response)
	}
}
