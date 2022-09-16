package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/guilleamutio/go4money/cmd/user"
	_ "github.com/guilleamutio/go4money/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{db: db}
	router := gin.Default()

	userCtrl := user.InitializeDependencies(db)

	registerHandlers(userCtrl, router)

	server.router = router

	return server
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func registerHandlers(userCtrl user.UserController, router *gin.Engine) {
	routerGroup := router.Group("/api/v1")
	{
		accRG := routerGroup.Group("/users")
		{
			userCtrl.RegisterUserRoutes(accRG)
		}

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
