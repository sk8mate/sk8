package spots

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"sk8.town/backside/errs"

	"sk8.town/backside/spots/dto"
	"sk8.town/backside/utils"
)

type SpotsHandler struct {
	service SpotsService
}

func (handler SpotsHandler) AddSpot(writer http.ResponseWriter, request *http.Request) {
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
		utils.WriteProtoJSON(writer, http.StatusOK, response)
	}
}

func (handler SpotsHandler) GetSpot(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idAsString := vars["id"]
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		appError := errs.NewBadRequestError("")
		utils.WriteError(writer, appError)
		return
	}

	spot, appError := handler.service.Get(id)
	if appError != nil {
		utils.WriteError(writer, appError)
	} else {
		response := &dto.SpotsGetResponse{
			Status: "success",
			Data:   spot,
		}
		utils.WriteProtoJSON(writer, http.StatusOK, response)
	}
}

func (handler SpotsHandler) GetSpots(writer http.ResponseWriter, request *http.Request) {
	spots, appError := handler.service.GetAll()
	if appError != nil {
		utils.WriteError(writer, appError)
	} else {
		response := &dto.SpotsGetAllResponse{
			Status: "success",
			Data:   spots,
		}
		utils.WriteProtoJSON(writer, http.StatusOK, response)
	}
}

func (handler SpotsHandler) UpdateSpot(writer http.ResponseWriter, request *http.Request) {
	var spotsRequest dto.SpotsUpdateRequest
	if err := json.NewDecoder(request.Body).Decode(&spotsRequest); err != nil {
		utils.WriteError(writer, errs.NewBadRequestError(""))
		return
	}

	vars := mux.Vars(request)
	idAsString := vars["id"]
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		appError := errs.NewBadRequestError("")
		utils.WriteError(writer, appError)
		return
	}

	spot, appError := handler.service.Update(id, &spotsRequest)
	if appError != nil {
		utils.WriteError(writer, appError)
	} else {
		response := &dto.SpotsUpdateResponse{
			Status: "success",
			Data:   spot,
		}
		utils.WriteProtoJSON(writer, http.StatusOK, response)
	}
}

func (handler SpotsHandler) DeleteSpot(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idAsString := vars["id"]
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		appError := errs.NewBadRequestError("")
		utils.WriteError(writer, appError)
		return
	}

	appError := handler.service.Delete(id)
	if appError != nil {
		utils.WriteError(writer, appError)
	} else {
		utils.WriteStatus(writer, http.StatusOK)
	}
}
