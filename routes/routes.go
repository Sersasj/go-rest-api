package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.GetPersonalidades).Methods("GET")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersonalidadesById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
