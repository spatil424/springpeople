package routers

import (
	"rest-api/controllers"
	"rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(userController *controllers.UserController, orderController *controllers.OrderController) *gin.Engine {
	r := gin.Default()

	public := r.Group("/")
	{
		public.POST("/login", userController.Login)
		public.POST("/createuser", userController.CreateUser)
	}

	authed := r.Group("/")
	{
		authed.GET("/users", middleware.AuthMiddleware("admin"), userController.GetAllUsers)
		authed.GET("/orders", middleware.AuthMiddleware(), orderController.GetOrders)
		authed.PUT("/orders/:id", middleware.AuthMiddleware("admin"), orderController.UpdateOrder)
	}
	return r
}
