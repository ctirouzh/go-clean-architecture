package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/go-clean-architecture/config"
	"github.com/tahadostifam/go-clean-architecture/database"
)

func main() {
	server := gin.Default()

	cfg, configsErr := config.LoadConfigs()
	if configsErr != nil {
		log.Fatal("failed to load config file", configsErr)
	}

	db := database.GetSqliteInstance()
	defer db.Close()

	// studentRepository := repository.NewStudentRepository(db)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Server has listening on %s\n", addr)
	server.Run(addr)
}
