package service

import (
	"go.uber.org/zap"
	"subsctiption-service/internal/repository"
)

type Service struct {
	EchoServiceInterface
	UserServiceInterface
	SubscriberServiceInterface
}

func NewService(repo repository.RepositoryInterface, logger *zap.Logger) ServiceInterface {
	return &Service{
		EchoServiceInterface:       NewEchoService(logger),
		UserServiceInterface:       NewUserService(repo, logger),
		SubscriberServiceInterface: NewSubscriber(repo, logger),
	}
}
