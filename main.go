package main

import (
	"log"

	"github.com/nehalshaquib/GoShellCommander/api"
	"github.com/nehalshaquib/GoShellCommander/config"
)

func main() {
	err := config.Configure()
	if err != nil {
		log.Println("error in configuration ", err)
	}

	logger := config.Logger

	logger.Infoln("starting shellCommander server...")
	server := api.NewServer()
	err = server.Run()
	if err != nil {
		logger.Errorln("error in starting server: ", err)
	}
}
