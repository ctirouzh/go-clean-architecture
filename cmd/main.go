package main

import (
	"fmt"
	"log"

	"lms/config"
	"lms/internal/adapters/controller"
	"lms/internal/adapters/repository/memory"
	"lms/internal/app/auth"
	"lms/internal/ports/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, cfgErr := config.Parse()
	if cfgErr != nil {
		log.Fatal("failed to load config file", cfgErr)
	}

	userRepo := memory.NewUserRepo()
	authService := auth.NewService(userRepo)
	authCtrl := controller.NewAuthController(authService)

	r := gin.Default()
	http.InitAuthRouter(r, authCtrl)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("server listening on %s\n", addr)
	r.Run(addr)
}
