package main

import (
	"github.com/blupulov/xwsDislinkt/company-service/startup"
	"github.com/blupulov/xwsDislinkt/company-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
