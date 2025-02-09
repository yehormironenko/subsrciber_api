package service

import (
	"context"
	"time"

	"go.uber.org/zap"
	"subsctiption-service/internal/model/db"
	"subsctiption-service/internal/model/request"
	"subsctiption-service/internal/model/response"
	"subsctiption-service/internal/repository"
)

type Subscriber struct {
	repo   repository.RepositoryInterface
	logger *zap.Logger
}

func NewSubscriberService(repository repository.RepositoryInterface, logger *zap.Logger) *Subscriber {
	return &Subscriber{
		repo:   repository,
		logger: logger,
	}
}

func (s *Subscriber) Subscribe(ctx context.Context, user request.Subscriber) (response.Subscriber, error) {
	s.logger.Info("subscriber service called")
	sub, err := s.repo.Subscribe(ctx, user)
	if err != nil {
		s.logger.Error("subscriber service failed", zap.Error(err))
		return response.Subscriber{}, err
	}

	return dbModelToJsonModel(sub), nil
}

func dbModelToJsonModel(subscriber db.Subscribe) response.Subscriber {
	return response.Subscriber{
		UserId:    subscriber.UserID,
		Email:     subscriber.Email,
		CreatedAt: subscriber.CreateTime.Format(time.RFC3339),
	}
}
