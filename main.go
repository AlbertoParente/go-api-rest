package main

import (
	"github.com/go-api-rest/models"
	"github.com/go-api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Nme 1", History: "History 1"},
		{Id: 2, Name: "Nme 2", History: "History 2"},
	}
	routes.HandleRequest()
}
