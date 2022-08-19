package main

import (
	"fmt"
	"log"

	"lms/config"
	"lms/database"
	"lms/internal/adapters/controller"
	"lms/internal/adapters/repository/postgres"
	"lms/internal/app/auth"
	"lms/internal/ports/http"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, cfgErr := config.Parse("./config/config.json")
	if cfgErr != nil {
		log.Fatal("failed to load config file", cfgErr)
	}
	// Prepare database connenction
	db, dbErr := database.ConnectToPostgres(cfg.Postgres)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	log.Println("postgres connected...")
	database.Migrate(db)

	// Prepare auth controller dependencies
	userRepo := postgres.NewUserRepo(db)
	jwtManager := auth.NewJwtManager(cfg.JWT)
	authService := auth.NewService(userRepo, jwtManager)
	authCtrl := controller.NewAuthController(authService)

	// Initialize gin framework engine
	r := gin.Default()
	http.InitAuthRouter(r, authCtrl)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("server listening on %s\n", addr)
	r.Run(addr)
}
