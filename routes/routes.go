package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadesById).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
