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

func GetByAllPersonalities(w, http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetByPersonality(w, http.ResponseWriter, r *http.Request) {
	vars := mux.vars(r)
	id := vars["id"]
	var personality models.Personalities
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(personality)
}

func NewPersonality(w, http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(personality)
	database.DB.Create(&personality)
	json.NewEncoder(w).Encode((personality))
}
