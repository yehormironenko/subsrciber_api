package main

import (
	"fmt"

	"subsctiption-service/internal/client"
	"subsctiption-service/internal/config"
	"subsctiption-service/internal/controller/handlers"
	"subsctiption-service/internal/repository"
	"subsctiption-service/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	c, err := config.NewConfig("internal/config/config.yaml")
	if err != nil {
		panic(err)
	}

	logger.Info("configuration loaded successful", zap.Any("config", c))

	// PostgresClient
	postgres := client.CreatePostgresClient(c, logger)

	repository := repository.NewRepository(postgres, logger)
	// services
	services := service.NewService(repository, logger)

	r := gin.Default()
	handlers := handlers.NewHandler(services, logger)

	handlers.Register(r)

	r.Run(fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port))
}
