package spots

import (
	"github.com/gorilla/mux"
	"net/http"
	"sk8.town/backside/spots/domain"
)

func Make(router *mux.Router) {
	config := getConfig()

	spotsDb := domain.NewSpotDb(config.DbHost, config.DbPort, config.DbName, config.DbUser, config.DbPassword)
	spotsService := NewSpotService(spotsDb)
	handler := Handler{spotsService}

	router.
		HandleFunc("/spots", handler.AddSpot).
		Methods(http.MethodPost).
		Name("AddNewSpot")
}
