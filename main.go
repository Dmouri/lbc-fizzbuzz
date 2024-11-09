package main

import (
	"lbc/fizzbuzz/api"
	"lbc/fizzbuzz/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()

	router := gin.New()

	fizzBuzzService := service.NewFizzBuzzService()
	api.SetupFizzBuzzController(logger, router, fizzBuzzService)

	logger.Info("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
