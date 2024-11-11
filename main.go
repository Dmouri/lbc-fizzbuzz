package main

import (
	"lbc/fizzbuzz/api"
	"lbc/fizzbuzz/internal"
	"lbc/fizzbuzz/repository"
	"lbc/fizzbuzz/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()

	router := gin.New()

	fizzBuzzRepository := repository.NewFizzBuzzRepository(internal.Clients.PostgreSQL(), logger)
	fizzBuzzService := service.NewFizzBuzzService(fizzBuzzRepository)
	api.SetupFizzBuzzController(logger, router, fizzBuzzService, fizzBuzzRepository)

	logger.Info("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
