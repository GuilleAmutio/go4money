package user

import (
	"github.com/gin-gonic/gin"
	"github.com/guilleamutio/go4money/models"
	"net/http"
)

func (database *Database) createUser(ctx *gin.Context) {
	database.db.Create(&models.User{Username: "Willy", Password: "mysecretpassword"})
	ctx.JSON(http.StatusOK, "Create user")
}
