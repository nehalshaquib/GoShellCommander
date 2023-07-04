package main

import (
	"log"

	"github.com/nehalshaquib/GoShellCommander/api"
	"github.com/nehalshaquib/GoShellCommander/config"
)

func main() {
	// Call the Configure function from the config package to initialize the application configuration.
	// This includes setting up the logger, loading environment variables, setting the server port and Gin mode,
	// and populating the AuthorizedTokens map.
	err := config.Configure()
	if err != nil {
		log.Println("error in configuration: ", err)
	}

	// Get the logger instance from the config package.
	logger := config.Logger

	logger.Infoln("shellCommander server staring...")

	// Create a new server instance using the NewServer function from the api package.
	server := api.NewServer()

	// Start the server.
	server.Run()
}
