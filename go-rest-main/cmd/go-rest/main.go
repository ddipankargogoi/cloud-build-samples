package main

import (
	"log"

	"github.com/ddipankargogoi/go-rest/internal/animals"
	"github.com/ddipankargogoi/go-rest/internal/api"
	"github.com/ddipankargogoi/go-rest/internal/cars"
	"github.com/ddipankargogoi/go-rest/internal/database"
)

func main() {
	log.Println("I'm go-rest, nice to meet you!")

	db := database.NewDatabase()
	defer db.Disconnect()

	api := api.NewAPI()

	animals := animals.NewAnimalsResource(db)
	api.AddResource(animals)

	cars := cars.CarsResource{}
	api.AddResource(&cars)

	api.Start()
}
