package main

import (
	"fmt"
	"log"

	"lms/config"
	"lms/internal/adapters/controller"
	"lms/internal/adapters/repository/memory"
	"lms/internal/app/auth"

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
	auth := r.Group("/api/auth")
	auth.POST("/signup", authCtrl.SignUp)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("server listening on %s\n", addr)
	r.Run(addr)
}
