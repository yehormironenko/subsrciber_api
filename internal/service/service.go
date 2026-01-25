package service

import (
	"subscription-service/internal/repository"

	"go.uber.org/zap"
)

type Service struct {
	EchoServiceInterface
	UserServiceInterface
	SubscriberServiceInterface
	SubscriptionsServiceInterface
}

func NewService(repo repository.RepositoryInterface, logger *zap.Logger) ServiceInterface {
	return &Service{
		EchoServiceInterface:          NewEchoService(logger),
		UserServiceInterface:          NewUserService(repo, logger),
		SubscriberServiceInterface:    NewSubscriber(repo, logger),
		SubscriptionsServiceInterface: NewSubscriptions(repo, logger),
	}
}
