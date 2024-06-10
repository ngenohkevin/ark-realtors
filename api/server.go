package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngenohkevin/ark-realtors/internal/handlers"
	"github.com/ngenohkevin/ark-realtors/internal/store"
	"github.com/ngenohkevin/ark-realtors/internal/token"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
)

type Server struct {
	Config     utils.Config
	Store      store.Store
	TokenMaker token.Maker
	Router     *gin.Engine
}

func NewServer(config utils.Config, store store.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
		Router:     gin.Default(),
	}
	server.SetUpRouter()
	return server, nil
}

func (server *Server) SetUpRouter() {
	router := server.Router

	router.GET("/albums", handlers.GetAlbums)

	//paths
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(AuthMiddleware(server.TokenMaker))

	authRoutes.GET("/users/:username", server.getUser)
	authRoutes.PUT("/users/:id", server.updateUser)
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
