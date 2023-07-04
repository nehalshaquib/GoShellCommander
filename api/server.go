package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nehalshaquib/GoShellCommander/config"
	"go.uber.org/zap"
)

// log is a global logger instance used throughout the api package.
var log *zap.SugaredLogger

// Server struct represents the server with its router.
type Server struct {
	router *gin.Engine
}

// NewServer initializes a new server with its routes and middleware.
func NewServer() *Server {
	// Initialize the logger
	log = config.Logger
	server := &Server{}

	// Set the gin mode and initialize the router
	gin.SetMode(config.GinMode)
	router := gin.Default()
	router.Use(GinZapMiddleware(log))

	// Define the root route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "GoShellCommander is running.")
	})

	// Define the API routes
	api := router.Group("api")
	api.Use(authMiddleWare)

	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "This is an authorized api group")
	})
	api.POST("/cmd", server.cmdHandler)

	server.router = router
	return server
}

// Run starts the server and listens for incoming requests.
func (server *Server) Run() {
	// Create a channel to listen for OS signals
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Start the server
		addr := "0.0.0.0:" + config.Port
		err := server.router.Run(addr)
		if err != nil {
			log.Errorln("error in starting server: ", err)
			signalChannel <- os.Interrupt
		}
		log.Infoln("goShellCommander running on: ", addr)
	}()
	log.Infoln("shellCommander server started")

	// Wait for an OS signal to stop the server
	<-signalChannel
	close(signalChannel)
	log.Infoln("shellCommander server stopping...")
	os.Exit(0)
}

// errorResponse creates a JSON response with an error message.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// GinZapMiddleware is middleware function that logs incoming requests.
func GinZapMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		// Process request
		ctx.Next()

		// Log request details
		logger.Infof("Incoming request: {path: %v}, {method: %v}, {ip: %v}, {user-agent: %v}, {status: %v}, {latency: %v}",
			ctx.Request.URL.Path,
			ctx.Request.Method,
			ctx.ClientIP(),
			ctx.Request.UserAgent(),
			fmt.Sprintf("%d", ctx.Writer.Status()),
			time.Since(start),
		)
	}
}
