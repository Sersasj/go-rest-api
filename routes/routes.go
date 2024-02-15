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
	r.HandleFunc("/api/personalidades", controllers.GetPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadesById).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))
}
