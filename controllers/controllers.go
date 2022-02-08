package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-api-rest/database"
	"github.com/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func getByAllPersonalities(w, http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func getByPersonality(w, http.ResponseWriter, r *http.Request) {
	vars := mux.vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(w).Encode(personality)
		}
	}
}
