package api

import (
	"fmt"

	db "github.com/codernirmalnp/golang/db/sqlc"
	"github.com/codernirmalnp/golang/token"
	"github.com/codernirmalnp/golang/util"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewMarker(config.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot Create token maker:%w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	router.GET("/", server.Test)
	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.LoginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.CreateAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/transfers", server.CreateTransfer)

	server.router = router
	return server, nil

}
func (server *Server) Start(address string) error {
	return server.router.Run(address)

}
