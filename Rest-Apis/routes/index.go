package routes

import (
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

// "net/http"

// "github.com/gin-gonic/gin"

// func CreateRoutes(r *gin.Engine) {
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	r.POST("/welcome", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "welcome from POST",
// 		})
// 	})
// }

func SetUpRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()
	r.GET("/users", userController.GetAllUsers)

	return r
}
