package routes

import (
	"log"
	"net/http"

	"github.com/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.http.HandleFunc("/", controllers.Home).Methods("Get")
	r.http.HandleFunc("/api/personalities/{id}", controllers.getByPersonality).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
