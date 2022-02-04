package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func getByAllPersonalities(w, http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
