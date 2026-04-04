package main

import (
	"fmt"
	"os"

	"subscription-service/internal/client"
	"subscription-service/internal/config"
	"subscription-service/internal/controller/handlers"
	"subscription-service/internal/repository"
	"subscription-service/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	c, err := config.NewConfig(getConfigPath())
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

func getConfigPath() string {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return "internal/config/config.yaml"
	}
	return configPath
}
