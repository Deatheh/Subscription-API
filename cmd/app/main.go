package main

import (
	"fmt"
	"log"
	_ "subscription/docs"
	"subscription/internal/config"
	"subscription/internal/repository"
	"subscription/internal/repository/db"
	"subscription/internal/route"
	"subscription/internal/service"

	"github.com/joho/godotenv"
)

// @title Subscription API
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}
	envConf := config.NewEnvConfig()
	config.PrintConfigWithHiddenSecrets(envConf)

	dbRepo, err := db.NewDatabaseInstance(envConf)
	if err != nil {
		log.Fatal(err)
	}
	defer dbRepo.Close()

	repository := &repository.Repository{DatabaseRepository: dbRepo}
	services := service.NewService(repository, envConf)
	handlers := route.NewRouter(services, envConf)

	if err := handlers.InitRoutes().Run(fmt.Sprintf(":%v", envConf.Application.Port)); err != nil {
		log.Fatal(fmt.Errorf("server run error: %w", err))
	}
}
