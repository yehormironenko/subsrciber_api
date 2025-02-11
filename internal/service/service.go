package service

import (
	"go.uber.org/zap"
	"subsctiption-service/internal/repository"
)

type Service struct {
	EchoServiceInterface
	UserServiceInterface
}

func NewService(repo repository.RepositoryInterface, logger *zap.Logger) ServiceInterface {
	return &Service{
		EchoServiceInterface: NewEchoService(logger),
		UserServiceInterface: NewUserService(repo, logger),
	}
}
