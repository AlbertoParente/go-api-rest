package routes

import (
	"log"
	"net/http"

	"github.com/go-api-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.getByAllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
