package dependency

import (
	"restapi/controllers"
	"restapi/repository"
	"restapi/services"
)

type Container struct {
	UserController *controllers.UserController
}

func BuildContainer() *Container {
	userRepo := repository.NewRepostiory()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
