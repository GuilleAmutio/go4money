package user

import (
	"github.com/gin-gonic/gin"
	"github.com/guilleamutio/go4money/util"
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

// createUser godoc
// @Summary createUser
// @Schemes
// @Description create user in database
// @Param        user  body      User  true  "Create User"
// @Tags USERS
// @Accept json
// @Produce json
// @Success 200 {string} Finished
// @Router /api/v1/users/createUser [post]
func (userService *UserService) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	user := User{
		Username: req.Username,
		Password: req.Password,
	}

	err := userService.UserRepository.createUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
