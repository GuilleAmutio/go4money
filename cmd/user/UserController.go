package user

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService UserService
}

func NewUserController(userSvc UserService) UserController {
	return UserController{
		UserService: userSvc,
	}
}

func (userController UserController) RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/createUser", userController.UserService.createUser)
}
