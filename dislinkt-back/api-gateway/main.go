package main

import (
	"log"
	"os"

	"github.com/blupulov/xwsDislinkt/api-gateway/startup"
	"github.com/blupulov/xwsDislinkt/api-gateway/startup/config"
)

func main() {
	logger := log.New(os.Stdout, "API-GW => ", log.LstdFlags|log.Lshortfile)
	config := config.NewConfig()
	server := startup.NewServer(config, logger)
	server.Start()
}
