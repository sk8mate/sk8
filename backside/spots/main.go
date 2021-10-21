package spots

import (
	"net/http"

	"github.com/gorilla/mux"

	"sk8.town/backside/config"
	"sk8.town/backside/spots/domain"
)

func Make(router *mux.Router) {
	cfg := config.Get()

	db:= domain.Connect(cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPassword)

	spotsDb := domain.NewSpotDb(db)
	spotsService := NewSpotService(spotsDb)
	spotsHandler := SpotHandler{spotsService}

	router.
		HandleFunc("/spots", spotsHandler.AddSpot).
		Methods(http.MethodPost).
		Name("AddNewSpot")
	router.
		HandleFunc("/spots/{id:[0-9]+}", spotsHandler.GetSpot).
		Methods(http.MethodGet).
		Name("GetSpotById")
	router.
		HandleFunc("/spots", spotsHandler.GetSpots).
		Methods(http.MethodGet).
		Name("GetAllSpots")
	router.
		HandleFunc("/spots/{id:[0-9]+}", spotsHandler.UpdateSpot).
		Methods(http.MethodPut).
		Name("UpdateSpot")
	router.
		HandleFunc("/spots/{id:[0-9]+}", spotsHandler.DeleteSpot).
		Methods(http.MethodDelete).
		Name("DeleteSpot")

	filtersDb := domain.NewFilterDb(db)
	filtersService := NewFilterService(filtersDb)
	filtersHandler := FilterHandler{filtersService}

	router.
		HandleFunc("/spots/filters", filtersHandler.GetFilters).
		Methods(http.MethodGet).
		Name("GetAllFilters")
}
