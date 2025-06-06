package main

import (
	"rest-api/config"
	"rest-api/dependency"
	"rest-api/routers"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Application starting...")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	logrus.Debug("Starting the application...")

	db := config.ConnectDB()
	if db == nil {
		logrus.Fatalf("Failed to connect to database. Application cannot start.")
	}

	container := dependency.BuildContainer(db)
	r := routers.SetUpRoutes(container.UserController, container.OrderController)

	r.Run(":8081")
}
