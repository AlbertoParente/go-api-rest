package routes

import (
	"log"
	"net/http"

	"github.com/go-api-rest/controllers"
	"github.com/go-api-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetByAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetByPersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.NewPersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
