package controllers

import (
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	logrus.Info("Getting all users")
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
func (c *UserController) CreateUser(ctx *gin.Context) {
	logrus.Println("CreateUser")
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	id, err := c.service.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	} else {
		ctx.JSON(http.StatusOK, id)
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	logrus.Println("Inside login func")
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	token, err := c.service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}

}
