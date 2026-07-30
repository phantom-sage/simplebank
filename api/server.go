package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/phantom-sage/simplebank/db/sqlc"
)

// Server serve HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer create new server instance.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// HTTPs routes
	router.POST("/api/accounts", server.createAccount)
	router.GET("/api/accounts/:id", server.getAccount)
	router.GET("/api/accounts", server.listAccounts)

	server.router = router

	return server
}

// Start the HTTP server and listen on address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
