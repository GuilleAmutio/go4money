package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return UserService{
		UserRepository: userRepo,
	}
}

func (userService *UserService) createUser(ctx *gin.Context) {
	// HACER PUTA MIERDA DE LOGICA DE APLICACION
	ctx.JSON(http.StatusOK, "He sido invocado por mis cojones")
	// PEDIR AL DATA QUE REALIZA LA PUTA MIERDA DE TRANSACCION
	user := User{
		Username: "pabloMyGod",
		Password: "mysecret",
	}

	userService.UserRepository.createUser(user)

	ctx.JSON(http.StatusOK, "He terminado")
}
