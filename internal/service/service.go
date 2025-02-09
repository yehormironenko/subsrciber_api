package service

import (
	"go.uber.org/zap"
	"subsctiption-service/internal/repository"
)

type Service struct {
	EchoServiceInterface
	SubscriberServiceInterface
}

func NewService(repo repository.RepositoryInterface, logger *zap.Logger) ServiceInterface {
	return &Service{
		EchoServiceInterface:       NewEchoService(logger),
		SubscriberServiceInterface: NewSubscriberService(repo, logger),
	}
}
