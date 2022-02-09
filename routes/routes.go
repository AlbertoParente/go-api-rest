package routes

import (
	"log"
	"net/http"

	"github.com/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetByAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetByPersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.NewPersonality).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
