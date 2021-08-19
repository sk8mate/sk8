package places

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Make(router *mux.Router) {
	repository := NewRepository()
	handler := Handler{NewService(repository)}

	router.
		HandleFunc("/places/autocomplete", handler.GetPlacesAutocomplete).
		Methods(http.MethodGet).
		Name("GetPlacesAutocomplete")
}