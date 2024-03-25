package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngenohkevin/ark-realtors/internal/handlers"
	"github.com/ngenohkevin/ark-realtors/internal/store"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
)

type Server struct {
	Config utils.Config
	Store  *store.Store
	Router *gin.Engine
}

func NewServer(config utils.Config, store *store.Store) (*Server, error) {
	server := &Server{
		Config: config,
		Store:  store,
		Router: gin.Default(),
	}
	server.SetUpRouter()
	return server, nil
}

func (server *Server) SetUpRouter() {
	router := server.Router

	router.GET("/albums", handlers.GetAlbums)

	router.POST("/users", server.createUser)
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
