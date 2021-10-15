package spots

import (
	"net/http"

	"github.com/gorilla/mux"

	"sk8.town/backside/config"
	"sk8.town/backside/spots/domain"
)

func Make(router *mux.Router) {
	cfg := config.Get()

	spotsDb := domain.NewSpotDb(cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPassword)
	spotsService := NewSpotService(spotsDb)
	handler := Handler{spotsService}

	router.
		HandleFunc("/spots", handler.AddSpot).
		Methods(http.MethodPost).
		Name("AddNewSpot")
	router.
		HandleFunc("/spots/{id:[0-9]+}", handler.GetSpot).
		Methods(http.MethodGet).
		Name("GetSpotById")
	router.
		HandleFunc("/spots", handler.GetSpots).
		Methods(http.MethodGet).
		Name("GetAllSpots")
	router.
		HandleFunc("/spots/{id:[0-9]+}", handler.UpdateSpot).
		Methods(http.MethodPut).
		Name("UpdateSpot")
	router.
		HandleFunc("/spots/{id:[0-9]+}", handler.DeleteSpot).
		Methods(http.MethodDelete).
		Name("DeleteSpot")
}
