package main

import (
	"fmt"
	"log"

	"lms/config"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	cfg, cfgErr := config.Parse()
	if cfgErr != nil {
		log.Fatal("failed to load config file", cfgErr)
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("server listening on %s\n", addr)
	server.Run(addr)
}
