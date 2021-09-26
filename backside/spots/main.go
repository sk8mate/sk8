package spots

import (
	"net/http"

	"github.com/gorilla/mux"

	"sk8.town/backside/db"
	"sk8.town/backside/spots/domain"
)

func Make(router *mux.Router, config db.Config) {
	spotsDb := domain.NewSpotDb(config.DbHost, config.DbPort, config.DbName, config.DbUser, config.DbPassword)
	spotsService := NewSpotService(spotsDb)
	handler := Handler{spotsService}

	router.
		HandleFunc("/spots", handler.AddSpot).
		Methods(http.MethodPost).
		Name("AddNewSpot")
}
