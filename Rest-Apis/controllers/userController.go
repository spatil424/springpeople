package controllers

import (
	"net/http"
	"restapi/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (u *UserController) GetAllUsers(ctx *gin.Context) {
	users := u.service.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}
