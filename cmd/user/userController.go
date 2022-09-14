package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserHandlers(db *gorm.DB, routerGroup *gin.RouterGroup) {
	database := &Database{db: db}

	routerGroup.GET("/createUser", database.createUser)
}
