package config

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nehalshaquib/GoShellCommander/logger"
	"go.uber.org/zap"
)

var (
	Logger           *zap.SugaredLogger
	AuthorizedTokens map[string]bool
	Host             string
	Port             string
	GinMode          string
)

func Configure() error {
	// Initializing logger
	log, err := logger.InitLogger()
	if err != nil {
		return err
	}
	Logger = log

	// Load env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Errorln("loading env: ", err)
		return err
	}

	host := os.Getenv("HOST")
	if host == "" {
		Host = "localhost"
		log.Warnln("HOST not provided, using default value: localhost")
	}
	Host = host

	port := os.Getenv("PORT")
	if port == "" {
		Port = "8085"
		log.Warnln("PORT not provided, using default value: 8085")
	}
	Port = port

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		GinMode = gin.DebugMode
		log.Warnln("GIN_MODE not provided, using default value: debug")
	}
	switch ginMode {
	case "release":
		GinMode = gin.ReleaseMode
	case "test":
		GinMode = gin.TestMode
	case "debug":
		GinMode = gin.DebugMode
	case "":
		GinMode = gin.DebugMode
		log.Warnln("GIN_MODE not provided, using default value: debug")
	default:
		log.Errorf("Invalid value: %s provided for GIN_MODE. Valid values are: %s, %s, %s", ginMode, gin.ReleaseMode, gin.TestMode, gin.DebugMode)
	}

	tokens := os.Getenv("TOKENS")
	if tokens == "" {
		return errors.New("TOKENS cannot be empty")
	}

	AuthorizedTokens = make(map[string]bool)
	tokenList := strings.Split(tokens, ",")
	for _, token := range tokenList {
		AuthorizedTokens[token] = true
	}

	return nil
}
