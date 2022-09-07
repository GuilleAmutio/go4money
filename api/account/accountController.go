package account

import (
	"github.com/gin-gonic/gin"
	db "github.com/guilleamutio/go4money/db/sqlc"
)

func RegisterAccountHandlers(store *db.Store, routerGroup *gin.RouterGroup) {
	server := &Server{store: store}

	routerGroup.POST("/createAccount", server.createAccount)
	routerGroup.PUT("/updateAccount", server.updateAccount)
	routerGroup.DELETE("/deleteAccount/:id", server.deleteAccount)
	routerGroup.GET("/:id", server.getAccount)
	routerGroup.GET("/listAccounts", server.listAccounts)
	routerGroup.GET("/listAllAccounts", server.listAllAccounts)
}
