package main

import (
	"log"
	"restapi/config"
	"restapi/dependency"
	"restapi/routes"
)

func main() {

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	container := dependency.BuildContainer()
	r := routes.SetUpRouter(container.UserController)
	r.Run(":8081")

}
