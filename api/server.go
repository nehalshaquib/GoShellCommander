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

var log *zap.SugaredLogger

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	log = config.Logger
	server := &Server{}

	gin.SetMode(config.GinMode)
	router := gin.Default()
	router.Use(GinZapMiddleware(log))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "GoShellCommander is running.")
	})

	api := router.Group("api")
	api.Use(authMiddleWare)

	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "This is an authorized api group")
	})
	api.POST("/cmd", server.cmdHandler)

	server.router = router
	return server
}

func (server *Server) Run() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		addr := "0.0.0.0:" + config.Port
		err := server.router.Run(addr)
		if err != nil {
			log.Errorln("error in starting server: ", err)
			signalChannel <- os.Interrupt
		}
		log.Infoln("goShellCommander running on: ", addr)
	}()
	log.Infoln("shellCommander server started")

	<-signalChannel
	close(signalChannel)
	log.Infoln("shellCommander server stopping...")
	os.Exit(0)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

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
