package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/guilleamutio/go4money/cmd/user"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{db: db}
	router := gin.Default()

	registerHandlers(db, router)

	server.router = router

	return server
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func registerHandlers(db *gorm.DB, router *gin.Engine) {
	routerGroup := router.Group("/users")
	{
		user.RegisterUserHandlers(db, routerGroup)
	}
}
