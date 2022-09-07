package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/guilleamutio/go4money/db/sqlc"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts/createAccount", server.createAccount)
	router.PUT("/accounts/updateAccount", server.updateAccount)
	router.DELETE("/accounts/deleteAccount/:id", server.deleteAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/listAccounts", server.listAccounts)
	router.GET("/accounts/listAllAccounts", server.listAllAccounts)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
