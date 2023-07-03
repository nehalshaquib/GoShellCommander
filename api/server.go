package api

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
		err := server.router.Run(config.Host + ":" + config.Port)
		if err != nil {
			log.Errorln("error in starting server: ", err)
			signalChannel <- os.Interrupt
		}
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

func authMiddleWare(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("authorization key missing")))
		ctx.Abort()
		return
	}
	if isTokenValid(token) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid authorization key")))
		ctx.Abort()
		return
	}

	ctx.Next()
}

func isTokenValid(token string) bool {
	//TODO: Implement JWT token
	_, ok := config.AuthorizedTokens[token]
	return ok
}
