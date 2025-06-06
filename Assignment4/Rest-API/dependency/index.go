package dependency

import (
	"rest-api/controllers"
	"rest-api/repository"
	"rest-api/services"

	"gorm.io/gorm"
)

type Container struct {
	UserController  *controllers.UserController
	OrderController *controllers.OrderController
}

func BuildContainer(db *gorm.DB) *Container {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	return &Container{
		UserController:  userController,
		OrderController: orderController,
	}
}
