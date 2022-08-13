package main

import (
	"fmt"
	"log"

	"lms/config"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	cfg, configsErr := config.LoadConfigs()
	if configsErr != nil {
		log.Fatal("failed to load config file", configsErr)
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Server has listening on %s\n", addr)
	server.Run(addr)
}
