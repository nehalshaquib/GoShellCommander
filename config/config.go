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

// Global variables that will be used throughout the application
var (
	Logger           *zap.SugaredLogger // Logger instance for logging
	AuthorizedTokens map[string]bool    // Map to hold authorized tokens
	Port             string             // Port on which the server will run
	GinMode          string             // Mode in which Gin framework will run
)

// Configure function initializes the logger, loads environment variables from .env file,
// sets the port and Gin mode, and populates the AuthorizedTokens map.
func Configure() error {
	// Initializing logger
	log, err := logger.InitLogger()
	if err != nil {
		return err
	}
	Logger = log

	// Load environment variables from .env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Warnln("loading env: ", err)
	}

	// Get the port from environment variables, if not provided, use 8085 as default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
		log.Warnln("PORT not provided, using default value: 8085")
	}
	Port = port

	// Get the Gin mode from environment variables, if not provided, use debug as default
	ginMode := os.Getenv("GIN_MODE")
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

	// Get the authorized tokens from environment variables, if not provided, return an error
	tokens := os.Getenv("TOKENS")
	if tokens == "" {
		return errors.New("TOKENS cannot be empty")
	}

	// Populate the AuthorizedTokens map
	AuthorizedTokens = make(map[string]bool)
	tokenList := strings.Split(tokens, ",")
	for _, token := range tokenList {
		AuthorizedTokens[token] = true
	}

	return nil
}
