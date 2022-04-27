package main

import (
	"fmt"

	"github.com/Johnman67112/go_api/database"
	"github.com/Johnman67112/go_api/models"
	"github.com/Johnman67112/go_api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", Story: "Story 1"},
		{Id: 2, Name: "Name 2", Story: "Story 2"},
	}

	database.DatabaseConnect()

	fmt.Println("Start Rest server with Go")
	routes.HandleRequest()
}
