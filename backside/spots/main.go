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
}
