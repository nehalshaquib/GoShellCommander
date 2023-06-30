package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()
	api := router.Group("api")
	api.POST("/cmd", server.cmdHandler)

	server.router = router
	return server
}

func (server *Server) Run() error { return server.router.Run("localhost:8085") }

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
