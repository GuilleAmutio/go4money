package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/guilleamutio/go4money/db/sqlc"
	"github.com/guilleamutio/go4money/token"
	"github.com/guilleamutio/go4money/util"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			return nil, fmt.Errorf("cannot create validator: %w", err)
		}
	}

	server.registerRouter()

	return server, nil
}

func (server *Server) registerRouter() {
	router := gin.Default()

	router.POST("accounts/login", server.loginUser)
	router.POST("/users/createUser", server.createUser)
	router.GET("/accounts/listAccounts", server.listAccounts)
	router.GET("/accounts/listAllAccounts", server.listAllAccounts)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts/createAccount", server.createAccount)
	authRoutes.PUT("/accounts/updateAccount", server.updateAccount)
	authRoutes.DELETE("/accounts/deleteAccount/:id", server.deleteAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)

	authRoutes.POST("/transfers/createTransfer", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
