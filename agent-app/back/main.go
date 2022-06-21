package main

import (
	"github.com/blupulov/xwsDislinkt/agent-app/back/startup"
	"github.com/blupulov/xwsDislinkt/agent-app/back/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
