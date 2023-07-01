package main

import (
	"log"

	"github.com/nehalshaquib/GoShellCommander/api"
	"github.com/nehalshaquib/GoShellCommander/config"
)

func main() {
	err := config.Configure()
	if err != nil {
		log.Println("error in configuration: ", err)
	}

	logger := config.Logger

	logger.Infoln("shellCommander server staring...")
	server := api.NewServer()
	server.Run()
}
