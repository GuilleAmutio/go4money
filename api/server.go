package api

import (
	"github.com/gin-gonic/gin"
	"github.com/guilleamutio/go4money/api/account"
	db "github.com/guilleamutio/go4money/db/sqlc"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	registerHandlers(store, router)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func registerHandlers(store *db.Store, router *gin.Engine) {

	routerGroup := router.Group("/accounts")
	{
		account.RegisterAccountHandlers(store, routerGroup)
	}

}
