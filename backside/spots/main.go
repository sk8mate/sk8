package places

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type UserModel struct{
	Id int `gorm:"primary_key"`
	Name string
	Address string
}

func main(){

}

func Make(router *mux.Router) {
	//config := getConfig()

	//locationService := NewLocationService(config.TomtomApiKey)
	//handler := Handler{NewService(locationService)}

	//router.
	//	HandleFunc("/spots", handler.GetPlacesAutocomplete).
	//	Methods(http.MethodPost).
	//	Name("AddNewSpot")
}
